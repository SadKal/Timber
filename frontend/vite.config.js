import routify from '@roxi/routify/vite-plugin';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import preprocess from 'svelte-preprocess';
import { defineConfig } from 'vite';
import { resolve } from 'path';
import dotenv from 'dotenv';

// Load environment variables from .env file
dotenv.config();

// Determine environment mode
const production = process.env.NODE_ENV === 'production';

// Determine hydratable value
const hydratable = process.env.ROUTIFY_SSR_ENABLE === 'true';

export default defineConfig({
    clearScreen: false,
    resolve: { alias: { '@': resolve('src') } },
    plugins: [
        routify({
            render: {
                ssr: { enable: false },
            },
        }),
        svelte({
            compilerOptions: {
                dev: !production,
                hydratable: hydratable,
            },
            extensions: ['.svelte'],
            preprocess: [preprocess()],
        })
    ],

    server: { port: 1337 },
});
