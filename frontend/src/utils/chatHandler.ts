import chatStore from "@/stores/chats";

const backend_url  = import.meta.env.VITE_BACKEND_URL;

export async function createChat(){
    try {
        const response = await fetch(`${backend_url}/createchat`, {
            method: 'POST',
        });

        chatStore.update(store => {
            store.fetchChats();
            return store;
        });
    } catch (error) {
        return error;
    }
}