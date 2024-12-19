import { create } from 'zustand'

interface AddNewCommentState {
    isShown: boolean
    change: () => void
}

export const useAddNewCommentStore = create<AddNewCommentState>()((set) => ({
    isShown: false,
    change: () => set((state) => ({ isShown: !state.isShown }))
}))