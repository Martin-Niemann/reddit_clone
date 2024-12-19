'use server'

import { revalidatePath } from "next/cache"
import { cookies } from "next/headers"

export async function logout() {
    const cookieStore = await cookies()
    cookieStore.delete("auth")

    // TODO this doesn't revalidate the account icon after logout
    revalidatePath("/", "layout")
}