<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {messagebox, MSG_ERROR, MSG_INFO, MSG_WARNING} from "../../../store/messagebox";
    import {dont_click} from "../../../store/changeBody";
    import {GetMachineCode} from "../../../../wailsjs/go/main/App";
    import {GetTodayLucky} from "../../../../wailsjs/go/main/App.js";

    export let slide = null
    export let after_leave = null

    async function dontClick() {
        await messagebox("警告", "PCL Nova Plus 作者不会受理由于点击千万别点造成的任何 Bug。这是最后的警告，是否继续操作？", MSG_ERROR, ['ok', 'ok', 'ok'])
        let rand = Math.floor(Math.random() * 3 + 1)
        // let rand = 2
        dont_click.set(rand)
    }
    async function todayLucky() {
        let luck = await GetTodayLucky(await GetMachineCode())
        if(luck == -1) {
            await messagebox("警告", "PCL Nova Plus 无法正常获取到 当前硬件配置UUID 你可能在虚拟机中运行 Nova Plus!", MSG_WARNING)
        }else if(luck == -2) {
            await messagebox("警告", "PCL Nova Plus 无法正常获取到 PCL ID 请尝试在官版 PCL 获取一次识别码，随后再使用 Nova 查看今日人品~", MSG_WARNING)
        }else if(luck == -3) {
            await messagebox("暂不支持", "PCL Nova Plus 暂时无法为除了 Windows 操作系统的系统提供 今日人品，请见谅~")
        }else{
            let str;
            if(luck == 100) {
                str = "！100！100！！！！！"
            } else if (luck == 99) {
                str = "！额，但不是 100……"
            } else if (luck >= 90) {
                str = "！好评如潮！"
            } else if (luck >= 60) {
                str = "！是不错的一天呢！"
            } else if (luck > 50) {
                str = "！还行啦还行啦。"
            } else if (luck == 50) {
                str = "！五五开……"
            } else if (luck >= 40) {
                str = "！还……还行吧……？"
            } else if (luck >= 11) {
                str = "！呜哇……"
            } else if (luck >= 1) {
                str = "……（没错，是百分制）"
            } else {
                str = "……"
                if(await messagebox("今日人品 - 附加使用条款", "在查看结果前，请先同意以下附加使用条款：<br><br>1. 我知晓并了解 PCL2 的今日人品功能完全没有出 Bug。<br>2. PCL2 不对使用本软件所间接造成的一切财产损失（如砸电脑等）等负责。", MSG_ERROR, ["再见", "同意并查看结果"]) == 0) {
                    return
                }
            }
            const date = new Date()
            await messagebox("今日人品 - " + date.getFullYear() + "/" + (date.getMonth() + 1) + "/" + date.getDate(), "你今天的人品值是：" + luck + str, MSG_INFO, ["我知道了"])
        }
    }
</script>
<div
        class="component-treasure_box"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <MySelectCard title="百宝箱">
        <div class="proc">
            <div style="margin: 0 20px; width: calc(100% - 40px)">
                {#if $dont_click === 0}
                    <MyNormalButton style_in="border: 1px solid red; color: red; width: 130px; height: 35px" on:click={dontClick}>千万别点</MyNormalButton>
                {/if}
                <MyNormalButton style_in="width: 130px; height: 35px; margin-left: 10px" on:click={todayLucky}>今日人品</MyNormalButton>
            </div>
        </div>
    </MySelectCard>
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