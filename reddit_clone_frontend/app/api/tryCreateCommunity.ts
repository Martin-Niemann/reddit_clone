'use server'

import * as v from 'valibot';
import { getJwtPayload, JwtPayload } from './getJwtPayload';
import { tryCreateCommunityDTO } from '@/types/tryCreateCommunityDTO';
import { redirect } from 'next/navigation';

const CreateCommunitySchema = v.object({
    url: v.pipe(
        v.string(),
        v.nonEmpty()
    ),
    displayname: v.pipe(
        v.string(),
        v.nonEmpty()
    ),
    description: v.pipe(
        v.string(),
        v.nonEmpty()
    )
});

export async function tryCreateCommunity(prevState: any, formData: FormData) {
    const rawFormData = {
        url: formData.get('url'),
        displayname: formData.get('displayname'),
        description: formData.get('description'),
    }

    const parseResult = await v.safeParseAsync(CreateCommunitySchema, Object.fromEntries(formData.entries()))

    if (parseResult.issues) {
        //console.log(`The following issues were found: '${v.flatten<typeof LoginSchema>(parseResult.issues).nested?.password}'.`)
        if (v.flatten<typeof CreateCommunitySchema>(parseResult.issues).nested) {
            const urlIssues = v.flatten<typeof CreateCommunitySchema>(parseResult.issues).nested?.url
            const displaynameIssues = v.flatten<typeof CreateCommunitySchema>(parseResult.issues).nested?.displayname
            const descriptionIssues = v.flatten<typeof CreateCommunitySchema>(parseResult.issues).nested?.description

            return {
                success: false,
                data: {
                    url: rawFormData.url,
                    displayname: rawFormData.displayname,
                    description: rawFormData.description
                },
                errors: {
                    url: urlIssues,
                    displayname: displaynameIssues,
                    description: descriptionIssues
                }
            }
        }
    }

    const authCookie: JwtPayload | null = await getJwtPayload()

    if (authCookie == null) {
        return {
            success: false,
            data: {
                url: rawFormData.url,
                displayname: rawFormData.displayname,
                description: rawFormData.description
            },
            errors: {
                url: "",
                displayname: "",
                description: "Please log in."
            }
        }
    }

    const communityDTO: tryCreateCommunityDTO = {
        url: rawFormData.url?.toString(),
        displayname: rawFormData.displayname?.toString(),
        description: rawFormData.description?.toString()
    }

    const addCommunityResponse = await fetch(process.env.BACKEND_HOST + `/subreddit`, { method: 'POST', headers: { Cookie: "auth" + "=" + authCookie.token + ";" }, body: JSON.stringify(communityDTO) })

    if (addCommunityResponse.status == 201) {
        redirect("/c/" + rawFormData.url)
    } else {
        return {
            success: false,
            data: {
                url: rawFormData.url,
                displayname: rawFormData.displayname,
                description: rawFormData.description
            },
            errors: {
                url: "",
                displayname: "",
                description: "Please try again later."
            }
        }
    }
}