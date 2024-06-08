<script lang="ts">
    import { afterUpdate } from "svelte"
    import Message from "./_Messages/Message.svelte"
    import { _ } from 'svelte-i18n';
    import chatStore from "@/stores/chats"
    import LoadMoreMessages from "./_Messages/LoadMoreMessages.svelte"

    export let chatID: string = '';

    $: chatWith = $chatStore.chats.find(chat => chat.ID === chatID);
    $: messages = $chatStore.chats.find(chat => chat.ID === chatID)?.messages || [];
    let container: HTMLElement;
    let messageBox: HTMLElement;

    afterUpdate(scrollToBottom);

    function scrollToBottom() {
        const currentScroll = chatWith?.currentScroll ? chatWith?.currentScroll : container.scrollTop;
        if (container) {
            container.scrollTop = currentScroll;
        }
    }

    function handleMessage(event) {
        if (event.key === "Enter"){
            event.preventDefault();
            sendMessage()
        }
    }

    function sendMessage(){
        if (messageBox.textContent !== ""){
            $chatStore.addMessage(chatID, messageBox.textContent, 0)
            messageBox.textContent = ""
        }
    }

</script>

<div class="bg-darkwood-700 w-full flex flex-col justify-between relative" style="background-image: linear-gradient(to bottom, rgba(175, 132, 71, 0.85), rgba(175, 132, 71, 0.85)), url('/assets/backgrounds/wood_texture.webp');">
    <div class="top-0 bg-leaf-700 w-full pl-5 h-20 shadow-3xl shrink-0 flex items-center justify-between">
        <div class="flex items-center">
            <img class="rounded-full object-cover w-16 h-16" src={chatWith?.pfp} alt="profile_picture"/>
            <span class="text-lightwood-100 text-4xl pl-10">{$_("ChatWith") + chatWith?.user}</span>
        </div>
    </div>
    <div bind:this={container}  class="flex flex-col-reverse overflow-scroll overflow-x-hidden grow px-6 py-3">
        {#each messages as message (message.id)}
            <Message {chatID} {message}/>
        {/each}
        <LoadMoreMessages {chatID}/>
    </div>
    <div class="bg-leaf-500 w-full h-30 shrink-0 py-3 pl-3 flex justify-center">
        <div contenteditable class="bg-lightwood-100 border-leaf-900 chatbox " bind:this={messageBox} on:keydown={handleMessage}/>
        <img class="p-3 hover:cursor-pointer hover:scale-110 hover:drop-shadow-xl transition-all" alt="sending button" src="/assets/components/leaf-send.svg" width="50" height="50" on:click={sendMessage}/>
    </div>
</div>