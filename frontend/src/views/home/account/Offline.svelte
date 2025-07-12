<script lang="ts">
    import {
        AccountPart,
        Alex,
        Ari,
        current_account_page,
        Efe,
        Kai,
        Makena,
        Noor,
        Steve,
        Sunny,
        Zuri
    } from "../../../store/changeBody";
    import MyNormalLabel from "../../../component/input/MyNormalLabel.svelte";
    import MyTextInput from "../../../component/input/MyTextInput.svelte";
    import MyRadioButton from "../../../component/button/MyRadioButton.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {GenerateBukkitUUID, UUIDToAvatar} from "../../../../wailsjs/go/launcher/MainMethod";
    import {launcher} from "../../../../wailsjs/go/models";
    import {SetAccountConfig} from "../../../../wailsjs/go/launcher/Account";
    import {messagebox} from "../../../store/messagebox";
    export let opacity = null
    export let after_leave = null
    let uuid_standard = true
    let username = ""
    let useruuid = ""
    let avatar = Steve
    async function usernameInput(name: string) {
        username = name
        if(!uuid_standard) { return }
        useruuid = await GenerateBukkitUUID(name)
        let arr = [Alex, Ari, Efe, Kai, Makena, Noor, Steve, Sunny, Zuri]
        let mod = await UUIDToAvatar(useruuid) % 18
        avatar = arr[mod >= 9 ? mod - 9 : mod]
    }
    async function useruuidInput(uuid: string) {
        useruuid = uuid
        const reg: RegExp = /^[a-f0-9]{32}$/g;
        if(reg.test(uuid)) {
            let mod = await UUIDToAvatar(uuid) % 18
            let arr = [Alex, Ari, Efe, Kai, Makena, Noor, Steve, Sunny, Zuri]
            avatar = arr[mod >= 9 ? mod - 9 : mod]
        }
    }
    async function createAccount() {
        const re1: RegExp = /^[a-zA-Z0-9_]{3,16}$/g
        const re2: RegExp = /^[a-f0-9]{32}$/g
        if(!re1.test(username)) {
            await messagebox("账户名称错误", "输入的账号名称错误，请输入英文状态下的英文数字和下划线，长度需要在3-16个之间。", 2, ['ok'])
            return
        }
        if(!re2.test(useruuid)) {
            await messagebox("账户 UUID 错误", "输入的账号 UUID 错误，请输入 32 位 16 进制的字符串，无需分隔符。", 2, ['ok'])
            return
        }
        $AccountPart.accounts.push(<launcher.AccountType>{
            type: "Offline",
            name: username,
            uuid: useruuid,
            head_skin: avatar
        })
        await SetAccountConfig($AccountPart)
        current_account_page.set(true)
    }
</script>
<div
        id="component"
        in:opacity
        out:opacity
        on:outroend={after_leave}
>
    <div id="center">
        <img src="data:image/png;base64,{avatar}" id="avatar" alt="头像" />
        <div class="table">
        <table>
            <tr style="">
                <td style="text-align: right;"><MyNormalLabel>玩家名称</MyNormalLabel></td>
                <td><MyTextInput placeholder="请输入用户名" style_in="width: 160px; height: 24px; margin-left: 4px" title="在 3 - 16 位之间，只能输入英文、数字和下划线。" handleInput={usernameInput}/></td>
            </tr>
            <tr>
                <td colspan="2">
                    <div style="display: flex; align-items: center; justify-content: space-around;">
                    <MyRadioButton isChecked={uuid_standard} click={() => uuid_standard = true}>行业规范</MyRadioButton>
                    <MyRadioButton isChecked={!uuid_standard} click={() => uuid_standard = false}>自定义</MyRadioButton>
                    </div>
                </td>
            </tr>
            <tr style:display={uuid_standard ? 'none' : ''}>
                <td style="text-align: right;"><MyNormalLabel>玩家 UUID</MyNormalLabel></td>
                <td><MyTextInput placeholder="请输入 UUID" style_in="width: 160px; height: 24px; margin-left: 4px" title="应为 32 位 16 进制字符串，不含连字符。" value={useruuid} handleInput={useruuidInput}/></td>
            </tr>
            <tr>
                <td colspan="2">
                    <div style="display: flex; align-items: center; justify-content: space-around;">
                        <MyNormalButton style_in="width: 100px; height: 30px" click={createAccount}>创建</MyNormalButton>
                        <MyNormalButton style_in="width: 100px; height: 30px" click={() => current_account_page.set(true)}>返回</MyNormalButton>
                    </div>
                </td>
            </tr>
        </table>
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
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 250px;
    }
    #avatar {
        width: 64px;
        image-rendering: pixelated;
        box-shadow: 0 0 6px gray;
        border-radius: 4px;
        margin-bottom: 16px;
    }
    .table {
        width: max-content;
    }
</style>