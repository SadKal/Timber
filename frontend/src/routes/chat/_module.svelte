<script lang="ts">
    import chatStore from "@/stores/chats";
    import { connect, sendMsg } from "@/utils/ws"
    import { onMount } from "svelte"
    import { params } from '@roxi/routify';

    $: chatID = $params.chatID;
    $: chats = $chatStore.chats;
    let previousChatID = null;


    $: if ((chatID !== previousChatID) && chats != null) {
        fetchMessages();
        previousChatID = chatID;
    }

    async function fetchMessages() {
        if (chatID) {
            await $chatStore.fetchMessages(chatID);
        }
    }
    onMount(async () => {
        connect($chatStore.receiveMessage);
        await $chatStore.fetchChats()
        await fetchMessages();
    })
</script>

<slot />