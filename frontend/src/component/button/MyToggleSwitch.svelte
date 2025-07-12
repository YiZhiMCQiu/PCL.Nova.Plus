<script lang="ts">
    import {dark_mode} from "../../store/changeBody";

    export let isSelect = false
    export let html_close_span = ""
    export let html_open_span = ""
    $: ({dark, back} = $dark_mode ? {dark: "#303030cf", back: "#707070cf"} : {dark: "#eeeeeecf", back: "#909090cf"})
    export let click = null
</script>
<div class="switch" style="--dark-color: {dark}; --back-color: {isSelect ? 'skyblue' : back}; --left-value: {isSelect ? '28px' : '4px'}" on:click={click} on:keydown={(e) => e.preventDefault()}>
    <button>{@html isSelect ? html_open_span : html_close_span}</button>
</div>
<style>
    .switch {
        position: relative;
        border: 0;
        width: 48px;
        height: 24px;
        background-color: var(--back-color);
        border-radius: 34px;
        cursor: pointer;
        transition: 0.2s;
    }
    .switch button {
        border-radius: 50%;
        cursor: pointer;
        width: 16px;
        height: 16px;
        margin: 4px 4px 4px var(--left-value);
        border: 0;
        background-color: var(--dark-color);
        transition: 0.2s;
        box-shadow: 2px 2px 5px gray;
    }
</style>