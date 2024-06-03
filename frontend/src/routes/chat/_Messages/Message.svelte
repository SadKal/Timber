<script lang="ts">
    import { _ } from 'svelte-i18n';
    import ContextMenu from './_MessageContext/ContextMenu.svelte'
    
    export let message;
    export let chatID;

    const user = localStorage.getItem("user");
    const isUser: boolean = user === message.username

    const hour = new Date(message.created_at).getHours();
    const minute = new Date(message.created_at).getMinutes().toString().padStart(2, '0');

    let showContext = false;

    function handleContextMenu(event) {
        event.preventDefault();
        if (message.type != 5){
            showContext = true;
        }
    }

</script>

{#if isUser && showContext}
    <ContextMenu {message} {chatID} bind:showContext/>
{/if}

<div class="p-2 mb-1 min-w-min text-lg {isUser ? 'self-end mr-3 bg-leaf-500 text-light-50 items-end' : 'self-start ml-3 bg-lightwood-200 items-start'} rounded-lg flex flex-col min-w-24" on:contextmenu|preventDefault={handleContextMenu}>
    <div class="{isUser ? '' : ''}">
        {isUser ? $_("You") : message.username}
    </div>

    <div class="flex gap-2 justify-between">
        <div class="break-words max-w-3xl shrink {message.type == 5 ? 'italic' : ''}">
            {message.type == 5 ? $_("Deleted") : message.content}
        </div>
        {#if message.type == 7}
            <div class="font-light text-xs self-end italic">
                ({$_("Edited")})
            </div>
        {/if}
        <div class="font-light text-xs self-end">
            {hour}:{minute}
        </div>
    </div>
</div>

