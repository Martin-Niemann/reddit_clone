export interface CommentDTO {
    post_id?: number;
    parent_id?: number;
    user_id: number;
    text: string;
}