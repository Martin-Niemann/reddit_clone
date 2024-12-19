export type trySendCommentDTO = {
    postId: number,
    parentId?: number,
    isFailed: boolean,
    success: boolean
}