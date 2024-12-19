'use server'

import dayjs from 'dayjs';
import { redirect } from 'next/navigation';
import * as v from 'valibot';

const RegisterSchema = v.pipe(
    v.object({
        email: v.pipe(
            v.string('Your email must be a string.'),
            v.nonEmpty('Please enter an email.'),
            v.email('The email address is typed incorrectly.')
        ),
        // https://valibot.dev/api/string/
        password: v.pipe(
            v.string('Your password must be a string.'),
            v.nonEmpty('Please enter a password.'),
            v.minLength(8, 'Your password must have 8 characters or more.'),
            v.regex(/[a-z]/, 'Your password must contain a lowercase letter.'),
            v.regex(/[A-Z]/, 'Your password must contain a uppercase letter.'),
            v.regex(/[0-9]/, 'Your password must contain a number.'),
            v.regex(/[!@#$%^<>&*]/, 'Your password must contain a special character.'),
        ),
        password_repeat: v.string(),
        // https://valibot.dev/guides/pipelines/
        username: v.pipe(
            v.string('Your username must be a string.'),
            v.nonEmpty('Please enter a username.'),
            v.regex(/^[a-z0-9_-]{4,60}$/iu, 'Your username must be between 4 and 60 characters long and can only contain letters, numbers, underscores and hyphens.'),
        ),
        dateofbirth: v.pipe(
            v.date('Please select your date of birth.'),
            v.maxValue(new Date(dayjs().subtract(18, "year").format("YYYY-MM-DD")), "You must be above the age of 18 for Opinit to legally be allowed to store and process your data under EU law.")
        ),
        question: v.pipe(
            v.string('The answer must be a string.'),
            v.nonEmpty('Please enter an answer.'),
            v.regex(/^(\bTomas|\bArturo)$/iu, 'That is incorrect.'),
        ),
    }),
    v.forward(
        v.partialCheck(
            [['password'], ['password_repeat']],
            (input) => input.password === input.password_repeat,
            'The two passwords do not match.'
        ),
        ['password_repeat']
    )
);

export async function trySignUp(prevState: any, formData: FormData) {
    const rawFormData = [
        ["email", formData.get('email')],
        ["password", formData.get('password')],
        ["password_repeat", formData.get('password_repeat')],
        ["username", formData.get('username')],
        ["dateofbirth", formData.get('dateofbirth') ? new Date(Date.parse(formData.get('dateofbirth').toString())) : null],
        ["question", formData.get('question')],
    ]

    const parseResult = await v.safeParseAsync(RegisterSchema, Object.fromEntries(rawFormData))

    if (parseResult.issues) {
        //console.log(`The following issues were found: '${v.flatten<typeof LoginSchema>(parseResult.issues).nested?.password}'.`)
        if (v.flatten<typeof RegisterSchema>(parseResult.issues).nested) {
            const emailIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.email?.[0]
            const passwordIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.password?.[0]
            const password_repeatIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.password_repeat?.[0]
            const usernameIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.username?.[0]
            const dateofbirthIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.dateofbirth?.[0]
            const questionIssues = v.flatten<typeof RegisterSchema>(parseResult.issues).nested?.question?.[0]

            return {
                origin: prevState.origin,
                success: false,
                data: {
                    email: rawFormData[0][1],
                    password: rawFormData[1][1],
                    password_repeat: rawFormData[2][1],
                    username: rawFormData[3][1],
                    dateofbirth: formData.get('dateofbirth'),
                    question: rawFormData[5][1]
                },
                errors: {
                    email: emailIssues,
                    password: passwordIssues,
                    password_repeat: password_repeatIssues,
                    username: usernameIssues,
                    dateofbirth: dateofbirthIssues,
                    question: questionIssues
                }
            }
        }
    }

    const signUpDTO = {
        email: rawFormData[0][1],
        password: rawFormData[1][1],
        username: rawFormData[3][1]
    }

    const signUpResponse = await fetch(process.env.BACKEND_HOST + `/signup`, { method: 'POST', body: JSON.stringify(signUpDTO) })

    if (signUpResponse.status == 201) {
        redirect("/login/" + "?" + prevState.origin)
    } else if (signUpResponse.status == 400) {
        return {
            origin: prevState.origin,
            success: false,
            data: {
                email: rawFormData[0][1],
                password: rawFormData[1][1],
                password_repeat: rawFormData[2][1],
                username: rawFormData[3][1],
                dateofbirth: formData.get('dateofbirth'),
                question: rawFormData[5][1]
            },
            errors: {
                email: undefined,
                password: undefined,
                password_repeat: undefined,
                username: undefined,
                dateofbirth: undefined,
                question: undefined
            }
        }
    }
}