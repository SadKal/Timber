<script lang="ts">
    import { setupI18n, locale } from '@/i18n';
    import LanguageChange from '@/components/LanguageChange.svelte'
    import { checkAuth } from '@/utils/auth'

    let currentLocation: string = document.location.pathname;
    const authRoutes: string[] = ["/", "/login", "/register"]
    let loading: Boolean = true;

    async function authRouter(): Promise<void> {
        const error = await checkAuth();
        if (!error && authRoutes.includes(currentLocation)) {
            document.location.href = "/chat"
            return
        }
        else if (error && !authRoutes.includes(currentLocation)) {
            document.location.href = "/"
            return
        }
        loading = false;
    }

    authRouter();
</script>


{#if !loading}
    <div class="absolute right-10 top-5 z-50">
        <LanguageChange
            value={$locale}
            on:locale-changed={e => setupI18n(e.detail) }
        />
    </div>
    <slot />
{/if}