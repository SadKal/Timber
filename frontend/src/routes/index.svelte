<script lang="ts">
    import Button from "@/components/Button.svelte"
</script>

<div class="w-screen h-screen bg-leaf-600 bg-opacity-95 flex flex-col items-center" style="background-image: linear-gradient(
        to bottom,
        rgba(89, 126, 82, 0.95),
        rgba(89, 126, 82, 0.95)
        ), url('/assets/mountains.jpg');">
    <div class="mt-28 mb-4 w-52 h-52">
        <img width="200" height="200" src="assets/logos/logo_dark_bg.png" alt="Logo dark" />
    </div>
    <h2 class="text-lightwood-100 font-fanwood text-4xl mt-2 mb-10 text-center">Rooted in nature, connected in spirit</h2>

    <div class="w-96 h-52 bg-lightwood-100 rounded-xl flex flex-col flex-shrink-5 items-center justify-evenly shadow-md shadow-leaf-950">
        <Button link={true}  href="/login" content="Login"/>
        <Button link={true}  href="/register" content="Register"/>
    </div>
</div>

