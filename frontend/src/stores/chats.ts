import { writable } from 'svelte/store';
import { sendMsg } from "@/utils/ws"
import { v4 as uuidv4 } from 'uuid';
import { getImage, getInvitations } from '@/utils/chatHandler';

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
    lastMessage?: Message
    messages?: Message[];
    cache?: boolean;
    pfp?: string;
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
    fetchChatByID: (chatID) => Promise<any>;
    fetchMessages: (chatID: string) => Promise<void>;
    addMessage: (chatID, msg, type) => void;
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

            const newChats: Chat[] = await Promise.all(newChatsResponse?.map(async chat => {
                const otherUser = chat.users.find(user => user.id !== localStorage.getItem("uuid"))

                const otherUserPfp = await getImage(otherUser.id);
                return  {
                    ID: chat.chat_id,
                    user: otherUser.username,
                    lastMessage: '',
                    messages: [],
                    cache: false,
                    pfp: otherUserPfp
                }
            }) || [])

            chatStore.update(store => ({
                ...store,
                chats: newChats
            }));
        } catch (error) {
            console.error('Failed to fetch chats:', error);
        }
    },
    fetchChatByID: async (chatID) => {
        try {
            const response = await  fetch(`${backend_url}/chats/${chatID}`);

            const chats = await response.json();
            const chat = chats[0]

            const otherUser = chat.users.find(user => user.id !== localStorage.getItem("uuid"))
            const otherUserPfp = await getImage(otherUser.id);
            const newChat: Chat = {
                ID: chat.chat_id,
                user: otherUser.username,
                lastMessage: null,
                messages: [],
                cache: false,
                pfp: otherUserPfp
            }

            chatStore.update(store => {
                store.chats.push(newChat)
                return {...store}
            });
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
                    store.chats[chatIndex].cache = true;
                }
                return { ...store };
            });

        } catch (error) {
            console.error('Failed to fetch messages:', error);
        }
    },
    addMessage: (chatID, msg, type = 0) => {
        const message = {
            id: uuidv4(),
            type: type,
            content: msg,
            chat_id: chatID == null ? uuidv4() : chatID,
            user_id: localStorage.getItem("uuid"),
            username: localStorage.getItem("user"),
            created_at: new Date(Date.now()),
        }
        switch (type){
            case 0:
                chatStore.update(store => {
                    const chatIndex = store.chats.findIndex(chat => chat.ID === chatID);
                    store.chats[chatIndex].messages.unshift(message)
                    return {...store}
                })
                break;
            case 1:
            case 3:
            default:
                break;
        }

        sendMsg(message)
    },
    receiveMessage: (msg) => {
        chatStore.update(store => {
            const message = JSON.parse(msg.data)

            switch (message.type) {
                case 0:
                    const chatIndex = store.chats.findIndex(chat => chat.ID === message.chat_id);
                    store.chats[chatIndex].messages.unshift(message)
                    store.chats[chatIndex].lastMessage = message
                    break;
                case 1:
                    getInvitations();
                    break;
                case 3:
                    store.fetchChatByID(message.content)
                    store.fetchMessages(store.currentChat.toString())
                    break;
            }
            return {...store}
        })
    },
    deleteInvitation: async (invitation) => {
        try {
            const response = await fetch(`${backend_url}/invitations/${invitation.id}`, {
                method: 'DELETE',
            });

            if (response.ok) {
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
