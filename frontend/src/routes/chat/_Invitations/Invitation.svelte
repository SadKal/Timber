<script lang="ts">
    import ConfirmReject from '@/components/ConfirmReject.svelte'
    import chatStore from '@/stores/chats'
    import { createChat } from '@/utils/chatHandler'
    import { _ } from 'svelte-i18n';


    export let invitation;
    const {url, id, sender_username, sender, receiver } = invitation;

    const confirmOrReject = async (confirm) => {
        if (confirm){
            createChat(invitation)
            
        }else{
            $chatStore.deleteInvitation(invitation)
        }
    }

</script>

<div  class="flex justify-between items-center m-5 p-2 transition-all bg-leaf-600 hover:bg-leaf-500 rounded-xl">
    <div class="flex items-center gap-6">
        <img src={url} class="rounded-full object-cover w-20 h-20" alt="pfpicture" id={id}/>
        <span class="mr-5 text-lightwood-100 text-2xl">{sender_username}</span>
    </div>
    <div class="flex gap-3">
        <ConfirmReject callback={confirmOrReject} confirm={true} />
        <ConfirmReject callback={confirmOrReject} confirm={false} />
    </div>
</div>