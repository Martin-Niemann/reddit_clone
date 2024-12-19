'use client'

import { tryCreateCommunity } from "@/app/api/tryCreateCommunity";
import { Field } from "@/components/ui/field";
import { Button, Group, Input, InputAddon, Stack, Textarea, Text } from "@chakra-ui/react";
import Form from "next/form";
import React from "react";
import { useActionState } from "react";

export default function Create({
    searchParams,
}: {
    searchParams: { [key: string]: string | string[] | undefined }
}) {
    const origin = React.use(searchParams).origin

    const initialState = {
        origin: origin,
        success: Boolean,
        data: {
            url: "",
            displayname: "",
            description: ""
        },
        errors: {
            url: "",
            displayname: "",
            description: ""
        }
    };

    const [state, formAction] = useActionState(tryCreateCommunity, initialState)

    return (
        <>
            <Text fontWeight={"bold"} fontSize={24} marginBottom={6} marginTop={1}>Create a new community</Text >
            <Form action={formAction}>
                <Stack gap="4" align="flex-start" maxW="sm">
                    <Field label="URL" invalid={state.errors.url != "" ? true : false}>
                        <Group attached>
                            <InputAddon>c/</InputAddon>
                            <Input name="url" type="text" defaultValue={state.data.url} />
                        </Group>
                    </Field>
                    <Field label="Display name" invalid={state.errors.displayname != "" ? true : false}>
                        <Input name="displayname" type="text" defaultValue={state.data.displayname} />
                    </Field>
                    <Field label="Description" marginBottom={1.5} invalid={state.errors.description != "" ? true : false}>
                        <Textarea name="description" borderWidth={1.5} rounded={"xl"} resize="both" defaultValue={state.data.description} />
                    </Field>
                    <Button type="submit" alignSelf="flex-center">
                        Create
                    </Button>
                </Stack>
            </Form>
        </>

    )
}