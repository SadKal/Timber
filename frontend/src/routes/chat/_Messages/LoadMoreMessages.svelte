<script lang="ts">
    import chatStore from '@/stores/chats'
    import { _ } from 'svelte-i18n';

    export let chatID;

    let loading: boolean = false;
    $: moreToLoad = $chatStore.chats.find( (chat) => chat.ID == chatID)?.moreToLoad

    const loadMessages = async () => {
        loading = true
        try {
            const response = await $chatStore.fetchMessages(chatID);
            loading = false;
        } catch (e) {
            console.log("Error fetching messages", e);
        } 
    }
</script>

{#if !loading && moreToLoad}
    <div class="self-center m-10 p-2 rounded-lg bg-leaf-800 text-light-50 text-xl cursor-pointer transition-all hover:bg-leaf-600 hover:scale-105" on:click={loadMessages}>
        {$_("LoadMore")}
    </div>
{/if}