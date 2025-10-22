<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    import { createEventDispatcher } from "svelte";
    $: light = $dark_mode ? "#e6e6e6cf" : "#1a1a1acf";
    export let isChecked = false;
    export let title = "";
    export let style_in = "";
    let dispatch = createEventDispatcher();
    function onCheckClick() {
        dispatch("click");
    }
</script>

<div
    class="check cursor-pointer"
    style="{style_in}; --light-color: {light}"
    {title}
    on:click={onCheckClick}
    on:keydown|preventDefault
>
    <div class="correct">
        <svg
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
            stroke-width="3"
            stroke-linecap="round"
            stroke-linejoin="round"
            fill="none"
            style="width: 13px; height: 13px;"
            style:transform={isChecked ? "scale(1)" : "scale(0)"}
        >
            <path d="M1.5 16.5 L7.5 22.5 M7.5 22.5 L22.5 7.5" />
        </svg>
    </div>
    <div class="slot">
        <slot />
    </div>
</div>

<style>
    .check {
        color: var(--light-color);
        height: 30px;
        width: max-content;
        display: flex;
        align-items: center;
    }
    .slot {
        position: relative;
        top: -1px;
    }
    .check .correct {
        margin-right: 5px;
        transition: all 0.2s;
        border-radius: 5px;
        width: 17px;
        height: 17px;
        border: 2px solid deepskyblue;
        position: relative;
    }
    .check:hover .correct {
        border: 2px solid #00abebcf;
    }
    .check:active .correct {
        transform: scale(90%);
    }
    .correct svg {
        position: absolute;
        top: 2px;
        left: 2px;
        width: 13px;
        height: 13px;
        stroke: deepskyblue;
        transition: transform 0.2s;
    }
</style>
