<script lang="ts">
    import TiPlus from 'svelte-icons/ti/TiPlus.svelte'
    import { _ } from 'svelte-i18n';
    import { createChat } from '@/utils/chatHandler';
    import { getContext } from 'svelte';
    import type { Context } from 'svelte-simple-modal';
    import NewChatUserList from './NewChatUserList.svelte'
    const { open, close } = getContext<Context>('simple-modal');
    const showUserList = () => open(NewChatUserList);
</script>

<div on:click={showUserList} on:keydown={createChat} class="bg-leaf-700 hover:bg-leaf-400 flex items-center justify-between gap-2 p-3 rounded-xl transition-all text-lightwood-200  hover:text-lightwood-900 text-lg cursor-pointer w-5/6">
    <div class="w-5 h-5">
        <TiPlus />
    </div>
    {$_("NewChat")}
</div>