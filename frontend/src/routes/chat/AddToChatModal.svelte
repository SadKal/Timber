<script lang="ts">
    import TiPlus from 'svelte-icons/ti/TiPlus.svelte'
    import { _ } from 'svelte-i18n';
    import { createChat } from '@/utils/chatHandler';
    import { getContext } from 'svelte';
    import type { Context } from 'svelte-simple-modal';
    import NewChatUserList from './NewChatUserList.svelte'
    const { open } = getContext<Context>('simple-modal');
    const showUserList = () => open(NewChatUserList, {newChat: false});
</script>

<div on:click={showUserList} on:keydown={createChat} class="bg-lightwood-100 hover:bg-leaf-400 flex items-center justify-between gap-2 p-3 rounded-xl transition-all text-lightwood-900  hover:text-lightwood-50 cursor-pointer w-52">
    <div class="w-5 h-5">
        <TiPlus />
    </div>
    {$_("AddToChat")}
</div>