'use client'

import { Button, Flex, Textarea } from "@chakra-ui/react";
import Form from "next/form";
import { useActionState, useEffect, useRef, useState } from "react";
import { Field } from "./ui/field";
import { trySendComment } from "@/app/api/trySendComment";
import { trySendCommentDTO } from "@/types/trySendCommentDTO";

export default function AddComment({ params }: { params: { postId: number } }) {
    const [isHidden, changeIsHidden] = useState(true)

    const initialState: trySendCommentDTO = {
        postId: params.postId,
        isFailed: false,
        success: false
    };

    const [state, formAction] = useActionState(trySendComment, initialState)

    // https://www.reddit.com/r/reactjs/comments/pk9jmd/comment/hc1xp1a/
    const isFirstRun = useRef(true)
    useEffect(() => {
        if (!isFirstRun.current) {
            changeIsHidden(state.success)
        } else {
            isFirstRun.current = false
        }
    }, [state])

    return (
        <>
            <Flex hidden={!isHidden} onClick={() => changeIsHidden(!isHidden)} backgroundColor={"gray.100"} rounded={"full"} justifyContent={"center"} padding={2} borderWidth={1} borderColor={"gray.400"} marginTop={3} _hover={{ backgroundColor: "gray.200" }}>
                Write a comment
            </Flex>

            <Flex paddingTop={3} gap={1.5} direction={"column"}>
                <Form action={formAction} hidden={isHidden}>
                    <Field invalid={state.isFailed} marginBottom={1.5}>
                        <Textarea borderWidth={1.5} name="text" rounded={"xl"} resize="none" placeholder="Actually, I find that…" />
                    </Field>
                    <Flex gap={2} justifyContent={"space-between"}>
                        <Button flex={1} type="submit" width={"full"} rounded={"full"} size={"xs"} backgroundColor={"gray.100"} color={"gray.800"} borderWidth={1} borderColor={"gray.400"}>Send</Button>
                        <Button flex={1} onClick={() => changeIsHidden(!isHidden)} width={"full"} rounded={"full"} size={"xs"} backgroundColor={"gray.500"} color={"white"} borderWidth={1} borderColor={"gray.500"}>Cancel</Button>
                    </Flex>
                </Form>
            </Flex>
        </>
    )
}