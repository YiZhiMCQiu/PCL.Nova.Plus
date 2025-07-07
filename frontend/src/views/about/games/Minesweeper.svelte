<script lang="ts">
    import MyNormalLabel from "../../../component/input/MyNormalLabel.svelte";

    export let slide = null
    export let after_leave = null
</script>
<div
        class="component-treasure_box"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <MyNormalLabel>NovaPlus暂不支持扫雷游戏，请去使用Nova 0.0.12版本~</MyNormalLabel>
</div>
<style>
    .component-treasure_box {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>