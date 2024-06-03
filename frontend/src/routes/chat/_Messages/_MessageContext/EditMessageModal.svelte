<script lang="ts">
    import { _ } from 'svelte-i18n';
    import EditMessageForm from "./EditMessageForm.svelte"
    import Pen from "@/components/svg/Pen.svelte"
    import { getContext } from 'svelte'
    import type { Context } from 'svelte-simple-modal';
    const { open } = getContext<Context>('simple-modal');

    export let message;
    const handleEdition = async () => {
        open(EditMessageForm, {message: message});
    }
</script>



<div class="flex p-2 cursor-pointer rounded-lg transition-all text-darkwood-950 hover:bg-lightwood-100" on:click={handleEdition}>
    <Pen/>
    <span class="pl-2">{$_("EditMessage")}</span>
</div>