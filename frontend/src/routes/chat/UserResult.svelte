<script lang="ts">
    import { sendInvitation } from '@/utils/chatHandler';
    import { getContext } from 'svelte';
    import type { Context } from 'svelte-simple-modal';
    const { close } = getContext<Context>('simple-modal');

    export let id;
    export let username;
    export let url;
    export let invitationAlreadyExists;

    const sendUserInvitation = async () => {
        const error = await sendInvitation(id);
        if (error){
            invitationAlreadyExists = true;
        } else{
            invitationAlreadyExists = false;
            setTimeout(() => {
                close()
            }, 500)
        }
    }
</script>

<div on:click={sendUserInvitation}  class="flex justify-between items-center cursor-pointer m-5 p-2 bg-leaf-600 rounded-xl">
    <img src={url} class="rounded-full object-cover w-20 h-20" alt="pfpicture" id={id}/>
    <p class="mr-5 text-lightwood-100 text-2xl">{username}</p>
</div>