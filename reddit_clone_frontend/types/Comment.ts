export interface Comment {
    id: number;
    parent_id?: number;
    created_at: Date;
    updated_at: Date;
    username: string;
    text: string;
    score: number;
}