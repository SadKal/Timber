<script lang="ts">
    import chatStore from '@/stores/chats'
import { getUsersByUsername } from '@/utils/chatHandler'
    import { debounce } from '@/utils/functions'
    import { _ } from 'svelte-i18n';
    import UserResult from './UserResult.svelte'

    let userToSearch = '';

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

<div class="flex flex-col items-center text-darkwood-950">
    <span>
        {$_("SearchUser")}
    </span>
    <div>
        <input bind:value={userToSearch} type="text" />
    </div>
    <div class="w-1/2">
        {#each usersResult as {id, url, username} (id)}
            <UserResult {id} {url} {username} />
        {/each}
    </div>
</div>