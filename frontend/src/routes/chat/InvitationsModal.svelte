<script lang="ts">
    import { _ } from 'svelte-i18n';
    import { createChat } from '@/utils/chatHandler';
    import { getContext } from 'svelte';
    import type { Context } from 'svelte-simple-modal';
    import InvitationsList from './InvitationsList.svelte'
    const { open } = getContext<Context>('simple-modal');
    const showInvitationsList = () => open(InvitationsList);
</script>

<div on:click={showInvitationsList} on:keydown={createChat} class="flex items-center p-3 rounded-xl transition-all text-lightwood-200 text-lg cursor-pointer invitations_fluctuation">
    {$_("Invitations")}
</div>