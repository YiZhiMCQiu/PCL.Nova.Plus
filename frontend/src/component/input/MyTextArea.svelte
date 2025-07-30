<script lang="ts">
    import {dark_mode} from "../../store/changeBody";
    export let value = ""
    export let placeholder = ""
    export let handleInput: ((e: string) => void | null) = null
    export let style_in = ""
    export let title = ""
    export let isReadonly = false
    function textAreaChange(e: Event) {
        if (handleInput == null) {
            return
        }else{
            handleInput((e.target as HTMLTextAreaElement).value)
        }
    }
    $: ({light, dark} = $dark_mode ? {light: '#e6e6e6cf', dark: '#1a1a1acf'} : {light: '#1a1a1acf', dark: '#e6e6e6cf'})
</script>
<textarea readonly={isReadonly} class="text-input font-pcl" title={title} style="{style_in}; --light-color: {light}; --dark-color: {dark};" bind:value={value} placeholder={placeholder} on:change={textAreaChange}/>
<style>
    .text-input {
        border: 1px solid var(--light-color);
        border-radius: 8px;
        background-color: var(--dark-color);
        color: var(--light-color);
        padding: 10px;
    }
</style>