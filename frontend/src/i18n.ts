import { dictionary, locale, _ } from 'svelte-i18n';
import { derived } from 'svelte/store';


const MESSAGE_FILE_URL_TEMPLATE: string = '/lang/{locale}.json';


const isLocaleLoaded = derived(locale, $locale => typeof $locale === 'string');


async function setupI18n(_locale: string = 'es') {

    const messsagesFileUrl: string = MESSAGE_FILE_URL_TEMPLATE.replace('{locale}', _locale);

    return await fetch(messsagesFileUrl)
        .then(response => response.json())
        .then((translations) => {
            dictionary.set({ [_locale]: translations });
            locale.set(_locale);
        });
}
export { _, setupI18n, isLocaleLoaded, locale };

