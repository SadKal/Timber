<script lang="ts">

    import { setupI18n, isLocaleLoaded, locale } from '@/i18n';
    import LanguageChange from '@/components/LanguageChange.svelte'
    import { checkAuth } from '@/utils/auth'

    const authRoutes: string[] = ["/", "/login", "/register"]
    let loading: Boolean = true;

    async function authRouter(): Promise<void> {
        const error = await checkAuth();

        if (!error && authRoutes.includes(document.location.pathname)) {
            document.location.href = "/dashboard"
            return
        }
        else if (error && !authRoutes.includes(document.location.pathname)) {
            document.location.href = "/login"
            return
        }
        loading = false;
    }

    authRouter()
</script>


{#if !loading}
    <div class="absolute right-20 top-10">
        <LanguageChange
            value={$locale}
            on:locale-changed={e => setupI18n(e.detail) }
        />
    </div>
    <slot />
{/if}