<script lang="ts">
    import Modal from 'svelte-simple-modal'
    import NewChatModal from './NewChatModal.svelte'
    import InvitationsModal from './InvitationsModal.svelte'
    import chatStore from '@/stores/chats'
</script>

<div class="flex justify-center items-center gap-5 p-3">
    <Modal
    classContent="bg-lightwood-100"
    classWindowWrap="rounded-xl">
        <NewChatModal />
    </Modal>
    {#if $chatStore.invitations.length > 0}
        <Modal
        classContent="bg-lightwood-100"
        classWindowWrap="rounded-xl">
            <InvitationsModal />
        </Modal>
    {/if}
</div>
