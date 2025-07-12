<script lang="ts">
    import MySidebarButton from "../../../component/button/MySidebarButton.svelte";
    export let slide = null
    export let after_leave = null
    import { dark_mode } from "../../../store/changeBody";
    import { current_manual } from "../../../store/changeBody";
    import {quadInOut} from "svelte/easing";
    import {onDestroy} from "svelte";
    import Forge from "./manual/Forge.svelte";
    import Fabric from "./manual/Fabric.svelte";
    import Quilt from "./manual/Quilt.svelte";
    import NeoForge from "./manual/NeoForge.svelte";
    import Optifine from "./manual/Optifine.svelte";
    import LiteLoader from "./manual/LiteLoader.svelte";
    $: dark = $dark_mode ? "rgba(32, 32, 32, 0.6)" : "rgba(255, 255, 255, 0.8)"
    let isTransitioning = true
    function control_leave() {
        isTransitioning = true
    }
    let f = false
    const unsubscribe_current_manual = current_manual.subscribe((value) => {
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
    onDestroy(unsubscribe_current_manual)
</script>
<div
        class="component-manual_install"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <div id="left" style:background-color={dark}>
        <MySidebarButton isChecked={$current_manual === 'Forge'} in_style="margin-top: 8px" click={() => current_manual.set('Forge')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    fill="none" class="side-icon">
                <path d="M7 10H6a4 4 0 0 1-4-4a1 1 0 0 1 1-1h4m0 0a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1a7 7 0 0 1-7 7H8a1 1 0 0 1-1-1zm2 7v5m6-5v5M5 20a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1"/>
            </svg>
            <span style="margin-left: 8px;">Forge</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_manual === 'Fabric'} click={() => current_manual.set('Fabric')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    fill="none" class="side-icon">
                <path d="M3 5a1.41 1.41 0 0 0 0 2l1 1a1.41 1.41 0 0 1 0 2l-1 1a1.41 1.41 0 0 0 0 2l1 1a1.41 1.41 0 0 1 0 2l-1 1a1.41 1.41 0 0 0 0 2l2 2a1.41 1.41 0 0 0 2 0l1-1a1.41 1.41 0 0 1 2 0l1 1a1.41 1.41 0 0 0 2 0l1-1a1.41 1.41 0 0 1 2 0l1 1a1.41 1.41 0 0 0 2 0l2-2a1.41 1.41 0 0 0 0-2l-1-1a1.41 1.41 0 0 1 0-2l1-1a1.41 1.41 0 0 0 0-2l-1-1a1.41 1.41 0 0 1 0-2l1-1a1.41 1.41 0 0 0 0-2l-2-2a1.41 1.41 0 0 0-2 0l-1 1a1.41 1.41 0 0 1-2 0l-1-1a1.41 1.41 0 0 0-2 0l-1 1a1.41 1.41 0 0 1-2 0L7 3a1.41 1.41 0 0 0-2 0Z"/>
            </svg>
            <span style="margin-left: 8px;">Fabric</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_manual === 'Quilt'} click={() => current_manual.set('Quilt')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    fill="none" class="side-icon">
                <g>
                    <rect
                            width="18"
                            height="18"
                            x="3" y="3"
                            rx="2"/>
                    <path d="M12 9v6m4 0v6m0-18v6M3 15h18M3 9h18M8 15v6M8 3v6"/>
                </g>
            </svg>
            <span style="margin-left: 8px;">Quilt</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_manual === 'NeoForge'} click={() => current_manual.set('NeoForge')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2" class="side-icon">
                <g>
                    <path d="M19.9 8.3C20.6 7 21 5.6 21 4c0-.6-.4-1-1-1c-2.3 0-4.3.8-5.9 2.2a15 15 0 0 0-4.2 0A8.78 8.78 0 0 0 4 3c-.6 0-1 .4-1 1c0 1.6.4 3 1.1 4.3c-.6.7-1.1 1.4-1.4 2.2C4 13 11 16 12 16s8-3 9.3-5.5c-.3-.8-.8-1.5-1.4-2.2M9 9v.5m4 3.5h-2m1 3v-3m3-4v.5"/>
                    <path d="M6.3 20.5A6.87 6.87 0 0 0 9 15H2.2c.8 4 4.9 7 9.8 7c5.5 0 10-3.8 10-8.5c0-1.1-.2-2.1-.7-3"/>
                </g>
            </svg>
            <span style="margin-left: 8px;">NeoForge</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_manual === 'Optifine'} click={() => current_manual.set('Optifine')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="1.5" class="side-icon">
                <g>
                    <path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/>
                    <path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524M3.528 7.294L8.4 10m12.1-2.722L15.6 10M12 21v-5"/>
                </g>
            </svg>
            <span style="margin-left: 8px;">Optifine</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_manual === 'LiteLoader'} click={() => current_manual.set('LiteLoader')}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="1.5" class="side-icon">
                <g>
                    <path d="M12 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>
                    <path d="M15.019 17h-6.04m6.04 0h3.614c1.876 0 1.559-1.86.61-2.804C15.825 10.801 20.68 3 11.999 3s-3.825 7.8-7.243 11.196c-.913.908-1.302 2.804.61 2.804H8.98m6.039 0c0 1.925-.648 4-3.02 4s-3.02-2.075-3.02-4"/>
                </g>
            </svg>
            <span style="margin-left: 8px;">LiteLoader</span>
        </MySidebarButton>
    </div>
    <div id="right">
        {#if $current_manual === "Forge" && isTransitioning}
            <Forge slide={slide_up} after_leave={control_leave} />
        {:else if $current_manual === "Fabric" && isTransitioning}
            <Fabric slide={slide_up} after_leave={control_leave} />
        {:else if $current_manual === "Quilt" && isTransitioning}
            <Quilt slide={slide_up} after_leave={control_leave} />
        {:else if $current_manual === "NeoForge" && isTransitioning}
            <NeoForge slide={slide_up} after_leave={control_leave} />
        {:else if $current_manual === "Optifine" && isTransitioning}
            <Optifine slide={slide_up} after_leave={control_leave} />
        {:else if $current_manual === "LiteLoader" && isTransitioning}
            <LiteLoader slide={slide_up} after_leave={control_leave} />
        {/if}
    </div>
</div>
<style>
    .component-manual_install {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
        display: flex;
    }
    #left {
        width: 144px;
        height: 100%;
        display: flex;
        flex-direction: column;
        background-color: rgb(0, 0, 0);
    }
    #right {
        position: relative;
        flex: 1;
    }
</style>