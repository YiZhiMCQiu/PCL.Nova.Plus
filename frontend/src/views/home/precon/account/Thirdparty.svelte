<script lang="ts">
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyTextInput from "../../../../component/input/MyTextInput.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {current_account_page} from "../../../../store/changeBody";
    import {showHint} from "../../../../store/messagebox";

    export let opacity = null
    export let after_leave = null

    let server = ""
    function serverInput(event: CustomEvent) {
        server = event.detail.value
    }
    let username = ""
    function usernameInput(event: CustomEvent) {
        username = event.detail.value
    }
    let password = ""
    function passwordInput(event: CustomEvent) {
        password = event.detail.value
    }
    function startLogin() {
        // console.log(server)
        // console.log(username)
        // console.log(password)
        showHint("目前外置登录暂时还没有做好😭，请敬请期待吧！")
    }
</script>
<div
        id="component"
        in:opacity
        out:opacity
        on:outroend={after_leave}
>
    <div id="center">
        <table>
            <tr>
                <td colspan="2">
                    <div style="display: flex; align-items: center; justify-content: center">
                        <MyNormalButton style_in="width: 200px; height: 30px" on:click={() => {server = 'https://littleskin.cn/api/yggdrasil'}}>应用 Littleskin 服务器</MyNormalButton>
                    </div>
                </td>
            </tr>
            <tr>
                <td><MyNormalSpan>服务器</MyNormalSpan></td>
                <td><MyTextInput placeholder="请输入服务器地址" title="后面必须跟着 /api/yggdrasil，否则登录不成功！" style_in="width: 160px; height: 24px;" value={server} on:blur={serverInput}/></td>
            </tr>
            <tr>
                <td><MyNormalSpan>账号</MyNormalSpan></td>
                <td><MyTextInput placeholder="请输入账号" title="多半是邮箱地址" style_in="width: 160px; height: 24px;" on:blur={usernameInput} /></td>
            </tr>
            <tr>
                <td><MyNormalSpan>密码</MyNormalSpan></td>
                <td><MyTextInput placeholder="请输入密码" title="输入密码即可" style_in="width: 160px; height: 24px;" on:blur={passwordInput} /></td>
            </tr>
            <tr>
                <td colspan="2">
                    <div style="display: flex; align-items: center; justify-content: space-around;">
                        <MyNormalButton style_in="width: 100px; height: 30px" on:click={startLogin}>登录</MyNormalButton>
                        <MyNormalButton style_in="width: 100px; height: 30px" on:click={() => current_account_page.set(true)}>返回</MyNormalButton>
                    </div>
                </td>
            </tr>
        </table>
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
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 200px;
    }
    #center table {
        border-spacing: 10px;
    }
    #center table tr td {
        text-align: right;
    }
</style>