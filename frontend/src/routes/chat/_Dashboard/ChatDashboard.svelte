<script>
    import ChatThumb from "./ChatThumb.svelte"
    import NewChat from "../_NewChat/NewChat.svelte"

    export let chats;
</script>



<div class="w-96 h-full pt-14 bg-darkwood-400">
    <NewChat />
    <div class="overflow-scroll overflow-x-hidden h-5/6">
        {#each chats as { ID, user, lastMessage, pfp } (ID) }
            <ChatThumb {ID} {user} {lastMessage} {pfp}/>
        {/each}
    </div>
</div>

