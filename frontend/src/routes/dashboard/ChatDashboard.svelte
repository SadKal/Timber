<script>
    import ChatThumb from "./ChatThumb.svelte"

    export let chats;
</script>



<div class="w-96 h-full bg-darkwood-400 z-10">
    <div class="pt-16 overflow-scroll overflow-x-hidden max-h-full">
        {#each chats as { user, lastMessage } (user) }
            <ChatThumb {user} {lastMessage} />
        {/each}
    </div>
</div>