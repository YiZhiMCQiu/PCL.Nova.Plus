<script lang="ts">
    export let slide = null
    export let after_leave = null
</script>
<div
        class="component-extension_resourcepack"
        in:slide
        out:slide
        on:outroend={after_leave}
>

</div>
<style>
    .component-extension_resourcepack {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
        background-color: lightcoral;
    }
</style>