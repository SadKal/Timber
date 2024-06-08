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
    ID?: string;
    user?: string;
    lastMessage?: Message
    messages?: Message[];
    cache?: boolean;
    pfp?: string;
    offset?: number;
    moreToLoad?: boolean;
    currentScroll?: number;
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
    scroll?: boolean;
    fetchChats: () => Promise<void>;
    fetchChatByID: (chatID) => Promise<any>;
    fetchMessages: (chatID: string) => Promise<string>;
    addMessage: (chatID, msg, type) => void;
    receiveMessage: (msg) => void;
    deleteInvitation: (invitation) => void;
    editMessage: (infoToEdit) => Promise<void>;
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
                    pfp: otherUserPfp,
                    offset: 0,
                    moreToLoad: true
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
                pfp: otherUserPfp,
                offset: 0,
                moreToLoad: true,
                currentScroll: 0
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
            let currentChat: Chat = {}
            const unsuscribe = chatStore.subscribe((store) => {
                currentChat = store.chats.find(chat => chat.ID === chatID)
            })
            unsuscribe();
            const offset = currentChat?.offset ? currentChat.offset : 0;

            const response = await fetch(`${backend_url}/messages/${chatID}?offset=${currentChat?.offset ? currentChat?.offset : 0}`);
            const messages = await response.json();

            chatStore.update(store => {
                const chatIndex = store.chats.findIndex(chat => chat.ID === chatID);
                const filteredMessages = messages.filter(message => message.type != 4)
                if (chatIndex !== -1) {
                    store.chats[chatIndex].messages = [
                        ...store.chats[chatIndex].messages,
                        ...filteredMessages
                    ];
                    store.chats[chatIndex].cache = true;
                    store.chats[chatIndex].offset = offset + messages.length;
                    store.chats[chatIndex].moreToLoad = messages.length < 20 ? false : true;
                }
                return { ...store };
            });

            return "Done";
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
                    store.chats[chatIndex].offset += 1;
                    return {...store}
                })
                break;
            case 1:
            case 3:
                break;
            case 4:
                chatStore.update(store => {
                    store.chats.find(chat => chat.ID === message.chat_id)
                    .messages.find(msg => msg.id === message.content)
                    .type = 5;
                    return {...store}
                })
            default:
                break;
        }

        sendMsg(message)
    },
    receiveMessage: (msg) => {
        chatStore.update(store => {
            const message = JSON.parse(msg.data)
            let chatIndex;
            chatIndex = store.chats.findIndex(chat => chat.ID === message.chat_id);
            switch (message.type) {
                case 0:
                    store.chats[chatIndex].messages.unshift(message)
                    store.chats[chatIndex].lastMessage = message
                    store.chats[chatIndex].offset += 1
                    break;
                case 1:
                    getInvitations();
                    break;
                case 3:
                    store.fetchChatByID(message.content)
                    store.fetchMessages(store.currentChat.toString())
                    break;
                case 4:
                    console.log(message)
                    store.chats[chatIndex].messages.find(msg => msg.id === message.content).type = 5;
                    break;
                case 6:
                    const infoToEdit = JSON.parse(message.content)
                    let messageReceived
                    messageReceived = store.chats
                    .find(chat => chat.ID === infoToEdit.chat_id)
                    .messages.find((message) => message.id === infoToEdit.id)

                    messageReceived.content = infoToEdit.content
                    messageReceived.type = 7
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
    },
    editMessage: async (infoToEdit) =>{
        let message
        chatStore.update((store) => {
            message = store.chats
            .find(chat => chat.ID === infoToEdit.chat_id)
            .messages.find((message) => message.id === infoToEdit.id)

            message.content = infoToEdit.content
            message.type = 7

            store.addMessage(infoToEdit.chat_id, JSON.stringify(infoToEdit), 6)
            return {...store}
        })
    }
});

export default chatStore;
