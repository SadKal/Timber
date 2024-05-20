<script lang="ts">
    import { afterUpdate, onMount } from "svelte"
    import ChatDashboard from "./ChatDashboard.svelte"
    import ChatMain from "./ChatMain.svelte"
    import chatStore from '@/stores/chats'
    import { params } from '@roxi/routify';

    $: chatID = $params.chatID;
    $: chats = $chatStore.chats;

</script>

<div class="w-full h-screen bg-leaf-600 flex justify-between">
    <ChatMain {chatID}/>
    <ChatDashboard {chats} />
</div>

