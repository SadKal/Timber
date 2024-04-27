<script lang="ts">
    import { _ } from 'svelte-i18n';

    export let content: string = "";
    export let link: boolean = false;
    export let href: string = "";
    export let lightwood: boolean = false;
</script>



{#if link}
    <a href={href} class="{lightwood ? "bg-lightwood-100" : "bg-light-50"} w-3/4 text-center p-3 pl-5 pr-5 text-2xl rounded-xl transition-shadow hover:shadow-md hover:shadow-lightwood-700">{$_(content)}</a>
{:else}
    <button type="submit" class="bg-light-50 w-3/4 text-center p-3 text-2xl rounded-xl transition-shadow hover:shadow-md hover:shadow-lightwood-700">{$_(content)}</button>
{/if}
