<script lang="ts">
    export let width = "370px"
    export let slide = null
    export let after_leave = null
</script>
<div
        class="component"
        style:width={width}
        in:slide={{ x: Number(width.replace("px", "")) }}
        out:slide={{ x: Number(width.replace("px", "")) }}
        on:outroend={after_leave}
>

</div>
<style>
    .component {
        height: 100%;
        display: flex;
        flex-direction: column;
    }
</style>