'use client'

import { Field } from "@/components/ui/field"
import { Button, Input, Stack } from "@chakra-ui/react"
import Form from 'next/form'
import React, { useActionState } from "react";
import { tryLogin } from "../api/tryLogin";

// https://next.chakra-ui.com/docs/components/field
// https://valibot.dev/guides/issues/

export default function Login({
    searchParams,
}: {
    searchParams: { [key: string]: string | string[] | undefined }
}) {
    const origin = React.use(searchParams).origin

    const initialState = {
        origin: origin,
        success: Boolean,
        errors: {
            password: "",
        }
    };

    const [state, formAction] = useActionState(tryLogin, initialState)

    return (
        <>
            <Form action={formAction}>
                <Stack gap="4" align="flex-start" maxW="sm">
                    <Field label="Email">
                        <Input name="email" type="email" />
                    </Field>
                    <Field label="Password" invalid={state.errors.password != ""} errorText={state.errors.password != "" ? state.errors.password : null}>
                        <Input name="password" type="password" />
                    </Field>
                    <Button type="submit" alignSelf="flex-center">
                        Log in
                    </Button>
                </Stack>
            </Form>
        </>

    )
}