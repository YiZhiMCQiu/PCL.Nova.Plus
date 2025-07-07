<script lang="ts">
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {current_account_page} from "../../../logic/changeBody";
    import {OpenCustomURL} from "../../../logic/functions";

    export let opacity = null
    export let after_leave = null
</script>
<div
        id="component"
        in:opacity
        out:opacity
        on:outroend={after_leave}
>
    <div id="center">
        <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" id="login-img">
            <path
                    d="M4,20 C4,17 8,17 10,15 C11,14 8,14 8,9 C8,5.667 9.333,4 12,4 C14.667,4 16,5.667 16,9 C16,14 13,14 14,15 C16,17 20,17 20,20"/>
        </svg>
        <div id="button-group">
            <MyNormalButton style_in="width: 100px; height: 30px; border: 1px solid skyblue;" click={() => {current_account_page.set(true)}}>添加新账号</MyNormalButton>
            <MyNormalButton style_in="width: 80px; height: 30px; border: 1px solid skyblue;" click={() => {OpenCustomURL("https://www.xbox.com/zh-cn/games/store/minecraft-java-bedrock-edition-for-pc/9nxp44l49shj")}}>购买正版</MyNormalButton>
            <MyNormalButton style_in="width: 50px; height: 30px; border: 1px solid skyblue;" click={() => {current_account_page.set(true)}}>返回</MyNormalButton>
        </div>
    </div>
</div>
<style>
    #component {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
    }
    #center {
        width: 100%;
        height: 120px;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    #login-img {
        stroke: rgb(0, 186, 254);
        width: 80px;
        height: 80px;
    }
    #button-group {
        display: flex;
        justify-content: space-around;
        width: calc(100% - 50px);
    }
</style>