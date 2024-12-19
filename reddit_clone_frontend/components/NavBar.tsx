'use client'

import { Image, Flex, Grid, GridItem, Box, Heading, Input, Icon, IconButton, Spinner } from "@chakra-ui/react";
import NextImage from "next/image"
import logo from "@/public/logo5.avif"
import { HiOutlineMagnifyingGlass, HiOutlineBars3, HiOutlineBarsArrowDown, HiOutlineUserCircle } from "react-icons/hi2";
import { InputGroup } from "./ui/input-group";
import { useBottomDrawerStore } from "@/stores/BottomDrawerStore";
import { useShallow } from "zustand/shallow";
import React, { useEffect, useState } from "react";
import { getJwtPayload, JwtPayload } from "@/app/api/getJwtPayload";

export default function NavBar() {
    const { drawerState, changeDrawerState } = useBottomDrawerStore(useShallow((state) => {
        return { drawerState: state.isOpen, changeDrawerState: state.change }
    }))

    const [jwtPayload, setJwtPayload] = useState<JwtPayload>()

    useEffect(() => {
        const updatejwtPayload = async () => { setJwtPayload(await getJwtPayload()) }
        updatejwtPayload()
    }, [])

    return (
        <Grid templateColumns="repeat(4, 1fr)" h="16" alignItems="center">
            <GridItem colSpan={1} /*backgroundColor="pink.200"*/>
                <Flex direction="row" alignItems="center">
                    <Image asChild shadow={"inner"} boxSize="45px" alt="reddit_clone logo" rounded="full">
                        <NextImage src={logo} alt="reddit_clone logo"></NextImage>
                    </Image>
                    <Heading lg={{ fontSize: "2xl" }} md={{ fontSize: "xl" }} hideBelow="md" fontWeight="bold" marginStart="4">opinit</Heading>
                </Flex>
            </GridItem>
            <GridItem paddingStart={2} colSpan={2} marginInline={-6} /*backgroundColor="green.200"*/>
                <Box textAlign={"center"}>
                    <InputGroup startElement={
                        <Icon fontSize={"18px"} color={"gray.400"}>
                            <HiOutlineMagnifyingGlass />
                        </Icon>}>
                        <Input rounded={"full"} placeholder="Search" />
                    </InputGroup>
                </Box>
            </GridItem>
            <GridItem colSpan={1} /*backgroundColor="yellow.200"*/>
                <Flex justifyContent={"end"} marginEnd={-2}>
                    <IconButton onClick={changeDrawerState} aria-label="Open bottom drawer with user settings." variant={"ghost"} size={"2xl"}>
                        {jwtPayload != undefined ? (jwtPayload.id != 0 ? <HiOutlineUserCircle /> : (drawerState ? <HiOutlineBarsArrowDown /> : <HiOutlineBars3 />)) : <Spinner size={"md"} color={"gray.900"} borderWidth={1.5} />}
                    </IconButton>
                </Flex>
            </GridItem>
        </Grid>
    )
}