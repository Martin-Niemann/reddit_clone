import CommentCard from "@/components/CommentCard"
import PostBanner from "@/components/PostBanner"
import { Post } from "@/types/Post"
import { Box, Button, Collapsible, Flex, List, ListItem, Textarea } from "@chakra-ui/react"
import { Metadata, ResolvingMetadata } from "next"
import { cache } from "react"
import akademi from "@/public/soroe_akademi.avif"
import AddComment from "@/components/AddComment"

export async function generateMetadata(
    { params }: { params: Promise<{ slug: string }> },
    parent: ResolvingMetadata
): Promise<Metadata> {
    const slug = (await params).slug
    const post: Post = await getPost(slug)

    return {
        openGraph: {
            title: post.title,
            type: "article",
            images: akademi.src,
            url: "https://opinit.vercel.app/p/" + post.id,
            siteName: "Opinit",
            description: post.text != undefined ? post.text.substring(0, 140) + "..." : undefined,
            authors: post.username,
            publishedTime: new Date(post.created_at).toDateString(),
            modifiedTime: new Date(post.updated_at).toDateString(),
            section: "c/" + post.community_name
        }
    }
}

//https://stackoverflow.com/a/64093016
function partition(array: [], predicate) {
    return array.reduce((acc, item) => predicate(item)
        ? (acc[0].push(item), acc)
        : (acc[1].push(item), acc), [[], []]);
}

export default async function Page({ params }: { params: Promise<{ slug: string }> }) {
    const slug = (await params).slug
    const post: Post = await getPost(slug)
    console.log(post)

    let topLevelComments = post.comments.filter(comment => comment.parent_id == null)
    topLevelComments = topLevelComments.map(comment => {
        return { ...comment, post_id: post.id }
    })

    function checkSubLevel(id: number) {
        let subLevelComments = post.comments.filter(comment => comment.parent_id == id)
        subLevelComments = subLevelComments.map(comment => {
            return { ...comment, post_id: post.id }
        })

        if (subLevelComments != undefined) {
            return subLevelComments.map(comment =>
                <List.Item key={comment.id}>
                    <Flex>
                        <Box backgroundColor={"gray.300"} width={0.49} marginRight={4} marginBlock={3}></Box>
                        <CommentCard props={comment} />
                    </Flex>
                    <List.Root ps={5} listStyle={"none"}>
                        {checkSubLevel(comment.id)}
                    </List.Root>
                </List.Item>
            )
        } else {
            return <></>
        }
    }

    const commentsList = topLevelComments.map(comment =>
        <List.Item key={comment.id}>
            <CommentCard props={comment} />
            <List.Root ps={3} listStyle={"none"}>
                {checkSubLevel(comment.id)}
            </List.Root>
        </List.Item>
    )

    return (
        <>
            <PostBanner props={post} />

            <AddComment params={{ postId: post.id }} />

            <List.Root paddingTop={4} listStyle={"none"}>
                {commentsList}
            </List.Root>
        </>
    )
}

const getPost = cache(async (slug: string) => {
    const res = await fetch(process.env.BACKEND_HOST + `/post/${slug}`, { next: { tags: ["p"] } })
    return res.json()
})