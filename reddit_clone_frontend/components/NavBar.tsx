import { Image, Flex, Grid, GridItem, Box, Heading } from "@chakra-ui/react";
import NextImage from "next/image"
import logo from "../public/logo5.avif"

export default function NavBar() {
    return (
        <Grid templateColumns="repeat(4, 1fr)" gap="2" h="16" alignItems="center">
            <GridItem colSpan={1} backgroundColor="pink.200">
                <Flex direction="row" alignItems="center">
                    <Image asChild boxSize="45px" alt="reddit_clone logo" rounded="full">
                        <NextImage src={logo} alt="reddit_clone logo"></NextImage>
                    </Image>
                    <Heading lg={{ fontSize: "2xl"}} md={{ fontSize: "xl"}} hideBelow="md" fontWeight="bold" marginStart="4">reddit_clone</Heading>
                </Flex>
            </GridItem>
            <GridItem colSpan={2} backgroundColor="green.200">
                <Box rounded="md" borderColor="blackAlpha.400">
                    search bar
                </Box>
            </GridItem>
            <GridItem colSpan={1} backgroundColor="yellow.200">
                some other stuff
            </GridItem>
        </Grid>

        //<Stack justifyContent="space-between" direction="row" h="20" gap="2">
        //</Stack>
    )
}