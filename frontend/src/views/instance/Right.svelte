<script lang="ts">
    import { current_instance } from "../../store/changeBody";
    export let slide = null;
    export let after_leave = null;
    import { slide_up } from "../../store/functions";
    import { onDestroy } from "svelte";
    import Overview from "./content/Overview.svelte";
    import Setting from "./content/Setting.svelte";
    let isTransitioning = true;
    function control_leave() {
        isTransitioning = true;
    }
    let f = false;
    const unsubscribe_current_instance = current_instance.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_current_instance);
</script>

<div class="component-instance" in:slide out:slide on:outroend={after_leave}>
    {#if $current_instance === "overview" && isTransitioning}
        <Overview slide={slide_up} after_leave={control_leave} />
    {:else if $current_instance === "setting" && isTransitioning}
        <Setting slide={slide_up} after_leave={control_leave} />
    {/if}
</div>

<style>
    .component-instance {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>
