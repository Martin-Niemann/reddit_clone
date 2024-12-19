'use client'

import { Post } from "@/types/Post"
import { Card, Text, HStack, Box, Image, Grid, Flex, GridItem, Icon } from "@chakra-ui/react"
import NextImage from "next/image"
import card_image from "@/public/card_image.avif"
import { HiOutlineHandThumbUp, HiOutlineHandThumbDown, HiOutlineUserCircle, HiOutlineChatBubbleOvalLeftEllipsis } from "react-icons/hi2";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

export default function PostBanner({ props }: { props: Post }) {
    return (
        <Flex direction={"column"} paddingTop={2} rounded={"2xl"} _hover={{ background: "gray.300" }} /*backgroundColor={"orange.200"}*/>
            <Flex paddingStart={0.5} fontSize={14} justifyContent={"flex-start"} alignItems={"center"} paddingBottom={1.5}>
                <Text fontWeight={"semibold"} color={"gray.500"}>c/{props.community_name}</Text>
                <Text fontWeight={"normal"} color={"gray.500"} paddingStart={3}>{dayjs(props.updated_at).fromNow()}</Text>
            </Flex>
            <Flex fontSize={14} justifyContent={"flex-start"} alignItems={"center"}>
                <Icon fontSize={"20px"} color={"gray.600"}>
                    <HiOutlineUserCircle />
                </Icon>
                <Text paddingStart={0.5} fontWeight={"semibold"} color={"gray.600"} >{props.username}</Text>
            </Flex>
            <Flex paddingStart={0.5} gap={2} paddingTop={2} alignItems={"center"}>
                <Text fontWeight={"semibold"} color={"gray.800"} fontSize={"xl"} flexGrow={2}>{props.title}</Text>
            </Flex>
            <Flex paddingStart={0.5} paddingTop={3}>
                <Text fontWeight={"normal"} color={"gray.solid"} flexGrow={2}>{props.text}</Text>
            </Flex>
            <Flex paddingTop={4} gap={3} alignItems={"center"}>
                <Flex justifyContent={"flex-start"} alignItems={"center"} backgroundColor={"gray.100"} borderWidth={1} borderColor={"gray.400"} rounded={"full"} padding={1.5}>
                    <Icon fontSize={"20px"} color={"gray.solid"}>
                        <HiOutlineHandThumbUp />
                    </Icon>
                    <Text paddingInline={2.5} fontWeight={"semibold"} color={"gray.solid"}>{props.score}</Text>
                    <Icon fontSize={"20px"} color={"gray.solid"}>
                        <HiOutlineHandThumbDown />
                    </Icon>
                </Flex>
                <Flex backgroundColor={"gray.100"} borderWidth={1} borderColor={"gray.400"} rounded={"full"} padding={1.5}>
                    <Text paddingStart={2} paddingEnd={1.5} fontWeight={"semibold"} color={"gray.solid"}>{props.comments.length}</Text>
                    <Icon fontSize={"24px"} color={"gray.800"}>
                        <HiOutlineChatBubbleOvalLeftEllipsis />
                    </Icon>
                </Flex>
            </Flex>
        </Flex >
    )
}