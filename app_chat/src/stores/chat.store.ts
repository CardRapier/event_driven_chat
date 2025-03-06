// store.js
import {create} from 'zustand';

const useChatStore = create((set) => ({
  messages: [],
  socket: null,

  addMessage: (msg) =>
    set((state) => ({
      messages: [...state.messages, msg],
    })),
  // You can also store the socket reference here if needed:
  setSocket: (socketInstance) => set({ socket: socketInstance }),
}));

export default useChatStore;
