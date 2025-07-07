<script lang="ts">
    import {dark_mode} from "../../logic/changeBody";
    export let placeholder = ""
    export let style_in = ""
    export let handleInput = null
    export let title = ""
    function onInput(e: Event) {
        handleInput((e.target as HTMLInputElement).value)
    }
    $: light = $dark_mode ? '#e6e6e6cf' : '#1a1a1acf'
    $: dark = $dark_mode ? '#1a1a1acf' : '#e6e6e6cf'
</script>
<input type="text" title={title} placeholder={placeholder} class="text-input font-pcl" style="{style_in}; --light-color: {light}; --dark-color: {dark}" on:input={onInput}>
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