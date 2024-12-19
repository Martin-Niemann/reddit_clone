import { Post } from "./Post";

export interface Community {
    id: number;
    name: string;
    description: string;
    moderator: string;
    posts: Post[];
}