<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    import { createEventDispatcher } from "svelte";
    export let isChecked = false;
    export let in_style = "";
    const dispatch = createEventDispatcher();
    function buttonClick() {
        dispatch("click");
    }
    // export let click = null
    $: light = $dark_mode ? "#e6e6e6" : "#1a1a1a";
</script>

<button
    class={isChecked ? "button-active" : "button-style cursor-pointer"}
    style="{in_style}; --light-color: {light};"
    on:click={buttonClick}
>
    <slot></slot>
</button>

<style>
    .button-style,
    .button-active {
        font-weight: bold;
        border-radius: 50px;
        transition: all 0.2s;
        border: 0;
    }
    .button-style {
        background-color: transparent;
        color: var(--light-color);
        stroke: var(--light-color);
    }
    .button-active {
        background-color: rgba(19, 85, 206, 0.72);
        color: #e6e6e6;
        stroke: #e6e6e6;
    }
    .button-style:hover {
        background-color: rgba(207, 207, 207, 0.72);
    }
    .button-style:active {
        background-color: rgba(207, 207, 207, 0.81);
        transform: scale(0.96);
    }
    .button-active:hover {
        background-color: rgba(0, 66, 187, 0.81);
    }
</style>
