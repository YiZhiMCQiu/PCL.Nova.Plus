<script lang="ts">
    export let slide = null;
    export let after_leave = null;
</script>

<div
    class="component-download-file-right"
    in:slide
    out:slide
    on:outroend={after_leave}
>
    <div></div>
</div>

<style>
    .component-download-file-right {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        overflow: hidden;
    }
</style>
