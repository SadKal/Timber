<script lang="ts">
    import ChatDashboard from "./ChatDashboard.svelte"
    import ChatMain from "./ChatMain.svelte"
    import chatStore from '@/stores/chats'
    import { params } from '@roxi/routify';

    $: chatID = $params.chatID;
    $: chats = $chatStore.chats;

    $: $chatStore.currentChat = chatID

</script>

<div class="w-full h-screen bg-leaf-600 flex justify-between">
    <ChatMain {chatID}/>
    <ChatDashboard {chats} />
</div>

