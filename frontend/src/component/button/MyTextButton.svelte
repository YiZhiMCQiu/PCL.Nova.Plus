<script lang="ts">
    import {dark_mode} from "../../store/changeBody";
    import {createEventDispatcher, onMount} from "svelte"
    export let style_in = ""
    export let hover_style_in = ""
    export let active_style_in = ""
    export let title = ""
    const dispatch = createEventDispatcher()
    function buttonClick() {
        dispatch('click')
    }
    function getColors(dark: boolean) {
        return {
            light: dark ? '#e6e6e6cf' : '#0a0a0aaf',
            hov: dark ? '#f6f6f6cf' : '#1a1a1acf',
            act: dark ? '#f0f0f0cf' : '#101010cf'
        }
    }
    $: ({light, hov, act} = getColors($dark_mode))
    let tc: (HTMLSpanElement | null) = null
    let rawStyle = `font-size: 16px; color: ${light}; stroke: ${light}; transition: all 0.2s;`
    let hovStyle = `font-size: 16px; color: ${hov}; stroke: ${hov}; transition: all 0.2s; cursor: pointer;`
    let actStyle = `font-size: 16px; color: ${act}; stroke: ${act}; transition: all 0.2s; cursor: pointer;`
    function onMouseEnter() {
        tc.style.cssText = hovStyle + hover_style_in;
    }
    function onMouseLeave() {
        tc.style.cssText = rawStyle + style_in;
    }
    function onMouseDown() {
        tc.style.cssText = actStyle + active_style_in
    }
    onMount(() => {
        tc.style.cssText = rawStyle + style_in
    })
</script>
<span
        class="font-pcl"
        bind:this={tc}
        title={title}
        on:mouseenter={onMouseEnter}
        on:mouseleave={onMouseLeave}
        on:mousedown={onMouseDown}
        on:mouseup={onMouseEnter}
        on:click={buttonClick}
        on:keydown|preventDefault>
    <slot />
</span>