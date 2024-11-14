import { cache } from "react"

export default async function Page({ params }: { params: Promise<{ slug: string }> }) {
  const slug = (await params).slug
  const data = await getPosts(slug)
  console.log(data)
  return <p>My Post: {data.description}</p>
}

const getPosts = cache(async (slug: string) => {
  const res = await fetch(`http://127.0.0.1:8090/subreddit/${slug}?sort_by=newest`)
  return res.json()
})