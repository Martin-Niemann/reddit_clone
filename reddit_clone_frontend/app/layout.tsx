'use client'

import type { Metadata } from "next";
import { Provider } from "@/components/ui/provider"
import { Grid, GridItem, DrawerTrigger, DrawerBackdrop, DrawerRoot, DrawerContent, DrawerHeader, DrawerTitle, DrawerBody, DrawerFooter, DrawerActionTrigger, DrawerCloseTrigger, Button, Flex, Icon, Text } from "@chakra-ui/react";
import NavBar from "@/components/NavBar";
import { Inter } from 'next/font/google'
import { useBottomDrawerStore } from "@/stores/BottomDrawerStore";
import { useShallow } from "zustand/shallow";
import { HiOutlineIdentification, HiOutlineArrowRightOnRectangle, HiOutlineUserPlus, HiOutlineMegaphone } from "react-icons/hi2";
import { ColorModeButton, ColorModeProvider } from "@/components/ui/color-mode";
import { useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { getJwtPayload, JwtPayload } from "./api/getJwtPayload";
import { logout } from "./api/logout";

// If loading a variable font, you don't need to specify the font weight
const inter = Inter({
  variable: "--font-inter",
  subsets: ['latin'],
  display: 'swap',
})

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { drawerState, changeDrawerState } = useBottomDrawerStore(useShallow((state) => {
    return { drawerState: state.isOpen, changeDrawerState: state.change }
  }))

  const router = useRouter()
  const searchParams = useSearchParams()

  // https://nextjs.org/docs/app/api-reference/functions/use-search-params#updating-searchparams
  const createQueryString = useCallback(
    (name: string, value: string) => {
      const params = new URLSearchParams(searchParams.toString())
      params.set(name, value)

      return params.toString()
    },
    [searchParams]
  )

  const [jwtPayload, setJwtPayload] = useState<JwtPayload>({
    token: '',
    username: '',
    id: 0,
    ita: new Date(0)
  })

  useEffect(() => {
    const updatejwtPayload = async () => {
      const payload = await getJwtPayload()
      setJwtPayload(payload)
    }
    updatejwtPayload()
  }, [])

  return (
    <html className={inter.variable} lang="en" suppressHydrationWarning>
      <body>
        <Provider>
          <ColorModeProvider>
            <Grid paddingX={{ "2xl": 200, xl: 0 }} paddingX={3}
              templateAreas={{
                xl: `"nav nav"
                "main aside"`,
                mdTo2xl: `"nav nav"
                     "main aside"`,
                sm: `"nav" "main"`,
              }}
              templateColumns={{
                "2xl": "1fr 300px",
                xl: "1fr 300px",
                lg: "1fr 300px",
                sm: "1fr"
              }}>
              <GridItem area={"nav"}>
                <NavBar />
              </GridItem>
              <GridItem area={"main"} rowSpan={2} /*backgroundColor="blue.200"*/>
                {children}
              </GridItem>
              <GridItem area={"aside"} hideBelow="lg" /*backgroundColor="red.200"*/>
                Sidebar
              </GridItem>
            </Grid>
            <DrawerRoot placement={"bottom"} open={drawerState} onInteractOutside={changeDrawerState}>
              <DrawerBackdrop />
              <DrawerContent position={"fixed"} bottom={0} left={0} roundedTop={"2xl"}>
                <DrawerHeader hidden={jwtPayload.id == 0 ? true : false}>
                  <DrawerTitle marginBottom={-6}>{jwtPayload.id != 0 ? "Hello " + jwtPayload.username : ""}</DrawerTitle>
                </DrawerHeader>
                <DrawerBody marginInline={-6} marginBottom={-2}>
                  <Flex onClick={() => {
                    if (jwtPayload.id == 0) {
                      router.push("/login/" + "?" + createQueryString("origin", window.location.href.split(window.location.host)[1]))
                      changeDrawerState()
                    } else {
                      logout()
                      changeDrawerState()
                    }
                  }} roundedTop={"3xl"} direction={"row"} alignItems={"center"} gap={3} padding={5} paddingInline={6} _hover={{ bg: "gray.200" }} _active={{ bg: "gray.200" }}>
                    <Icon fontSize={"38px"} color={"gray.600"}>
                      {jwtPayload.id == 0 ? <HiOutlineIdentification /> : <HiOutlineArrowRightOnRectangle />}
                    </Icon>
                    {jwtPayload.id == 0 ? "Log in" : "Log out"}
                  </Flex>
                  <Flex hidden={jwtPayload.id != 0} onClick={() => {
                    router.push("/signup/" + "?" + createQueryString("origin", window.location.href.split(window.location.host)[1]))
                    changeDrawerState()
                  }
                  } direction={"row"} alignItems={"center"} gap={3} padding={5} paddingInline={6} paddingTop={1} _hover={{ bg: "gray.200" }} _active={{ bg: "gray.200" }}>
                    <Icon fontSize={"34px"} color={"gray.600"}>
                      <HiOutlineUserPlus />
                    </Icon>
                    <Text paddingStart={1}>Sign up</Text>
                  </Flex>
                  <Flex hidden={jwtPayload.id == 0} onClick={() => {
                    router.push("/c/create/" + "?" + createQueryString("origin", window.location.href.split(window.location.host)[1]))
                    changeDrawerState()
                  }
                  } direction={"row"} alignItems={"center"} gap={3} padding={5} paddingInline={6} paddingTop={1} _hover={{ bg: "gray.200" }} _active={{ bg: "gray.200" }}>
                    <Icon fontSize={"34px"} color={"gray.600"}>
                      <HiOutlineMegaphone />
                    </Icon>
                    <Text paddingStart={1}>Create new community</Text>
                  </Flex>
                  <Flex direction={"row"} alignItems={"center"} gap={3.5} padding={5} paddingInline={6} paddingTop={-1} _hover={{ bg: "gray.200" }}>
                    <ColorModeButton _icon={{ color: "gray.600" }} />
                    Change color mode
                  </Flex>
                </DrawerBody>
                <DrawerCloseTrigger />
              </DrawerContent>
            </DrawerRoot>
          </ColorModeProvider>
        </Provider>
      </body>
    </html >
  );
}
