<script lang="ts">
    import { connect, sendMsg } from "@/utils/ws"
    import { afterUpdate, onMount } from "svelte"
    import { writable } from 'svelte/store';
    import Message from "./Message.svelte"

    const messages = writable([]);

    function update(msg) {
        messages.update(currentMessages => [JSON.parse(msg.data), ...currentMessages]);
    }

    $: console.log($messages)

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
            sendMsg(messageBox.textContent)
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
    <div class="top-0 bg-leaf-700 w-full h-20 shadow-3xl shrink-0 " on:click={() => sendMsg("The curious cat chased the elusive butterfly through the sun-dappled meadow.")}>
        Hola header
    </div>
    <div bind:this={container}  class="flex flex-col-reverse overflow-scroll overflow-x-hidden grow">
        {#each $messages as message (message.id)}
            <Message {message}/>
        {/each}
    </div>
    <div class="bg-leaf-500 w-full h-30 shrink-0 p-3 flex justify-center">
        <div contenteditable class="bg-lightwood-100 border-leaf-900 chatbox " bind:this={messageBox} on:keydown={handleMessage}/>
    </div>
</div>