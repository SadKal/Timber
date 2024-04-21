<script lang="ts">
    import { Router, createRouter} from '@roxi/routify';
    import routes from '../.routify/routes.default.js';
    import { setupI18n, isLocaleLoaded } from '@/i18n';

    $: if (!$isLocaleLoaded) {
        setupI18n();
    }

    export const router = createRouter({ routes })
</script>
{#if $isLocaleLoaded}
    <Router {router} />
{/if}
