<script lang="ts">
    import { checkAuth } from "@/utils/auth";
    import { onMount } from "svelte"

    let loading = true;

    onMount( async () => {
        const error = await checkAuth();
        if (error) {
            document.location.href = "/login"
        }else{
            loading = false;
        }
    })

</script>

{#if !loading}
    Hola
{/if}


