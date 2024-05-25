<script lang="ts">
    import chatStore from '@/stores/chats'
import { getUsersByUsername } from '@/utils/chatHandler'
    import { debounce } from '@/utils/functions'
    import { _ } from 'svelte-i18n';
    import UserResult from './UserResult.svelte'

    let userToSearch = '';
    let invitationAlreadyExists = false;

    const searchUsers = debounce(async (username) => getUsersByUsername(username), 500)

    $: userToSearch && searchUsers(userToSearch)

    $: if (userToSearch.trim() !== '') {
        searchUsers(userToSearch);
    } else {
        searchUsers.cancel();
        $chatStore.usersResult = [];
    }

    $: usersResult = $chatStore.usersResult
</script>

<div class="flex flex-col gap-5 items-center text-darkwood-950">
    <span>
        {$_("SearchUser")}
    </span>
    <div>
        <input class="p-1" bind:value={userToSearch} type="text" />
    </div>
    {#if invitationAlreadyExists}
        <div>
            LA INVITACION YA EXISTE
        </div>
    {/if}
    <div class="w-1/2">
        {#each usersResult as {id, url, username} (id)}
            {#if id != localStorage.getItem("uuid")}
                <UserResult {id} {url} {username} bind:invitationAlreadyExists />
            {/if}
        {/each}
    </div>
</div>