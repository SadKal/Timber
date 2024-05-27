<script>
    import { _ } from 'svelte-i18n';
    import { createEventDispatcher } from 'svelte';

    export let value = 'en';
    
    const dispatch = createEventDispatcher();
    function handleLocaleChange(event) {
        event.preventDefault();
        dispatch('locale-changed', event.target.value);
    }
</script>

<div >
    <select class="bg-lightwood-100 p-1 bg-opacity-95 rounded-md" value={value} on:change={handleLocaleChange} >
        <option value="es">🇪🇸 {$_('Spanish')}</option>
        <option value="en">🇬🇧 {$_('English')}</option>
    </select>
</div>