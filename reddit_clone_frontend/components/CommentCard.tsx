'use client'

import { Comment } from "@/types/Comment";
import { Button, Collapsible, Flex, Icon, Text, Textarea } from "@chakra-ui/react";
import Form from "next/form";
import { HiOutlineChatBubbleOvalLeftEllipsis, HiOutlineHandThumbDown, HiOutlineHandThumbUp } from "react-icons/hi2";
import { Field } from "./ui/field";
import { useActionState, useEffect, useRef, useState } from "react";
import { trySendComment } from "@/app/api/trySendComment";
import { getShortenedDateString } from "@/helpers/getShortenedDateString";
import { trySendCommentDTO } from "@/types/trySendCommentDTO";

export interface CommentCardProps {
    id: number;
    parent_id?: number;
    created_at: Date;
    updated_at: Date;
    username: string;
    text: string;
    score: number;
    post_id: number;
}

export default function CommentCard({ props }: { props: CommentCardProps }) {
    const [isHidden, myChangeIsHidden] = useState(false)

    const initialState: trySendCommentDTO = {
        postId: props.post_id,
        parentId: props.id,
        isFailed: false,
        success: false
    };

    console.log(props.text, props.id, props.created_at)

    const [state, formAction] = useActionState(trySendComment, initialState)

    // https://www.reddit.com/r/reactjs/comments/pk9jmd/comment/hc1xp1a/
    const isFirstRun = useRef(true)
    useEffect(() => {
        if (!isFirstRun.current) {
            myChangeIsHidden(false)
        } else {
            isFirstRun.current = false
        }
    }, [state])

    return (
        <Collapsible.Root open={isHidden}>
            <Flex direction={"column"} marginBlock={2.5} marginBottom={3}>
                <Flex fontSize={14} justifyContent={"flex-start"} alignItems={"center"} paddingBottom={1}>
                    <Text fontWeight={"semibold"} color={"gray.500"}>{props.username}</Text>
                    <Text fontWeight={"normal"} color={"gray.500"} paddingStart={3}>{getShortenedDateString(props.updated_at)}</Text>
                </Flex>
                <Flex paddingStart={0}>
                    <Text fontWeight={"medium"} fontSize={14} color={"gray.700"}>{props.text}</Text>
                </Flex>
                <Flex paddingTop={3} gap={3} alignItems={"center"}>
                    <Flex justifyContent={"flex-start"} alignItems={"center"} backgroundColor={"gray.100"} rounded={"full"} padding={1.5}>
                        <Icon fontSize={"18px"} color={"gray.solid"}>
                            <HiOutlineHandThumbUp />
                        </Icon>
                        <Text paddingInline={3} paddingRight={3.5} fontWeight={"medium"} fontSize={14} color={"gray.700"}>{props.score}</Text>
                        <Icon fontSize={"18px"} color={"gray.solid"}>
                            <HiOutlineHandThumbDown />
                        </Icon>
                    </Flex>
                    <Flex onClick={() => myChangeIsHidden(!isHidden)} backgroundColor={"gray.100"} rounded={"full"} padding={1.5}>
                        <Text paddingStart={2} paddingEnd={1.5} fontWeight={"medium"} fontSize={14} color={"gray.700"}>Reply</Text>
                        <Icon fontSize={"20px"} color={"gray.800"}>
                            <HiOutlineChatBubbleOvalLeftEllipsis />
                        </Icon>
                    </Flex>
                </Flex>
                <Collapsible.Content>
                    <Flex paddingTop={3} gap={1.5} direction={"column"}>
                        <Form action={formAction}>
                            <Field invalid={state.isFailed} marginBottom={1.5}>
                                <Textarea borderWidth={1.5} name="text" rounded={"xl"} resize="none" placeholder="Actually, I find that…" />
                            </Field>
                            <Flex gap={2} justifyContent={"space-between"}>
                                <Button flex={1} type="submit" width={"full"} rounded={"full"} size={"xs"} backgroundColor={"gray.100"} color={"gray.800"} borderWidth={1} borderColor={"gray.400"}>Send</Button>
                                <Button onClick={() => myChangeIsHidden(!isHidden)} flex={1} width={"full"} rounded={"full"} size={"xs"} backgroundColor={"gray.500"} color={"white"} borderWidth={1} borderColor={"gray.500"}>Cancel</Button>
                            </Flex>
                        </Form>
                    </Flex>
                </Collapsible.Content>
            </Flex>
        </Collapsible.Root>
    )
}