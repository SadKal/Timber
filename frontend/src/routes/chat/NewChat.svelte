<script lang="ts">
    import TiPlus from 'svelte-icons/ti/TiPlus.svelte'
    import { _ } from 'svelte-i18n';
    import { createChat } from '@/utils/chatHandler';

</script>

<div on:click={createChat} on:keydown={createChat} class="bg-leaf-700 hover:bg-leaf-400 flex items-center justify-between gap-2 m-10 p-3 rounded-xl transition-all text-lightwood-200  hover:text-lightwood-900 text-lg cursor-pointer">
    <div class="w-5 h-5">
        <TiPlus />
    </div>
    {$_("NewChat")}
</div>