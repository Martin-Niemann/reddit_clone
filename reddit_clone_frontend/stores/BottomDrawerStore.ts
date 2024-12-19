import { create } from 'zustand'

interface BottomDrawerState {
    isOpen: boolean
    change: () => void
}

export const useBottomDrawerStore = create<BottomDrawerState>()((set) => ({
    isOpen: false,
    change: () => set((state) => ({ isOpen: !state.isOpen }))
}))