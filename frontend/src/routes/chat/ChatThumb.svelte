<script lang="ts">
    import { params } from '@roxi/routify';
    import { _ } from 'svelte-i18n';

    $: ({ chatID }= $params)
    export let ID: string;
    export let user: string;
    export let lastMessage: string;
</script>


<a href={`/chat/${ID}`}>
    <div class="{chatID === ID ? 'bg-leaf-500' : 'bg-leaf-700'} hover:bg-leaf-400 flex flex-col gap-2 m-4 p-3 rounded-xl transition-all text-lightwood-200  hover:text-lightwood-900 text-lg cursor-pointer ">
        <div class="text-darkwood-900 text-xl bg-lightwood-200 pl-2 rounded-md">
            {user}
        </div>
        <div class="truncate">
            {lastMessage ? lastMessage : $_("NoRecentMessages")}
        </div>
    </div>
</a>