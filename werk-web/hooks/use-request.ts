import { useState } from "react";
import { useAuth } from "../pages/context/AuthUserContext"

export interface ClientRequestParams {
    url: string;
    method: string;
    body: any;
}

export function useClientRequest() {
    const [waiting, setWaiting] = useState(true)

    const { getToken } = useAuth()

    const doRequest = async (url: string, method: string, body: any = undefined) => {
        setWaiting(true)

        const token = await getToken()

        const res = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify(body)
        })

        setWaiting(false)
        return res
    }

    return {
        doRequest,
        waiting,
    }
}
