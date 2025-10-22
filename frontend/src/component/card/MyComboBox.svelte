<!--
这只是一个未完成品喵~
-->
<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { dark_mode } from "../../store/changeBody";
    import { slide_opacity } from "../../store/functions";
    export let style_in = "";
    export let values = [];
    export let canInput = false;
    export let text = "";
    export let placeholder = "";
    const dispatch = createEventDispatcher();
    let isDropDown = false;
    let DropDownDOM: HTMLDivElement | null = null;
    let DropMenuDOM: HTMLDivElement | null = null;
    function allClick() {
        isDropDown = !isDropDown;
    }
    function inputClick() {
        if (!canInput) {
            allClick();
        }
    }
    function childClick(index: number) {}
    $: ({ light, dark } = $dark_mode
        ? { light: "#e6e6e6cf", dark: "#1a1a1acf" }
        : { light: "#1a1a1acf", dark: "#e6e6e6cf" });
</script>

<div
    style="width: 200px; height: 30px; {style_in}; --dark-color: {dark}; --light-color: {light}; position: relative; min-width: 200px; min-height: 30px;"
    class="dropdown"
    bind:this={DropDownDOM}
>
    <div class="dropdown-toggle">
        <input
            type="text"
            {placeholder}
            class="dropdown-input font-pcl"
            style:cursor={canInput ? "text" : "pointer"}
            readonly={!canInput}
            on:click={inputClick}
            value={text}
        />
        <button class="dropdown-button" on:click={allClick}>
            <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="3"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none"
                class="change-svg"
                style:stroke={light}
                style:transform={isDropDown ? "rotate(180deg)" : "rotate(0)"}
            >
                <polyline points="6 10 12 16 18 10" />
            </svg>
        </button>
    </div>
    {#if isDropDown}
        <div
            bind:this={DropMenuDOM}
            class={isDropDown ? "dropdown-menu-show" : "dropdown-menu"}
            in:slide_opacity
            out:slide_opacity
        ></div>
    {/if}
</div>

<style>
    .dropdown {
        border: 1px solid var(--light-color);
        border-radius: 5px;
    }
    .dropdown-toggle {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        border-radius: 5px;
    }
    .dropdown-menu,
    .dropdown-menu-show {
        position: absolute;
        top: 100%;
        left: 0;
        width: 100%;
        height: 100px;
        z-index: 99;
        overflow-y: auto;
        background-color: white;
    }
    .dropdown-input {
        border-bottom-left-radius: 5px;
        border-top-left-radius: 5px;
        /*flex: 1;*/
        width: calc(100% - 40px);
        height: 100%;
        outline: none;
        border: none;
        padding: 0 5px;
    }
    .dropdown-button {
        width: 30px;
        height: 100%;
        position: relative;
        background-color: white;
        border: none;
        border-top-right-radius: 5px;
        border-bottom-right-radius: 5px;
        cursor: pointer;
        flex-shrink: 0;
    }
    .change-svg {
        width: 12px;
        height: 12px;
        position: absolute;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        margin: auto;
        transition: all 0.2s;
    }
</style>
