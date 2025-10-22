<script lang="ts">
    import MyNormalSpan from "../../component/input/MyNormalSpan.svelte";
    import { download_list } from "../../store/changeBody";
    import { onDestroy } from "svelte";

    export let width = "211px";
    export let slide = null;
    export let after_leave = null;
    const unsubscribe_download_list = download_list.subscribe((value) => {
        console.log(value);
    });
    onDestroy(unsubscribe_download_list);
    let allProgress = 0;
    let downloadSpeed = 0.0;
</script>

<div
    class="component"
    style:width
    in:slide={{ x: Number(width.replace("px", "")) }}
    out:slide={{ x: Number(width.replace("px", "")) }}
    on:outroend={after_leave}
>
    <div class="container">
        <MyNormalSpan>下载进度</MyNormalSpan>
        <div class="progress">
            <MyNormalSpan style_in="font-size: 24px;"
                >{allProgress}%</MyNormalSpan
            >
        </div>
    </div>
    <div class="container">
        <MyNormalSpan>下载速度</MyNormalSpan>
        <div class="progress">
            <MyNormalSpan style_in="font-size: 24px;"
                >{downloadSpeed.toFixed(2)} MB/s</MyNormalSpan
            >
        </div>
    </div>
    <div class="container">
        <MyNormalSpan>剩余文件</MyNormalSpan>
        <div class="progress">
            <MyNormalSpan style_in="font-size: 24px;"
                >{$download_list.length}</MyNormalSpan
            >
        </div>
    </div>
</div>

<style>
    .component {
        height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
    }
    .container {
        width: 100%;
        height: 100px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    .progress {
        margin-top: 5px;
        padding-top: 5px;
        display: flex;
        align-items: center;
        flex-direction: column;
        width: calc(100% - 60px);
        border-top: 3px solid royalblue;
    }
</style>
