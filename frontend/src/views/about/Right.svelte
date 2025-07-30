<script lang="ts">
    import { current_about } from "../../store/changeBody";
    import {onDestroy} from "svelte";
    import Help from "./content/Help.svelte";
    import About from "./content/About.svelte";
    import TreasureBox from "./content/TreasureBox.svelte";
    import Feedback from "./content/Feedback.svelte";
    import Discussion from "./content/Discussion.svelte";
    import Minesweeper from "./games/Minesweeper.svelte";
    import P2048 from "./games/P2048.svelte";
    import {slide_up} from "../../store/functions";
    export let slide = null
    export let after_leave = null
    let isTransitioning = true
    function control_leave() {
        isTransitioning = true
    }
    let f = false
    const unsubscribe_current_about = current_about.subscribe((value) => {
        if(!f) {
            f = true
            isTransitioning = true
        }else{
            isTransitioning = !isTransitioning
        }
    })
    onDestroy(unsubscribe_current_about)
</script>
<div
        class="component-right"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    {#if $current_about === "Help" && isTransitioning}
        <Help slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "About" && isTransitioning}
        <About slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "Treasure-Box" && isTransitioning}
        <TreasureBox slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "Feedback" && isTransitioning}
        <Feedback slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "Discussion" && isTransitioning}
        <Discussion slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "Minesweeper" && isTransitioning}
        <Minesweeper slide={slide_up} after_leave={control_leave} />
    {:else if $current_about === "P2048" && isTransitioning}
        <P2048 slide={slide_up} after_leave={control_leave} />
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