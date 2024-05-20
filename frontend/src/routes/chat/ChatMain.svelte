<script lang="ts">
    import { connect, sendMsg } from "@/utils/ws"
    import { afterUpdate, onMount } from "svelte"
    import Message from "./Message.svelte"
    import { _ } from 'svelte-i18n';
    import chatStore from "@/stores/chats"
    import { login } from "@/utils/auth"

    export let chatID: string = '';

    $: chatWith = $chatStore.chats.find(chat => chat.ID === chatID)?.user;
    $: messages = $chatStore.chats.find(chat => chat.ID === chatID)?.messages || [];


    function update(msg) {
        $chatStore.addMessage(chatID, JSON.parse(msg.data))
    }
    let container: HTMLElement;
    let messageBox: HTMLElement;

    afterUpdate(scrollToBottom);

    function scrollToBottom() {
        if (container) {
            container.scrollTop = container.scrollHeight;
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
            sendMsg(messageBox.textContent, chatID)
            messageBox.textContent = ""
        }
    }

    onMount(() => {
        connect(update);
    })
</script>



<div class="bg-darkwood-700 w-full flex flex-col justify-between relative" style="background-image: linear-gradient(
    to bottom,
    rgba(175, 132, 71, 0.85),
    rgba(175, 132, 71, 0.85)
    ), url('/assets/backgrounds/wood_texture.webp');">
    <div class="top-0 bg-leaf-700 w-full h-20 shadow-3xl shrink-0 flex items-center justify-between">
        <span class="text-lightwood-100 text-4xl pl-10">{$_("ChatWith") + chatWith}</span>
    </div>
    <div bind:this={container}  class="flex flex-col-reverse overflow-scroll overflow-x-hidden grow px-6 py-3">
        {#each messages as message (message.id)}
            <Message {message}/>
        {/each}
    </div>
    <div class="bg-leaf-500 w-full h-30 shrink-0 py-3 pl-3 flex justify-center">
        <div contenteditable class="bg-lightwood-100 border-leaf-900 chatbox " bind:this={messageBox} on:keydown={handleMessage}/>
        <img class="p-3 hover:cursor-pointer hover:scale-110 hover:drop-shadow-xl transition-all" alt="sending button" src="/assets/components/leaf-send.svg" width="50" height="50" on:click={sendMessage}/>
    </div>
</div>