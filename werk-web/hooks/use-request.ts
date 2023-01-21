import { useAuth } from "../pages/context/AuthUserContext"

export interface ClientRequestParams {
    url: string;
    method: string;
    body: any;
}

export function useClientRequest() {
    const { getToken } = useAuth()

    const doRequest = async (url: string, method: string, body: any) => {
        const token = await getToken()

        return await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }, body: JSON.stringify(body)
        })
    }

    return {
        doRequest
    }
}
