<script>
    export let id;
    export let username;
    export let url;
    console.log(id)
</script>

<div class="flex justify-between items-center cursor-pointer m-5 p-2 bg-leaf-600 rounded-xl">
    <img src={url} class="rounded-full object-cover w-20 h-20" alt="pfpicture" id={id}/>
    <p class="mr-5 text-lightwood-100 text-2xl">{username}</p>
</div>