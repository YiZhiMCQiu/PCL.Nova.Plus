<script lang="ts">
    import {dark_mode} from "../../store/changeBody.js";
    import {onMount} from "svelte";
    $: ({light, dark, hov} = $dark_mode ? {light: '#f8f8f8cf', dark: '#151515cf', hov: '#151525cf'} : {light: '#151515cf', dark: '#f8f8f8cf', hov: '#e8e8f8cf'})
    export let in_style = ""
    export let title = ""
    export let isExpand = false
    export let maxHeight = 0
    export let onComp: ((height: number, isExpand: boolean, title: string) => void | null) = null
    let cardHeight = "43px"
    let isExpandComp = !isExpand
    let myCardRef: (HTMLElement | null) = null
    function changeProps() {
        if(!isExpand) return
        isExpandComp = !isExpandComp
        cardHeight = maxHeight == 0 ? (isExpandComp ? (myCardRef!.offsetHeight + 43) + 'px' : '43px') : (isExpandComp ? (maxHeight + 43) + "px" : "43px")
        if(onComp != null) {
            onComp(maxHeight != 0 ? maxHeight : myCardRef!.offsetHeight, isExpandComp, title)
        }
    }
    onMount(() => {
        if(!isExpand) {
            cardHeight = maxHeight == 0 ? ((myCardRef!.offsetHeight + 43) + "px") : (maxHeight + 43) + "px"
        }
    })
</script>
<div class="card-container" style="{in_style}; --card-height: {title !== '' ? cardHeight : (parseInt(cardHeight.substring(0, cardHeight.length - 2)) - 43) + 'px'}; --dark-color: {dark}; --light-color: {light}; --hov-color: {hov}">
    {#if title !== ''}
        <div class="grid {isExpand ? 'cursor-pointer' : ''}"
            data-isOpen={isExpandComp ? 'expand' : 'close'}
            on:click={() => changeProps()} on:keydown|preventDefault
        >
            {title}
            <svg
                    role="img"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-width="3"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    fill="none" class="{isExpandComp ? 'card-icon-expand' : 'card-icon'}"
                    style="{isExpand ? 'display: flow' : 'display: none'}">
                <polyline points="6 10 12 16 18 10"/>
            </svg>
        </div>
    {:else}
        <div class="grid-close" style="height: fit-content;"></div>
    {/if}
    <div bind:this={myCardRef}>
        <slot />
    </div>
</div>
<style>
    .card-container {
        flex-shrink: 0;
        height: var(--card-height);
        border-radius: 6px;
        transition: all 0.2s;
        margin: 15px 22px 0 22px;
        background-color: var(--dark-color);
        overflow: hidden;
    }
    .card-container:last-child {
        margin-bottom: 15px;
    }
    .card-container:hover {
        box-shadow: 0 0 4px #0077ff;
        background-color: var(--hov-color);
    }
    .grid, .grid-close {
        padding: 10px 20px;
        font-size: 16px;
        font-weight: bold;
        transition: all 0.2s;
        color: var(--light-color);
    }
    .grid-close {
        padding: 0;
    }
    .grid[data-isOpen="expand"] + div {
        height: fit-content;
    }
    .grid + div,
    .grid-close + div{
        width: 100%;
        transition: all 0.2s;
        color: var(--light-color);
        overflow: hidden;
    }
    .card-icon,
    .card-icon-expand {
        width: 20px;
        height: 20px;
        vertical-align: middle;
        transition: all 0.2s;
        float: right;
        stroke: var(--light-color);
    }
    .card-icon-expand {
        transform: rotate(180deg);
    }
</style>