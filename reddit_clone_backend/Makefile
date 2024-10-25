# https://github.com/sqlc-dev/sqlc/discussions/380#discussioncomment-3597491
# A nice solution for replacing SQLC's Nullstrings with normal Strings, in order to flatten the generated JSON
# install goimports: `go install golang.org/x/tools/cmd/goimports@latest`

run: 
	go run .
sql:
	sqlc generate
	sed -i -e 's/sql.NullString/*string/g' database/*.go
	sed -i -e 's/sql.NullInt32/*int32/g' database/*.go
	sed -i -e 's/sql.NullTime/*time.Time/g' database/*.go
	goimports -w ./database/