'use server'

import { cookies } from 'next/headers';

export interface JwtPayload {
    "token": string
    "username": string
    "id": number
    "ita": Date
}

export async function getJwtPayload(): Promise<JwtPayload> {
    const cookieStore = await cookies()

    const token = cookieStore.get("auth")?.value

    if (token != undefined) {
        const payload: JwtPayload = JSON.parse(Buffer.from(token.split('.')[1], 'base64').toString());
        payload.token = token
        return payload
    } else {
        const payload: JwtPayload = {
            token: '',
            username: '',
            id: 0,
            ita: new Date(0)
        }
        return payload
    }
}
