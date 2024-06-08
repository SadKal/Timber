<script lang="ts">
    import chatStore from '@/stores/chats'
    import { editMessage } from '@/utils/chatHandler'
    import { getContext } from 'svelte';
    import type { Context } from 'svelte-simple-modal';
    const { close } = getContext<Context>('simple-modal');
    import { _ } from 'svelte-i18n';

    export let message

    let editedMessage = message.content;

    function resizeTextarea(event) {
        const textarea = event.target;
        textarea.style.height = 'auto';
        textarea.style.height = (textarea.scrollHeight) + 'px';
    }

    function removeNewlines(message) {
        message = message.replace(/\n/g, '');
        return message
    }

    const handleEdit = async () => {
        editedMessage = removeNewlines(editedMessage)
        const infoToEdit = {
            id: message.id,
            chat_id: message.chat_id,
            content: editedMessage
        }
        $chatStore.editMessage(infoToEdit);
        editMessage(infoToEdit)
        close()
    }
</script>

<div class="text-center text-darkwood-950">
    <p>{$_("EditTheMessage")}</p>
    <form class="flex flex-col items-center gap-4 p-2" on:submit|preventDefault={handleEdit}>
        <textarea on:input={resizeTextarea} class="box-border resize-none p-3 min-h-28 w-3/4 " bind:value={editedMessage} />
        <input class="bg-leaf-600 text-light-50 p-2 rounded-xl cursor-pointer transition-all hover:scale-105 hover:bg-leaf-400"  type="submit" value={$_("Submit")}/>
    </form>
</div>