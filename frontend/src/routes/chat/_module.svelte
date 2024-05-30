<script lang="ts">
    import chatStore from "@/stores/chats";
    import { connect, sendMsg } from "@/utils/ws"
    import { onMount } from "svelte"
    import { params } from '@roxi/routify';
    import { getImage, getInvitations } from "@/utils/chatHandler"

    $: chatID = $params.chatID;
    $: chats = $chatStore.chats;
    let previousChatID = null;


    $: if ((chatID !== previousChatID) && chats != null) {
        fetchMessages();
        previousChatID = chatID;
    }

    async function fetchMessages() {
        const isCached = $chatStore.chats.find(chat => chat.ID == chatID)?.cache;
        if (chatID && !isCached) {
            console.log(chatID)
            await $chatStore.fetchMessages(chatID);
        }
    }
    onMount(async () => {
        connect($chatStore.receiveMessage);
        await getInvitations();
        await $chatStore.fetchChats()
        if (chatID){
            await fetchMessages();
        } 
        console.log($chatStore.chats)
    })
</script>


<slot />