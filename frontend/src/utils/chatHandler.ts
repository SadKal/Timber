import chatStore from "@/stores/chats";

const backend_url  = import.meta.env.VITE_BACKEND_URL;
const user_uuid = localStorage.getItem('uuid');
const user_username = localStorage.getItem('user');


export async function createChat(invitation){
    try {
        console.log(invitation)
        const response = await fetch(`${backend_url}/createchat`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(invitation)
        });

        const chat = await response.json();
        chatStore.update(store => {
            store.deleteInvitation(invitation);
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

export async function sendInvitation(id: string) {
    const response = await fetch(`${backend_url}/invitations`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            sender_username: user_username, 
            sender: user_uuid,
            receiver: id
        })
    })
    if (!response.ok) {
        return true;
    }
    return false;
}

export async function getInvitations() {
    const response = await fetch(`${backend_url}/invitations/${user_uuid}`)

    const invitations = await response.json()
    const invitationsWithUrl = await Promise.all(invitations.map(async (invitation) => {
        const url = await getImage(invitation.sender);
        return {
            url,
            ...invitation
        };
    }));
    chatStore.update(store => ({
        ...store,
        invitations: invitationsWithUrl
    }));
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