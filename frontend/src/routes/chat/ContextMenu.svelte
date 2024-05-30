<script lang="ts">
    import { _ } from 'svelte-i18n';
    import { deleteMessage } from "@/utils/chatHandler"
    import Bin from "./Bin.svelte"
    import chatStore from "@/stores/chats"
    import EditMessageModal from "./EditMessageModal.svelte"
    import { Modal } from "svelte-simple-modal"

    export let message;
    export let showContext;
    export let chatID;

    const id = message.id

    let posX = 0;
    let posY = 0;

    function handleClickOutside(event) {
        if (!event.target.closest(`.context-menu-${id}`)) {
            showContext = false;
        }
    }

    function handleContextMenu(event) {
        event.preventDefault();
        showContext = true;
        posX = event.clientX;
        posY = event.clientY;
    }

    const handleDeletion =  async () => {
        deleteMessage(id);
        $chatStore.addMessage(chatID, id, 4)
        showContext = false;
    }

    
</script>

<div class="context-menu-{id} bg-lightwood-200 p-2 rounded-md" style="position: fixed; left: {posX-100}px; top: {posY}px;">
    <Modal
    classContent="bg-lightwood-100"
    classWindowWrap="rounded-xl">
        <EditMessageModal {message}/>
    </Modal>

    <div class="flex p-2 cursor-pointer rounded-lg transition-all text-darkwood-950 hover:bg-lightwood-100" on:click={handleDeletion}>
        <Bin/>
        <span class="pl-2">{$_("DeleteMessage")}</span>
    </div>
</div>

<svelte:window on:click={handleClickOutside} on:contextmenu|preventDefault={handleContextMenu}/>
