import { writable } from 'svelte/store';
import { sendMsg } from "@/utils/ws"
import { v4 as uuidv4 } from 'uuid';

const backend_url  = import.meta.env.VITE_BACKEND_URL;

interface Message {
    id?: string;
    content: string;
    created_at?: Date;
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

interface Invitation {
    id: string;
    sender_username: string;
    sender: string;
    receiver: string;
    url: string;
}

interface ChatStore {
    chats: Chat[];
    invitations: Invitation[];
    currentChat: number;
    usersResult: any;
    fetchChats: () => Promise<void>;
    fetchMessages: (chatID: string) => Promise<void>;
    addMessage: (chatID, msg) => void;
    receiveMessage: (msg) => void;
    deleteInvitation: (invitation) => void;
}

const chatStore = writable<ChatStore>({
    chats: [],
    currentChat: 0,
    usersResult: [],
    invitations: [],
    fetchChats: async () => {
        try {
            const response = await fetch(`${backend_url}/chats`, {
                method: "POST",
                body: JSON.stringify({
                    uuid: localStorage.getItem("uuid")
                })
            });
            const newChatsResponse = await response.json();

            const newChats: Chat[] = newChatsResponse?.map( chat => {
                return  {
                    ID: chat.chat_id,
                    user: chat.users.find(user => user.id !== localStorage.getItem("uuid")).username,
                    lastMessage: '',
                    messages: []
                }
            }) || []

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
            console.error('Failed to fetch messages:', error);
        }
    },
    addMessage: (chatID, msg) => {
        const message = {
            id: uuidv4(),
            type: 0,
            content: msg,
            chat_id: chatID,
            user_id: localStorage.getItem("uuid"),
            username: localStorage.getItem("user"),
            created_at: new Date(Date.now()),
        }
        chatStore.update(store => {
            const chatIndex = store.chats.findIndex(chat => chat.ID === chatID);
            store.chats[chatIndex].messages.unshift(message)
            return {...store}
        })
        sendMsg(message)
    },
    receiveMessage: (msg) => {
        chatStore.update(store => {
            const message = JSON.parse(msg.data)
            console.log(message.chat_id)
            console.log(store)
            const chatIndex = store.chats.findIndex(chat => chat.ID === message.chat_id);
            store.chats[chatIndex].messages.unshift(message)
            return {...store}
        })
    },
    deleteInvitation: async (invitation) => {
        try {
            const response = await fetch(`${backend_url}/invitations/${invitation.id}`, {
                method: 'DELETE',
            });
    
            if (response.ok) {
                // Update the store after successful deletion
                chatStore.update(store => {
                    store.invitations = store.invitations.filter(inv => inv.id !== invitation.id);
                    return {...store};
                });
            } else {
                console.error('Failed to delete invitation');
            }
        } catch (error) {
            console.error('Error deleting invitation:', error);
        }
    }
});

export default chatStore;
