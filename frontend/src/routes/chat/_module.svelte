<script lang="ts">
    import chatStore from "@/stores/chats";
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
        await $chatStore.fetchChats()
        await fetchMessages();
    })
</script>

<slot />