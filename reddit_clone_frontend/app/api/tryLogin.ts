'use server'

import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import * as v from 'valibot';

const LoginSchema = v.object({
    email: v.pipe(
        v.string('Your email must be a string.'),
        v.nonEmpty('Please enter your email.'),
    ),
    password: v.pipe(
        v.string('Your password must be a string.'),
        v.nonEmpty('Please enter your password.'),
        v.minLength(8, 'Your password has 8 characters or more.'),
    ),
});

export async function tryLogin(prevState: any, formData: FormData) {
    const rawFormData = {
        email: formData.get('email'),
        password: formData.get('password'),
    }

    console.log(`You entered '${rawFormData.email}'.`)

    const parseResult = await v.safeParseAsync(LoginSchema, Object.fromEntries(formData.entries()))

    if (parseResult.issues) {
        if (v.flatten<typeof LoginSchema>(parseResult.issues).nested) {
            const passwordIssues = v.flatten<typeof LoginSchema>(parseResult.issues).nested?.password
            console.log(passwordIssues)
            return {
                origin: prevState.origin,
                succes: false,
                errors: {
                    password: passwordIssues?.[0]
                }
            }
        }
    }

    const loginResponse = await fetch(process.env.BACKEND_HOST + `/login`, { method: 'POST', body: JSON.stringify(rawFormData) })

    if (loginResponse.status == 200) {

        interface TokenCookie {
            token: string
            expires: Date
        }

        const tokenCookie: TokenCookie = await loginResponse.json()
        console.log(tokenCookie.token, tokenCookie.expires)

        const cookieStore = await cookies()
        cookieStore.set({
            name: "auth",
            value: tokenCookie.token,
            expires: tokenCookie.expires,
            httpOnly: true,
            secure: false,
        })

        redirect(prevState.origin)
    } else if (loginResponse.status == 401) {
        return {
            origin: prevState.origin,
            succes: false,
            errors: {
                password: "The password or email is wrong."
            }
        }
    }
}