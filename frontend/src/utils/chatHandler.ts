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

export async function getUsersByUsername(username: string): Promise<void> {
    try {
        if (username.trim() !== '') {
            const response = await fetch(`${backend_url}/users/${username}`);

            const users = await response.json();
            const usersWithUrl = await Promise.all(users.map(async (user) => {
                const url = await getImage(user.id);
                return {
                    url,
                    ...user
                };
            }));
            chatStore.update(store => ({
                ...store,
                usersResult: usersWithUrl
            }));
        }
    }
    catch (error) {
        return error;
    }
}

export async function getImage(imageName: string[]): Promise<string> {
    const response = await fetch(`${backend_url}/images/${imageName}`)
    if (response.ok) {
        const blob = await response.blob();

        var urlCreator = window.URL || window.webkitURL;
        const imageUrl = urlCreator.createObjectURL(blob);
        return imageUrl;
    } else {
        console.error('Failed to fetch image');
    }
}