'use client'

import { Button, Icon } from "@chakra-ui/react"
import { redirect } from "next/navigation"
import { HiOutlinePencil } from "react-icons/hi2"

// https://nextjs.org/docs/app/api-reference/functions/use-search-params#updating-searchparams
function createQueryString(name: string, value: string) {
    const params = new URLSearchParams()
    params.set(name, value)

    return params.toString()
}

export default function AddComment({ params }: { params: { communityId: number } }) {
    return (
        <Button onClick={() => redirect("/p/create" + "?" + createQueryString("community", params.communityId.toString()))} rounded={"full"} color={"gray.700"} size={"sm"} marginTop={-1} variant={"outline"}>
            <Icon fontSize={"16px"} color={"gray.700"}>
                <HiOutlinePencil />
            </Icon>Create post
        </Button>
    )
}