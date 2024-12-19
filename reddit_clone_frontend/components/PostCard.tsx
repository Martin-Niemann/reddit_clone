'use client'

import { Post } from "@/types/Post"
import { Card, Text, HStack, Box, Image, Grid, Flex, GridItem, Icon } from "@chakra-ui/react"
import NextImage from "next/image"
import card_image from "@/public/soroe_akademi.avif"
import { HiOutlineHandThumbUp, HiOutlineHandThumbDown, HiOutlineUserCircle, HiOutlineChatBubbleOvalLeftEllipsis } from "react-icons/hi2";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { useRouter } from "next/navigation";

dayjs.extend(relativeTime);

export default function PostCard({ props }: { props: Post }) {
    const router = useRouter()
    console.log("this is the id of the post: " + props.id)

    return (
        <Flex direction={"column"} padding={3} rounded={"2xl"} _hover={{ background: "gray.300" }} cursor={"pointer"} onClick={() => router.push("/p/" + props.id)} /*backgroundColor={"orange.200"}*/>
            <Flex fontSize={14} justifyContent={"flex-start"} alignItems={"center"}>
                <Icon fontSize={"20px"} color={"gray.600"}>
                    <HiOutlineUserCircle />
                </Icon>
                <Text paddingStart={0.5} fontWeight={"semibold"} color={"gray.600"} >{props.username}</Text>
                <Text fontWeight={"normal"} color={"gray.500"} paddingStart={4}>{dayjs(props.updated_at).fromNow()}</Text>
            </Flex>
            <Flex gap={2} paddingTop={2} alignItems={"center"}>
                <Text fontWeight={"medium"} color={"gray.800"} flexGrow={2}>{props.title}</Text>
                <Image rounded={"xl"} flexGrow={1} maxWidth={"120px"} objectFit={"contain"} asChild alt={props.title}>
                    <NextImage width={500} height={500} src={card_image} alt={props.title}></NextImage>
                </Image>
            </Flex>
            <Flex paddingTop={2.5} alignItems={"center"}>
                <Icon fontSize={"20px"} color={"gray.500"}>
                    <HiOutlineHandThumbUp />
                </Icon>
                <Text paddingInline={2.5} fontWeight={"semibold"} color={"gray.600"}>{props.score}</Text>
                <Icon fontSize={"20px"} color={"gray.500"}>
                    <HiOutlineHandThumbDown />
                </Icon>
                <Text paddingStart={7} paddingEnd={1.5} fontWeight={"semibold"} color={"gray.600"}>{props.comments_count}</Text>
                <Icon fontSize={"24px"} color={"gray.500"}>
                    <HiOutlineChatBubbleOvalLeftEllipsis />
                </Icon>
            </Flex>
        </Flex>
    )
}