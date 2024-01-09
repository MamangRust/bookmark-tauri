import { create } from 'zustand';

export const useAuthStore = create((set) => ({
  token: localStorage.getItem('token') || '',
  setToken: (newToken) => {
    set({ token: newToken });
    localStorage.setItem('token', newToken);
  },
}));
