<script lang="ts">
    import MySelectCard from "../../component/card/MySelectCard.svelte";
    import MyNormalButton from "../../component/button/MyNormalButton.svelte";
    import {
        HNT_ERROR,
        HNT_INFO, HNT_PASS,
        HNT_WARNING,
        inputbox,
        messagebox,
        MSG_ERROR,
        MSG_INFO,
        MSG_WARNING,
        showHint
    } from "../../store/messagebox";
    import MyLoadingPickaxe from "../../component/card/MyLoadingPickaxe.svelte";
    import MyProgressBar from "../../component/input/MyProgressBar.svelte";
    import MyNormalLabel from "../../component/input/MyNormalLabel.svelte";
    import MyTextArea from "../../component/input/MyTextArea.svelte";
    import {game_log} from "../../store/mc";
    import {EventsOn} from "../../../wailsjs/runtime";

    export let slide = null
    export let after_leave = null

    let loading_state = false
    let loading_text = "加载中~"

    EventsOn('game_log', (log: string) => {
        // game_log.set(log)
    })

    async function testMessageBox() {
        let test1 = await messagebox("信息测试", "这是一个信息测试");
        console.log(test1)
        let test2 = await messagebox("警告测试", "这是一个<br>警告测试<br>换行测试", MSG_WARNING);
        console.log(test2)
        let test3 = await messagebox("错误测试", "这是一个错误测试", MSG_ERROR);
        console.log(test3)
        let test4 = await messagebox("多按钮测试", "这是一个多按钮测试", MSG_INFO, ["ok", "ok", "ok", "ok", "ok", "ok"]);
        console.log(test4)
        let test5 = await messagebox("输入框测试", "这是一个文字非常长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长的输入框")
        console.log(test5)
        await messagebox("输入的文字2", `你的选项是：${test5}`)
    }
    function testHint() {
        showHint("这是一个提示框😀", HNT_INFO)
        setTimeout(() => {
            showHint("这是一个警告框😰", HNT_WARNING)
        }, 3200)
        setTimeout(() => {
            showHint("这是一个错误框😭", HNT_ERROR)
        }, 6400)
        setTimeout(() => {
            showHint("这是一个通过框😋", HNT_PASS)
        }, 9600)
    }
    async function testInput() {
        let test1 = await inputbox("输入框测试", "这是一个非常简单的输入框", 0, "请任意输入文字")
        console.log(test1)
        await messagebox("输入的文字1", `你输入的文字是：${test1}`)
    }
    let progress = 20
    function dragProgressBar(value: number) {
        progress = value
    }
</script>
<div
        class="component-right"
        in:slide
        out:slide
        on:outroend={after_leave}>
    <MySelectCard title="游戏日志">
        <div class="version-all">
            <MyTextArea title="启动游戏后会在这里显示日志~" style_in="margin-top: 2px; width: calc(100% - 25px); height: 500px" isReadonly={true} value="目前暂未实现，请等待下一个版本的更新~"/>
        </div>
    </MySelectCard>
    <MySelectCard title="可以折叠的卡片~" isExpand={true}>
        <div class="proc">
            <MyLoadingPickaxe state={loading_state} loading_text={loading_text} />
            <div style="margin: 10px 0">
                <MyNormalButton style_in="width: 80px; height: 30px" click={() => {loading_state = false; loading_text = "加载中~"}}>加载中</MyNormalButton>
                <MyNormalButton style_in="width: 80px; height: 30px" click={() => {loading_state = true; loading_text = "加载失败~"}}>失败</MyNormalButton>
            </div>
            <MyNormalButton style_in="width: 170px; height: 30px" click={testMessageBox}>测试信息框</MyNormalButton>
            <MyNormalButton style_in="width: 170px; height: 30px; margin-top: 10px" click={testInput}>测试输入框</MyNormalButton>
            <MyNormalButton style_in="width: 170px; height: 30px; margin-top: 10px" click={testHint}>测试提示</MyNormalButton>
            <MyProgressBar min={0} max={100} value={20} onDragging={dragProgressBar}/>
            <MyNormalLabel>当前滑动条值是：{progress}</MyNormalLabel>
        </div>
    </MySelectCard>
    <MySelectCard>
        <div class="proc">
            <p>这是无标题的卡片</p>
            <p>无标题的卡片默认会无视isExpand，因为它不会折叠</p>
        </div>
    </MySelectCard>
    <MySelectCard title="带有标题的卡片">
        <div class="proc">
            <p>这是带有标题的卡片</p>
            <p>卡片默认的isExpand是false，因此如果你想要显式的可折叠卡片，你需要手动将isExpand设为true</p>
        </div>
    </MySelectCard>
    <MySelectCard title="加载一个苹果~" isExpand={true}>
        <div class="proc">
            <p>加载一张苹果图片~</p>
            <img src="https://ts1.tc.mm.bing.net/th/id/R-C.54916b18a985e6a9c2b4cf1be60eef25?rik=8Ti1QEH7JkbCLA&riu=http%3a%2f%2fpic.616pic.com%2fys_bnew_img%2f00%2f03%2f69%2fg70yiNCFvx.jpg&ehk=uOYCWB%2fXSgQtsQC%2fRiCa8pW1wvaSMke8Md4zcEJTKUI%3d&risl=&pid=ImgRaw&r=0" alt="头像" width="300" height="300">
        </div>
    </MySelectCard>

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