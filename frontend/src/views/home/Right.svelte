<script lang="ts">
    import MySelectCard from "../../component/card/MySelectCard.svelte";
    import MyNormalButton from "../../component/button/MyNormalButton.svelte";
    import {messagebox} from "../../store/messagebox";
    import MyLoadingPickaxe from "../../component/card/MyLoadingPickaxe.svelte";

    export let slide = null
    export let after_leave = null

    let loading_state = false
    let loading_text = "加载中~"

    async function testMessageBox() {
        let test1 = await messagebox("信息测试", "这是一个信息测试", 0, ["ok", "no"]);
        console.log(test1)
        let test2 = await messagebox("警告测试", "这是一个<br>警告测试<br>换行测试", 1, ["ok", "no", "yes"]);
        console.log(test2)
        let test3 = await messagebox("错误测试", "这是一个错误测试", 2, ["ok", "no"]);
        console.log(test3)
        let test4 = await messagebox("多按钮测试", "这是一个多按钮测试", 0, ["ok", "ok", "ok", "ok", "ok", "ok"]);
        console.log(test4)
    }
</script>
<div
        class="component-right"
        in:slide
        out:slide
        on:outroend={after_leave}>
    <MySelectCard title="可以折叠的卡片~" isExpand={true}>
        <div class="test">
            <MyLoadingPickaxe state={loading_state} loading_text={loading_text} />
            <div style="margin: 10px 0">
                <MyNormalButton style_in="width: 80px; height: 30px" click={() => {loading_state = false; loading_text = "加载中~"}}>加载中</MyNormalButton>
                <MyNormalButton style_in="width: 80px; height: 30px" click={() => {loading_state = true; loading_text = "加载失败~"}}>失败</MyNormalButton>
            </div>
            <MyNormalButton style_in="width: 170px; height: 30px" click={testMessageBox}>测试信息框</MyNormalButton>
        </div>
    </MySelectCard>
    <MySelectCard title="">
        <div class="test">
            <p>这是无标题的卡片</p>
            <p>无标题的卡片默认会无视isExpand，因为它不会折叠</p>
        </div>
    </MySelectCard>
    <MySelectCard title="带有标题的卡片">
        <div class="test">
            <p>这是带有标题的卡片</p>
            <p>卡片默认的isExpand是false，因此如果你想要显式的可折叠卡片，你需要手动将isExpand设为true</p>
        </div>
    </MySelectCard>
    <MySelectCard title="加载一个苹果~" isExpand={true}>
        <div class="test">
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
    .test {
        height: max-content;
        margin: 20px;
        width: calc(100% - 40px);
        display: flex;
        flex-direction: column;
        align-items: center;
    }
</style>