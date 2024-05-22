export const compareDate = (a, b) => {
    return a.created_at < b.created_at ? a : b;
}

//En javascript, las funciones son clases, por lo que se pueden incluir métodos, así consigo cancelar la llamada en caso de que este vacio el parametro de busqueda
export function debounce(callback, wait) {
    let timeout;

    function debounced(...args) {
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            timeout = null;
            callback(...args);
        }, wait);
    }

    debounced.cancel = () => {
        clearTimeout(timeout);
    };

    return debounced;
}

// const searchGames = debounce(async (textToSearch) => {
//     if (textToSearch.trim() !== '') {
//         loading = true;
//         const response = await fetch(`/api?type=search&q=${textToSearch}`, {
//             method: 'GET',
//             headers: {
//                 'content-type': 'application/json'
//             }
//         });
//         const result = await response.json();

//         searchResult = result;
//         loading = false;
//     } else {
//         searchResult = [];
//     }
// }, 500);
