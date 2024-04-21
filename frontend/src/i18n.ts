import { dictionary, locale, _ } from 'svelte-i18n';
import { derived } from 'svelte/store';


const MESSAGE_FILE_URL_TEMPLATE: string = '/lang/{locale}.json';

//As i am fetching the language files, i need a variable to store
//the state of the download
//So here im checking if the locale is a string, as before loading it wont have a type(or will be undefined)
const isLocaleLoaded = derived(locale, $locale => typeof $locale === 'string');

//Setup the default locale to be es
async function setupI18n(_locale: string = 'es') {

    //Creates the route of the language needed
    const messsagesFileUrl: string = MESSAGE_FILE_URL_TEMPLATE.replace('{locale}', _locale);

    return await fetch(messsagesFileUrl)
        .then(response => response.json())
        .then((translations) => {
            dictionary.set({ [_locale]: translations });
            locale.set(_locale);
        });
}
export { _, setupI18n, isLocaleLoaded, locale };