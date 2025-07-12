<script lang="ts">
    import MyNormalLabel from "../../../component/input/MyNormalLabel.svelte";
    import MyTextInput from "../../../component/input/MyTextInput.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {current_account_page} from "../../../store/changeBody";

    export let opacity = null
    export let after_leave = null

    let server = ""
    function serverInput(s: string) {
        server = s
    }
    let username = ""
    function usernameInput(name: string) {
        username = name
    }
    let password = ""
    function passwordInput(pass: string) {
        password = pass
    }
    function startLogin() {
        console.log(server)
        console.log(username)
        console.log(password)
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
                        <MyNormalButton style_in="width: 200px; height: 30px" click={() => {server = 'https://littleskin.cn/api/yggdrasil'}}>应用 Littleskin 服务器</MyNormalButton>
                    </div>
                </td>
            </tr>
            <tr>
                <td><MyNormalLabel>服务器</MyNormalLabel></td>
                <td><MyTextInput placeholder="请输入服务器地址" title="后面必须跟着 /api/yggdrasil，否则登录不成功！" style_in="width: 160px; height: 24px;" value={server} handleInput={serverInput}/></td>
            </tr>
            <tr>
                <td><MyNormalLabel>账号</MyNormalLabel></td>
                <td><MyTextInput placeholder="请输入账号" title="多半是邮箱地址" style_in="width: 160px; height: 24px;" handleInput={usernameInput} /></td>
            </tr>
            <tr>
                <td><MyNormalLabel>密码</MyNormalLabel></td>
                <td><MyTextInput placeholder="请输入密码" title="输入密码即可" style_in="width: 160px; height: 24px;" handleInput={passwordInput} /></td>
            </tr>
            <tr>
                <td colspan="2">
                    <div style="display: flex; align-items: center; justify-content: space-around;">
                        <MyNormalButton style_in="width: 100px; height: 30px" click={startLogin}>登录</MyNormalButton>
                        <MyNormalButton style_in="width: 100px; height: 30px" click={() => current_account_page.set(true)}>返回</MyNormalButton>
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