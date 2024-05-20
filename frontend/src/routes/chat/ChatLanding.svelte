<script lang="ts">
    import { _ } from 'svelte-i18n';
</script>



<div class="bg-darkwood-700 w-full flex flex-col justify-center items-center text-lightwood-100 text-6xl text-center relative" style="background-image: linear-gradient(
    to bottom,
    rgba(175, 132, 71, 0.85),
    rgba(175, 132, 71, 0.85)
    ), url('/assets/backgrounds/wood_texture.webp');">
    {$_("LandingMessage")}
</div>