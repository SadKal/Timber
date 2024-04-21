<script>

    import { setupI18n, isLocaleLoaded, locale } from '@/i18n';
    import LanguageChange from '@/components/LanguageChange.svelte'

</script>

<div class="absolute right-20 top-10">
    <LanguageChange
        value={$locale}
        on:locale-changed={e => setupI18n(e.detail) }
    />
</div>
<slot />