<script lang="ts">
    import {dark_mode} from "../../logic/changeBody";
    export let isDisabled = false;
    export let style_in = ""
    export let click = null;
    $: light = $dark_mode ? "#f6f6f6cf" : "#0a0a0aaf"
    $: hov = $dark_mode ? "#213646cf" : "#c3e0fdcf"
</script>
<button on:click={click} style="{style_in}; --hov-color: {hov}; --light-color: {light};" class="button-style cursor-pointer" disabled={isDisabled}>
    <slot />
</button>
<style>
    .button-style {
        background-color: transparent;
        border-radius: 6px;
        border: 1px solid var(--light-color);
        transition: all 0.2s;
        color: var(--light-color);
    }
    .button-style[disabled] {
        color: #7f7f7f7f;
        border: 1px solid #7f7f7f7f;
        cursor: auto;
    }
    .button-style[disabled]:hover {
        background-color: transparent;
    }
    .button-style[disabled]:active {
        transform: scale(1);
    }
    .button-style:hover {
        background-color: var(--hov-color);
    }
    .button-style:active {
        transform: scale(0.96);
    }
</style>