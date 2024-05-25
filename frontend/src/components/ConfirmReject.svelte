<script lang="ts">
    import { _ } from 'svelte-i18n';

    export let confirm = true;
    export let callback;
</script>

<span on:click={() => callback(confirm)} class="bg-lightwood-100 text-darkwood-950 p-2 rounded-xl hover:scale-110 transition-all {confirm ? 'hover:bg-leaf-200' : 'hover:bg-lightwood-300'} cursor-pointer">
    {confirm ? $_("Confirm") : $_("Reject")}
</span>