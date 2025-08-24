<script lang="ts">
    import {dark_mode} from "../../store/changeBody.js";
    import {createEventDispatcher, onMount} from "svelte";
    $: ({light, dark, hov} = $dark_mode ? {light: '#f8f8f8cf', dark: '#151515cf', hov: '#151525cf'} : {light: '#151515cf', dark: '#f8f8f8cf', hov: '#e8e8f8cf'})
    export let in_style = ""
    export let title = ""
    // 是否处在 Expand 状态
    export let isExpand = false
    // 能否被 Expand，当此值为 false 时，isExpand 将无效（一直处于打开状态）
    export let canExpand = false
    // 手动指定 maxHeight
    export let maxHeight = 0
    const dispatch = createEventDispatcher()
    function cardComp() {
        if(!canExpand) return
        isExpandComp = !isExpandComp
        cardHeight = maxHeight == 0 ? (isExpandComp ? (myCardRef!.offsetHeight + 43) + 'px' : '43px') : (isExpandComp ? (maxHeight + 43) + "px" : "43px")
        dispatch('comp', {
            height: maxHeight != 0 ? maxHeight : myCardRef!.offsetHeight,
            isExpand: isExpandComp,
            title: title
        })
    }
    // export let onComp: ((height: number, isExpand: boolean, title: string) => void | null) = null
    let cardHeight = "43px"
    let isExpandComp = canExpand ? !isExpand : true
    let myCardRef: (HTMLElement | null) = null
    onMount(() => {
        if((canExpand && !isExpand) || !canExpand) {
            cardHeight = maxHeight == 0 ? ((myCardRef!.offsetHeight + 43) + "px") : (maxHeight + 43) + "px"
        }
    })
</script>
<div class="card-container" style:height={title !== '' ? cardHeight : (parseInt(cardHeight.substring(0, cardHeight.length - 2)) - 43) + 'px'} style="{in_style}; --dark-color: {dark}; --light-color: {light}; --hov-color: {hov}">
    {#if title !== ''}
        <div class="grid {canExpand ? 'cursor-pointer' : ''}"
            data-isOpen={isExpandComp ? 'expand' : 'close'}
            on:click={cardComp} on:keydown|preventDefault
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
                    style="{canExpand ? 'display: flow' : 'display: none'}">
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