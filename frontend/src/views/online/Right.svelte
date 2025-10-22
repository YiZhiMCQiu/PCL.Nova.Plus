<script lang="ts">
    import { current_online } from "../../store/changeBody";
    import { onDestroy } from "svelte";
    import { slide_up } from "../../store/functions";
    import IPv6 from "./content/IPv6.svelte";

    export let slide = null;
    export let after_leave = null;
    let isTransitioning = true;
    function control_leave() {
        isTransitioning = true;
    }
    let f = false;
    const unsubscribe_current_online = current_online.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_current_online);
</script>

<div class="component-right" in:slide out:slide on:outroend={after_leave}>
    {#if $current_online === "IPv6" && isTransitioning}
        <IPv6 slide={slide_up} after_leave={control_leave} />
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
