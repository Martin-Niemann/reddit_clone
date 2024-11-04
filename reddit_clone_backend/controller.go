package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gorilla/handlers"
)

type Controller struct {
	handler       http.Handler
	server        http.Server
	service       Service
	jwtSigningKey []byte
	validator     *validator.Validate
}

func (c *Controller) Setup(mysql MysqlConnectionVariables) error {
	c.service = NewService(mysql)

	mux := c.setupRoutes(http.NewServeMux())

	c.handler = handlers.LoggingHandler(os.Stdout, mux)

	c.server = http.Server{Addr: ":8090", Handler: c.handler}

	c.jwtSigningKey = []byte(os.Getenv("JWT_SIGNING_KEY"))

	c.validator = validator.New(validator.WithRequiredStructEnabled())

	return nil
}

func (c *Controller) Run() {
	go func() {
		log.Fatal(c.server.ListenAndServe())
	}()
	log.Printf("HTTP server succesfully started on address: %s.", c.server.Addr)
}

func (c *Controller) Shutdown(timeoutCtx context.Context) {
	c.service.closeConnection()
	c.server.Shutdown(timeoutCtx)
}

func sendJsonResponse[T any](bodyContent T, statusCode int, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	if statusCode != 0 {
		writer.WriteHeader(statusCode)
	}

	err := json.NewEncoder(writer).Encode(bodyContent)
	if err != nil {
		log.Println(err.Error())

		http.Error(writer, err.Error(), 500)
	}
}

func sendErrorResponse(writer http.ResponseWriter, serviceError ServiceError) {
	switch serviceError.Type {
	case NoResult:
		http.Error(writer, http.StatusText(204), 204)
	case InvalidInput:
		http.Error(writer, http.StatusText(400), 400)
	case InvalidArgument:
		sendJsonResponse(serviceError.ValidationErrors, 400, writer)
	case InvalidEmailAndOrPassword:
		http.Error(writer, http.StatusText(401), 401)
	case UserAlreadyExists:
		sendJsonResponse(serviceError.ValidationErrors, 403, writer)
	case UserDoesntExist:
		http.Error(writer, http.StatusText(400), 400)
	case UnexpectedError:
		http.Error(writer, http.StatusText(500), 500)
	default:
		panic(fmt.Sprintf("unexpected main.ServiceError: %+v", serviceError))
	}
}

func tryParseJson[T any](req *http.Request, model *T) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(model)
}

func createValidationErrorsStruct(err validator.ValidationErrors) []ValidationError {
	// https://dev.to/franciscomendes10866/how-to-validate-data-in-golang-1f87
	var errors []ValidationError

	for _, err := range err {
		var valiError ValidationError
		valiError.Field = err.Field()
		valiError.Tag = err.Tag()
		valiError.Value = err.Param()
		errors = append(errors, valiError)
	}

	return errors
}

// https://medium.com/geekculture/learn-go-middlewares-by-examples-da5dc4a3b9aa
func (c *Controller) authenticationAndAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

		token := req.CookiesNamed("Auth")[0]
		if token == nil {
			sendErrorResponse(writer, ServiceError{Type: MissingAuthCookie})
			return
		}

		tokenStruct := Token{Token: token.Value}
		println("the token string is: ", tokenStruct.Token)

		err := c.validator.Struct(tokenStruct)
		if err != nil {
			validationErrors := createValidationErrorsStruct(err.(validator.ValidationErrors))
			sendErrorResponse(writer, ServiceError{Type: InvalidArgument, ValidationErrors: validationErrors})
			return
		}

		parsedToken, err := jwt.Parse(tokenStruct.Token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return c.jwtSigningKey, nil
		})
		if err != nil {
			sendErrorResponse(writer, ServiceError{Type: InvalidArgument, Text: err.Error()})
			return
		}

		var ctx context.Context
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			fmt.Println("the id and iat are: ", claims["id"], claims["iat"])
			ctx = context.WithValue(req.Context(), "id", claims["id"])
		} else {
			sendErrorResponse(writer, ServiceError{Type: InvalidArgument})
			return
		}

		newReq := req.WithContext(ctx)

		next.ServeHTTP(writer, newReq)
	})
}
