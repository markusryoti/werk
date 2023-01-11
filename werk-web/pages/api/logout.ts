import { NextApiRequest, NextApiResponse } from "next";
import { setCookie } from "../../utils/cookies";

type ResponseBody = {
    message: string
}

export default function handler(_: NextApiRequest, res: NextApiResponse<ResponseBody>) {
    setCookie(res, "session", "", { expires: new Date(0) })

    res.status(200).json({ message: 'session destroyed' })
}
