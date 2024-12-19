'use client'

import { tryCreatePost } from "@/app/api/tryCreatePost";
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
    const community = React.use(searchParams).community

    const initialState = {
        community_id: community,
        success: Boolean,
        data: {
            title: "",
            link: "",
            text: ""
        },
        errors: {
            title: "",
        }
    };

    const [state, formAction] = useActionState(tryCreatePost, initialState)

    return (
        <>
            <Text fontWeight={"bold"} fontSize={24} marginBottom={6} marginTop={1}>What&apos;s on your mind?</Text >
            <Form action={formAction}>
                <Stack gap="4" align="flex-start" maxW="sm">
                    <Field label="Title">
                        <Input name="title" type="text" defaultValue={state.data.title} />
                    </Field>
                    <Field label="Link (optional)">
                        <Input name="link" type="url" defaultValue={state.data.link} />
                    </Field>
                    <Field label="Text (optional)" marginBottom={1.5}>
                        <Textarea name="text" borderWidth={1.5} rounded={"xl"} resize="both" defaultValue={state.data.text} />
                    </Field>
                    <Button type="submit" alignSelf="flex-center">
                        Post
                    </Button>
                    {state.errors.title}
                </Stack>
            </Form>
        </>

    )
}