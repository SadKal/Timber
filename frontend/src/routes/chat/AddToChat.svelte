<script lang="ts">
    import Modal from 'svelte-simple-modal'
    import AddToChatModal from './AddToChatModal.svelte'
</script>

<div class="flex justify-center items-center gap-5 p-3">
    <Modal
    classContent="bg-lightwood-100"
    classWindowWrap="rounded-xl">
        <AddToChatModal />
    </Modal>
</div>
