import { NextApiRequest, NextApiResponse } from "next";
import { setCookie } from "../../utils/cookies";

type ResponseBody = {
    message: string
}

export default function handler(req: NextApiRequest, res: NextApiResponse<ResponseBody>) {
    const token = JSON.parse(req.body).token

    if (!token) {
        return res.status(400).json({ message: 'must include access token' })
    }

    setCookie(res, "session", token, { path: '/', maxAge: 30 * 24 * 60, sameSite: 'lax' })

    res.status(200).json({ message: 'cookie created' })
}
