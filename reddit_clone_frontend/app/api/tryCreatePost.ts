'use server'

import { getJwtPayload, JwtPayload } from './getJwtPayload';
import { redirect } from 'next/navigation';
import { tryCreatePostDTO } from '@/types/tryCreatePostDTO';

export async function tryCreatePost(prevState: any, formData: FormData) {
    const rawFormData = {
        title: formData.get('title'),
        link: formData.get('link'),
        text: formData.get('text'),
    }

    const authCookie: JwtPayload | null = await getJwtPayload()

    if (authCookie == null) {
        return {
            community_id: prevState.community_id,
            success: false,
            data: {
                title: rawFormData.title,
                link: rawFormData.link,
                text: rawFormData.text
            },
            errors: {
                title: "Please log in.",
            }
        }
    }

    if (rawFormData.title == "") {
        return {
            community_id: prevState.community_id,
            success: false,
            data: {
                title: rawFormData.title,
                link: rawFormData.link,
                text: rawFormData.text
            },
            errors: {
                title: "A title is required.",
            }
        }
    }

    const postDTO: tryCreatePostDTO = {
        title: rawFormData.title?.toString(),
        link: rawFormData.link?.toString(),
        text: rawFormData.text?.toString(),
        community_id: Number(prevState.community_id)
    }

    console.log(postDTO.title, postDTO.link, postDTO.text, postDTO.community_id, authCookie.token)

    const addPostResponse = await fetch(process.env.BACKEND_HOST + `/post`, { method: 'POST', headers: { Cookie: "auth" + "=" + authCookie.token + ";" }, body: JSON.stringify(postDTO) })

    if (addPostResponse.status == 201) {
        const responseJson = await addPostResponse.json()
        redirect("/p/" + responseJson.post_id)
    } else {
        return {
            ommunity_id: prevState.community_id,
            success: false,
            data: {
                title: rawFormData.title,
                link: rawFormData.link,
                text: rawFormData.text
            },
            errors: {
                title: "Please try again later.",
            }
        }
    }
}