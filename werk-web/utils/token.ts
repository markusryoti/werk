import { parse } from "cookie";

export default function getTokenFromCookie(reqCookie: string) {
    return parse(reqCookie)
}
