<script lang="ts">
    import {dark_mode} from "../../store/changeBody";
    import {createEventDispatcher} from "svelte";

    export let isChecked = false
    export let style_in = ""
    const dispatch = createEventDispatcher()
    function buttonClick() {
        dispatch('click')
    }
    export let title = ""
    $: light = $dark_mode ? '#e6e6e6cf' : '#1a1a1acf'
</script>
<div
        class="radio {isChecked ? 'button-active' : 'button-style cursor-pointer'}"
        style="{style_in}; --light-color: {light}"
        title={title}
        on:click={buttonClick}
        on:keydown|preventDefault>
    <div class="circle"></div>
    <div class="slot">
        <slot />
    </div>
</div>
<style>
    .radio {
        color: var(--light-color);
        height: 30px;
        width: max-content;
        display: flex;
        align-items: center;
        flex-shrink: 0;
    }
    .slot {
        position: relative;
        top: -1px;
    }
    .circle {
        position: relative;
        margin-right: 5px;
        transition: all 0.2s;
        border-radius: 50%;
        width: 17px;
        height: 17px;
    }
    .circle::before {
        content: '';
        position: absolute;
        top: 2px;
        left: 2px;
        width: 13px;
        height: 13px;
        transform: scale(0);
        border-radius: 50%;
        background-color: deepskyblue;
        transition: all 0.2s;
    }
    .button-style .circle {
        border: 1px solid #a0a0a0;
    }
    .button-active .circle {
        border: 1px solid deepskyblue;
    }
    .button-style:hover .circle {
        border: 1px solid deepskyblue;
    }
    .button-style:active .circle {
        transform: scale(86%);
    }
    .button-active .circle::before {
        transform: scale(1);
    }
</style>