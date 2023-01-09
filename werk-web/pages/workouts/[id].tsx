import { GetServerSidePropsContext } from "next"

interface Props {
    id: number
}

export default function WorkoutDetail({ id }: Props) {
    return (
        <div>{id}</div>
    )
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
    const id = context.params?.id

    return {
        props: {
            id
        }
    }
}
