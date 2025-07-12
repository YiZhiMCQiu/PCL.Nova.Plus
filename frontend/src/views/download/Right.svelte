<script lang="ts">
    import {current_download} from "../../store/changeBody";
    import {quadInOut} from "svelte/easing";
    import {onDestroy} from "svelte";
    import AutoInstall from "./content/AutoInstall.svelte";
    import ManualInstall from "./content/ManualInstall.svelte";
    import ExtensionMod from "./content/ExtensionMod.svelte";
    import ExtensionModpack from "./content/ExtensionModpack.svelte";
    import ExtensionDatapack from "./content/ExtensionDatapack.svelte";
    import ExtensionResourcepack from "./content/ExtensionResourcepack.svelte";
    import ExtensionShaderpack from "./content/ExtensionShaderpack.svelte";
    export let slide = null
    export let after_leave = null
    let isTransitioning = true
    function control_leave() {
        isTransitioning = true
    }
    let f = false
    const unsubscribe_current_download = current_download.subscribe((value) => {
        if(!f) {
            f = true
            isTransitioning = true
        }else{
            isTransitioning = !isTransitioning
        }
    })
    function slide_up(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number, n: number) => {
                return `
                    transform: translateY(${-50 * n}%);
                    opacity: ${t};
                `
            }
        }
    }
    onDestroy(unsubscribe_current_download)
</script>
<div
        class="component-right"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    {#if $current_download === "Auto-Install" && isTransitioning}
        <AutoInstall slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Manual-Install" && isTransitioning}
        <ManualInstall slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Extension-Mod" && isTransitioning}
        <ExtensionMod slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Extension-Modpack" && isTransitioning}
        <ExtensionModpack slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Extension-Datapack" && isTransitioning}
        <ExtensionDatapack slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Extension-Resourcepack" && isTransitioning}
        <ExtensionResourcepack slide={slide_up} after_leave={control_leave} />
    {:else if $current_download === "Extension-Shaderpack" && isTransitioning}
        <ExtensionShaderpack slide={slide_up} after_leave={control_leave} />
    {/if}
</div>
<style>
    .component-right {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>