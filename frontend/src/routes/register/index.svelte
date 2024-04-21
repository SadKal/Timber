<script lang="ts">
    import Button from "@/components/Button.svelte"
    import { _ } from 'svelte-i18n';

    let username: string = "";
    let password: string = "";
    let imageUrl: string | ArrayBuffer;
    let file: File = null;


    function getImgData(event): void {
        const file = event.target.files[0];
        const reader = new FileReader();

        reader.onload = () => {
            imageUrl = reader.result;
        };

        reader.readAsDataURL(file);
    }

    async function register() {
        try {
            const uploadedFile = new FormData();
            uploadedFile.append('file', file);
            uploadedFile.append('username', username);
            uploadedFile.append('password', password);


            const response = await fetch('http://localhost:8080/register', {
                method: 'POST',
                body: uploadedFile
            });

            const data = await response.json();
        } catch (error) {
            console.error('Error:', error);
        }
    }
</script>


<div class="w-screen h-screen bg-leaf-600 bg-opacity-95 flex flex-col gap-10 items-center font-fanwood" style="background-image: linear-gradient(
    to bottom,
    rgba(89, 126, 82, 0.95),
    rgba(89, 126, 82, 0.95)
    ), url('/assets/mountains.jpg');">
    <div class="mt-12 mb-0">
        {#if imageUrl}
            <img width=200 height=200 src={imageUrl} alt="Uploaded Profile" />
        {:else}
            <img width=200 height=200 src="assets/logos/logo_dark_bg.png" alt="Logo dark" />
        {/if}
    </div>
    <h1 class="text-lightwood-100 text-5xl text-center">{$_("Register")}</h1>
    <form on:submit|preventDefault = {register} class="w-96 gap-10 p-10 bg-lightwood-100 rounded-xl flex flex-col items-center shadow-[4.0px_8.0px_8.0px_rgba(0,0,0,0.38)]">
        <div class="flex flex-col items-center text-darkwood-950 text-2xl">
            <label for="username">{$_("Username")}:</label>
            <input type="text" id="username" bind:value={username} />
        </div>
        <div class="flex flex-col items-center text-darkwood-950 text-2xl">
            <label for="password">{$_("Password")}:</label>
            <input type="password" id="password" bind:value={password} />
        </div>
        <div class="flex flex-col items-center text-darkwood-950 text-2xl">
            <label for="avatar">{$_("ChoosePfp")}:</label>
            <input class="text-sm" type="file" id="avatar" on:change={(event) => {
                getImgData(event)
                file = event.target.files[0]
                }}
                accept="image/*">
        </div>
        <Button content="Register"/>
    </form>
</div>

