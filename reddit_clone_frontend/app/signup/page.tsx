'use client'

import { Field } from "@/components/ui/field"
import { Button, Input, Stack, Text } from "@chakra-ui/react"
import Form from 'next/form'
import React, { useActionState } from "react";
import { trySignUp } from "../api/trySignUp";

// https://next.chakra-ui.com/docs/components/field
// https://valibot.dev/guides/issues/

export default function SignUp({
    searchParams,
}: {
    searchParams: { [key: string]: string | string[] | undefined }
}) {
    const origin = React.use(searchParams).origin

    const initialState = {
        origin: origin,
        success: Boolean,
        data: {
            email: "",
            password: "",
            password_repeat: "",
            username: "",
            dateofbirth: null,
            question: ""
        },
        errors: {
            email: undefined,
            password: undefined,
            password_repeat: undefined,
            username: undefined,
            dateofbirth: undefined,
            question: undefined
        }
    };

    const [state, formAction] = useActionState(trySignUp, initialState)

    return (
        <>
            <Text fontWeight={"bold"} fontSize={24} marginBottom={3} marginTop={1}>Create your account</Text >
            <Form action={formAction}>
                <Stack gap="4" align="flex-start" maxW="sm">
                    <Field label="Email" invalid={state.errors.email != undefined} errorText={state.errors.email != undefined ? state.errors.email : null}>
                        <Input name="email" type="email" defaultValue={state.data.email} />
                    </Field>
                    <Field label="Password" invalid={state.errors.password != undefined} errorText={state.errors.password != undefined ? state.errors.password : null}>
                        <Input name="password" type="password" defaultValue={state.data.password} />
                    </Field>
                    <Field label="Repeat password" invalid={state.errors.password_repeat != undefined} errorText={state.errors.password_repeat != undefined ? state.errors.password_repeat : null}>
                        <Input name="password_repeat" type="password" defaultValue={state.data.password_repeat} />
                    </Field>
                    <Field label="Username" invalid={state.errors.username != undefined} errorText={state.errors.username != undefined ? state.errors.username : null}>
                        <Input name="username" type="text" defaultValue={state.data.username} />
                    </Field>
                    <Field label="Date of birth" invalid={state.errors.dateofbirth != undefined} errorText={state.errors.dateofbirth != undefined ? state.errors.dateofbirth : null}>
                        <Input name="dateofbirth" type="date" defaultValue={state.data.dateofbirth} />
                    </Field>
                    <Field label="Please enter the first name of my course teacher or my examiner" invalid={state.errors.question != undefined} errorText={state.errors.question != undefined ? state.errors.question : null}>
                        <Input name="question" type="text" defaultValue={state.data.question} />
                    </Field>
                    <Button type="submit" alignSelf="flex-center">
                        Create
                    </Button>
                </Stack>
            </Form>
        </>

    )
}