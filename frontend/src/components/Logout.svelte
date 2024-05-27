<script>
    import { _ } from 'svelte-i18n';

    function logout () {
        document.cookie = "jwt=;expires=Thu, 01 Jan 1970 00:00:00 GMT;path=/";
        document.location.href = "/login";
    }
</script>

<div class="flex items-center gap-3 bg-lightwood-100 text-darkwood-950 rounded-md p-1.5 text-sm cursor-pointer transition-all hover:scale-105" on:click={logout}>
    <img class="h-5 w-5" alt="sending button" src="/assets/components/logout.svg"/>
    {$_("Logout")}
</div>