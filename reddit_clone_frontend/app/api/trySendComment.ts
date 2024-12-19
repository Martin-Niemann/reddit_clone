'use server'

import * as v from 'valibot';
import { CommentDTO } from '@/types/CommentDTO';
import { getJwtPayload, JwtPayload } from './getJwtPayload';
import { revalidatePath, revalidateTag } from 'next/cache';
import { trySendCommentDTO } from '@/types/trySendCommentDTO';

const AddCommentSchema = v.object({
    text: v.pipe(
        v.string(),
        v.nonEmpty(),
    )
});

export async function trySendComment(prevState: any, formData: FormData) {
    const rawFormData = {
        text: formData.get('text'),
    }

    const parseResult = await v.safeParseAsync(AddCommentSchema, Object.fromEntries(formData.entries()))

    if (parseResult.issues) {
        const fail: trySendCommentDTO = {
            postId: prevState.postId,
            isFailed: true,
            success: prevState.success
        }
        return fail
    }

    const authCookie: JwtPayload | null = await getJwtPayload()

    if (authCookie == null) {
        const fail: trySendCommentDTO = {
            postId: prevState.postId,
            isFailed: true,
            success: prevState.success
        }
        return fail
    }

    const commentDTO: CommentDTO = {
        post_id: prevState.postId,
        parent_id: prevState.parentId,
        user_id: authCookie.id,
        text: rawFormData.text
    }

    const addCommentResponse = await fetch(process.env.BACKEND_HOST + `/comment`, { method: 'POST', headers: { Cookie: "auth" + "=" + authCookie.token + ";" }, body: JSON.stringify(commentDTO) })

    if (addCommentResponse.status == 201) {
        const response = await addCommentResponse.json()
        console.log(response)

        revalidatePath("/p/" + prevState.postId)

        return {
            postId: prevState.postId,
            isFailed: false,
            success: true
        }
    }
    else {
        return {
            postId: prevState.postId,
            isFailed: true,
            success: prevState.success
        }
    }
}