import { Comment } from "@/types/Comment"

export interface Post {
    id: number;
    title: string;
    link: string;
    text: string;
    created_at: Date;
    updated_at: Date;
    username: string;
    score: number;
    comments_count: number;
    comments: Comment[];
    community_name: string;
}