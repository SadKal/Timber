/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: 'class',
    content: ['./src/**/*.svelte', './src/**/*.css', './index.html'],
    theme: {
        fontFamily: {
            'fanwood': ['Fanwood Text']
        },
        extend: {},
        colors: {
            transparent: 'transparent',
            lightwood: {
                '50': '#fbf8ef',
                '100': '#f1e4c3',
                '200': '#e9d49e',
                '300': '#ddb96c',
                '400': '#d5a34a',
                '500': '#cb8735',
                '600': '#b3692c',
                '700': '#964e27',
                '800': '#7b3f25',
                '900': '#653422',
                '950': '#391b0f',
            },
            darkwood: {
                '50': '#faf8f2',
                '100': '#f4f0e0',
                '200': '#e8dfc0',
                '300': '#d9c898',
                '400': '#c6a969',
                '500': '#bd9752',
                '600': '#af8447',
                '700': '#92693c',
                '800': '#765536',
                '900': '#60472e',
                '950': '#332417',
            },
            leaf: {
                '50': '#f6f8f5',
                '100': '#e9f0e8',
                '200': '#d3e1d1',
                '300': '#b1c9ac',
                '400': '#86a97f',
                '500': '#648a5d',
                '600': '#597e52',
                '700': '#40593c',
                '800': '#364833',
                '900': '#2d3c2b',
                '950': '#151f14',
            },
            light: {
                '50': '#ffffec',
                '100': '#fdffc0',
                '200': '#ffff85',
                '300': '#fff73f',
                '400': '#ffe80b',
                '500': '#f4ce00',
                '600': '#d3a000',
                '700': '#a87200',
                '800': '#8a5909',
                '900': '#75480e',
                '950': '#452603',
            }
        }
    },
    plugins: [require('tailwindcss-debug-screens')],
};
