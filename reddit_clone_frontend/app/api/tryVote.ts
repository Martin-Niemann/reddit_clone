import { getJwtPayload, JwtPayload } from "./getJwtPayload"

export async function tryVote({ props }: { props: { vote: Boolean } }) {
    const authCookie: JwtPayload | null = await getJwtPayload()

    if (authCookie == null) {
        return {
            insertedComment: null,
            postId: prevState.postId,
            isFailed: true,
            errors: {
                text: "Not logged in"
            }
        }
    }

    console.log("the post id is: " + prevState.postId)

    const commentDTO: CommentDTO = {
        post_id: prevState.postId,
        user_id: authCookie.id,
        text: rawFormData.text
    }

    const addCommentResponse = await fetch(process.env.BACKEND_HOST + `/comment`, { method: 'POST', headers: { Cookie: "auth" + "=" + authCookie.token + ";" }, body: JSON.stringify(commentDTO) })

    if (addCommentResponse.status == 201) {
        const response = await addCommentResponse.json()
        console.log(response)

        //revalidateTag('p')
        revalidatePath("/p/" + prevState.postId)

        return {
            insertedComment: response,
            postId: prevState.postId,
            isFailed: false,
            errors: {
                text: null
            }
        }
    }
    else {
        return {
            insertedComment: null,
            postId: prevState.postId,
            isFailed: true,
            errors: {
                text: addCommentResponse.json
            }
        }
    }
}