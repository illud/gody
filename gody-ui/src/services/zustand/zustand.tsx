import { create } from 'zustand';

// Define the shape of the store's state
interface TokenStore {
    token: string | null;
    setToken: (newToken: string) => void;
    clearToken: () => void;
}

// Create the Zustand store
export const useTokenStore = create<TokenStore>((set) => ({
    token: null, // initial state (no token by default)
    setToken: (newToken: string) => set({ token: newToken }), // method to set the token
    clearToken: () => set({ token: null }), // method to clear the token
}));
