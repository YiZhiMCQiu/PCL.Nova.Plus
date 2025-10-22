<script lang="ts">
    import { current_setting } from "../../store/changeBody";
    import { onDestroy } from "svelte";
    import Launch from "./content/Launch.svelte";
    import Personalization from "./content/Personalization.svelte";
    import Other from "./content/Other.svelte";
    import { slide_up } from "../../store/functions";
    import HomePage from "./content/HomePage.svelte";
    export let slide = null;
    export let after_leave = null;
    let isTransitioning = true;
    function control_leave() {
        isTransitioning = true;
    }
    let f = false;
    const unsubscribe_current_setting = current_setting.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_current_setting);
</script>

<div class="component-right" in:slide out:slide on:outroend={after_leave}>
    {#if $current_setting === "Launch" && isTransitioning}
        <Launch slide={slide_up} after_leave={control_leave} />
    {:else if $current_setting === "Personalization" && isTransitioning}
        <Personalization slide={slide_up} after_leave={control_leave} />
    {:else if $current_setting === "Other" && isTransitioning}
        <Other slide={slide_up} after_leave={control_leave} />
    {:else if $current_setting === "HomePage" && isTransitioning}
        <HomePage slide={slide_up} after_leave={control_leave} />
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
