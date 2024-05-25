<script lang="ts">
    import chatStore from '@/stores/chats'
    import { _ } from 'svelte-i18n';
    import Invitation from './Invitation.svelte'

    $: invitations = $chatStore.invitations
</script>

<div class="flex flex-col gap-5 items-center text-darkwood-950">
    <span class="text-2xl">
        {$_("Invitations")}
    </span>
    <div class="w-5/6">
        {#each invitations as invitation (invitation.sender)}
            <Invitation {invitation} />
        {/each}
    </div>
</div>