<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    import { createEventDispatcher } from "svelte";
    export let isChecked = false;
    export let in_style = "";
    const dispatch = createEventDispatcher();
    function buttonClick() {
        dispatch("click");
    }
    $: light = $dark_mode ? "#e6e6e6cf" : "#1a1a1aaf";
    $: dark = $dark_mode ? "#0a0a0aaf" : "#d6d6d6cf";
</script>

<button
    class="font-pcl {isChecked
        ? 'button-active'
        : 'button-style cursor-pointer'}"
    on:click={buttonClick}
    style="{in_style}; --dark-color: {dark}; --light-color: {light};"
>
    <slot />
</button>

<style>
    .button-style,
    .button-active {
        background-color: transparent;
        position: relative;
        width: 100%;
        height: 40px;
        border: 0;
        font-size: 15px;
        transition: all 0.2s;
        text-align: left;
        color: var(--light-color);
        stroke: var(--light-color);
        padding: 0 10px;
        box-sizing: border-box;
    }
    .button-active {
        color: rgba(0, 128, 255, 207);
        stroke: rgba(0, 128, 255, 207);
    }
    .button-active::before,
    .button-style::before {
        position: absolute;
        content: "";
        width: 4px;
        height: 0;
        background-color: rgba(0, 128, 255, 207);
        left: 1px;
        top: 20px;
        transition: all 0.15s;
        border-radius: 2px;
    }
    .button-active::before {
        height: 28px;
        top: 6px;
    }
    .button-style:hover,
    .button-active:hover {
        background-color: #0099ff3f;
        border: 1px solid #0066ff3f;
        border-left: 0;
    }
</style>
