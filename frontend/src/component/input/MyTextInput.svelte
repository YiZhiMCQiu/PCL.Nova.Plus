<script lang="ts">
    import {dark_mode} from "../../store/changeBody";
    export let placeholder = ""
    export let style_in = ""
    export let handleInput = (e: string) => {}
    export let title = ""
    export let value = ""
    function onInput(e: Event) {
        handleInput((e.target as HTMLInputElement).value)
    }
    $: ({light, dark} = $dark_mode ? {light: '#e6e6e6cf', dark: '#1a1a1acf'} : {light: '#1a1a1acf', dark: '#e6e6e6cf'})
</script>
<input type="text" bind:value={value} title={title} placeholder={placeholder} class="text-input font-pcl" style="{style_in}; --light-color: {light}; --dark-color: {dark};" on:input={onInput}>
<style>
    .text-input {
        border: 1px solid var(--light-color);
        border-radius: 5px;
        background-color: var(--dark-color);
        color: var(--light-color);
        padding-left: 10px;
        padding-right: 10px;
    }
</style>