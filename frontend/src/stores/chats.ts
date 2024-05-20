import { compareDate } from '@/utils/functions';
import { writable } from 'svelte/store';
import { connect, sendMsg } from "@/utils/ws"

const backend_url  = import.meta.env.VITE_BACKEND_URL;

interface Message {
    id: string;
    content: string;
    created_at: string;
    user_id: string;
    username: string;
    chat_id: string;
    type: number;
}

interface Chat {
    ID: string;
    user: string;
    lastMessage?: string;
    messages?: Message[];
}

interface ChatStore {
    chats: Chat[];
    currentChat: number;
    fetchChats: () => Promise<void>;
    fetchMessages: (chatID: string) => Promise<void>;
    addMessage: (chatID, msg) => void;
}

const chatStore = writable<ChatStore>({
    chats: [],
    currentChat: 0,
    fetchChats: async () => {
        try {
            const response = await fetch(`${backend_url}/chats`, {
                method: "POST",
                body: JSON.stringify({
                    uuid: localStorage.getItem("uuid")
                })
            });
            const newChatsResponse = await response.json();

            const newChats: Chat[] = newChatsResponse.map( chat => {
                return  {
                    ID: chat.chat_id,
                    user: chat.users.find(user => user.id !== localStorage.getItem("uuid")).username,
                    lastMessage: '',
                    messages: []
                }
            })

            chatStore.update(store => ({
                ...store,
                chats: newChats
            }));
        } catch (error) {
            console.error('Failed to fetch chats:', error);
        }
    },
    fetchMessages: async (chatID) => {
        try{
            const response = await fetch(`${backend_url}/messages/${chatID}`);
            const messages = await response.json();
            
            chatStore.update(store => {
                const chatIndex = store.chats.findIndex(chat => chat.ID === chatID);
                if (chatIndex !== -1) {
                    store.chats[chatIndex].messages = messages || [];
                }
                return { ...store };
            });

        } catch (error) {
            console.error('Failed to fetch chats:', error);
        }
    },
    addMessage: (chatID, msg) => {
        chatStore.update(store => {
            const chatIndex = store.chats.findIndex(chat => chat.ID === chatID);
            store.chats[chatIndex].messages.unshift(msg)
            return {...store}
        })
    }
});

export default chatStore;
