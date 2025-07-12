<script lang="ts">
    export let slide = null
    export let after_leave = null
</script>
<div
        class="component-quilt"
        in:slide
        out:slide
        on:outroend={after_leave}
>

</div>
<style>
    .component-quilt {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>