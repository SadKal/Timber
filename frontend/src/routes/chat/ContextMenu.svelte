<script lang="ts">
    import Bin from "./Bin.svelte"
    import Pen from "./Pen.svelte"

    export let id;
    export let showContext;

    let posX = 0;
    let posY = 0;

    function handleClickOutside(event) {
        if (!event.target.closest(`.context-menu-${id}`)) {
            showContext = false;
        }
    }

    function handleContextMenu(event) {
        event.preventDefault();
        showContext = true;
        posX = event.clientX;
        posY = event.clientY;
    }
</script>

<div class="context-menu-{id} bg-lightwood-200 p-2 rounded-md" style="position: fixed; left: {posX-100}px; top: {posY}px;">
    <div class="flex p-2 cursor-pointer rounded-lg transition-all hover:bg-lightwood-100">
        <Pen/>
        Editar mensaje
    </div>
    <div class="flex p-2 cursor-pointer rounded-lg transition-all hover:bg-lightwood-100">
        <Bin/>
        Borrar mensaje
    </div>
</div>

<svelte:window on:click={handleClickOutside} on:contextmenu|preventDefault={handleContextMenu}/>
