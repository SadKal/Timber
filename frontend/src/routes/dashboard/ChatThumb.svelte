<script lang="ts">
    export let user: string;
    export let lastMessage: string;
</script>



<div class="bg-leaf-700 hover:bg-leaf-400 flex flex-col gap-2 m-4 p-3 rounded-xl transition-all text-lightwood-200  hover:text-lightwood-900 text-lg">
    <div class="text-darkwood-900 text-xl bg-lightwood-200 pl-2 rounded-md">
        {user}
    </div>
    <div class="pl-3 truncate">
        {lastMessage}
    </div>
</div>