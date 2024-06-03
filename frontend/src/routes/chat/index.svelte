<script lang="ts">
    import ChatDashboard from "./_Dashboard/ChatDashboard.svelte"
    import ChatLanding from "./ChatLanding.svelte"
    import chatStore from '@/stores/chats'


    $: chats = $chatStore.chats
</script>


<div class="w-full h-screen bg-leaf-600 flex justify-between">
    <ChatLanding />
    <ChatDashboard {chats} />
</div>