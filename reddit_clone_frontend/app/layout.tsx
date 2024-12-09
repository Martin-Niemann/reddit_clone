import type { Metadata } from "next";
import { Provider } from "@/components/ui/provider"
import { Grid, GridItem } from "@chakra-ui/react";
import NavBar from "@/components/NavBar";

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <Provider>
          <Grid paddingX={{ "2xl": 200, xl: 0}}
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
            <GridItem area={"main"} rowSpan={2} backgroundColor="blue.200">
              {children}
            </GridItem>
            <GridItem area={"aside"} hideBelow="lg" backgroundColor="red.200">
              Sidebar
            </GridItem>
          </Grid>
        </Provider>
      </body>
    </html>
  );
}
