import { create } from 'zustand';

export const zUser = create((set) => ({
    isLoggedIn: false,
    setUserState: (state=false) => set({isLoggedIn: state}),
}));