<script lang="ts">
    import {
        select_account,
        Alex,
        Ari,
        current_account_page,
        Efe,
        Kai,
        Makena,
        Noor,
        Steve,
        Sunny,
        Zuri,
    } from "../../../../store/changeBody";
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyTextInput from "../../../../component/input/MyTextInput.svelte";
    import MyRadioButton from "../../../../component/button/MyRadioButton.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {
        GenerateBukkitUUID,
        UUIDToAvatar,
    } from "../../../../../wailsjs/go/launcher/MainMethod";
    import { launcher } from "../../../../../wailsjs/go/models";
    import { SetAccountConfig } from "../../../../../wailsjs/go/launcher/AccountMethod";
    import {
        HNT_PASS,
        messagebox,
        MSG_ERROR,
        showHint,
    } from "../../../../store/messagebox";
    export let opacity = null;
    export let after_leave = null;
    let uuid_standard = true;
    let username = "";
    let useruuid = "";
    let avatar = Steve;
    async function usernameInput(event: CustomEvent) {
        username = event.detail.value;
        if (!uuid_standard) {
            return;
        }
        useruuid = await GenerateBukkitUUID(username);
        let arr = [Alex, Ari, Efe, Kai, Makena, Noor, Steve, Sunny, Zuri];
        let mod = (await UUIDToAvatar(useruuid)) % 18;
        avatar = arr[mod >= 9 ? mod - 9 : mod];
    }
    async function useruuidInput(event: CustomEvent) {
        useruuid = event.detail.value;
        const reg: RegExp = /^[a-f0-9]{32}$/g;
        if (reg.test(useruuid)) {
            let mod = (await UUIDToAvatar(useruuid)) % 18;
            let arr = [Alex, Ari, Efe, Kai, Makena, Noor, Steve, Sunny, Zuri];
            avatar = arr[mod >= 9 ? mod - 9 : mod];
        }
    }
    async function createAccount() {
        const re1: RegExp = /^[a-zA-Z0-9_]{3,16}$/gi;
        const re2: RegExp = /^[a-f0-9]{32}$/gi;
        if (!re1.test(username)) {
            await messagebox(
                "账户名称错误",
                "输入的账号名称错误，请输入英文状态下的英文数字和下划线，长度需要在3-16个之间。",
                MSG_ERROR,
            );
            return;
        }
        if (!re2.test(useruuid)) {
            await messagebox(
                "账户 UUID 错误",
                "输入的账号 UUID 错误，请输入 32 位 16 进制的字符串，无需分隔符。",
                MSG_ERROR,
            );
            return;
        }
        select_account.set([
            ...$select_account,
            {
                type: "Offline",
                name: username,
                uuid: useruuid,
                head_skin: avatar,
            },
        ]);
        await SetAccountConfig(
            launcher.AccountList.createFrom({ accounts: $select_account }),
        );
        current_account_page.set(true);
        showHint(`添加成功😀！玩家名称：${username}！`, HNT_PASS);
    }
</script>

<div id="component" in:opacity out:opacity on:outroend={after_leave}>
    <div id="center">
        <img src="data:image/png;base64,{avatar}" id="avatar" alt="头像" />
        <div class="table">
            <table>
                <tr style="">
                    <td style="text-align: right;"
                        ><MyNormalSpan>玩家名称</MyNormalSpan></td
                    >
                    <td
                        ><MyTextInput
                            placeholder="请输入用户名"
                            style_in="width: 160px; height: 24px; margin-left: 4px"
                            title="在 3 - 16 位之间，只能输入英文、数字和下划线。"
                            on:input={usernameInput}
                        /></td
                    >
                </tr>
                <tr>
                    <td colspan="2">
                        <div
                            style="display: flex; align-items: center; justify-content: space-around;"
                        >
                            <MyRadioButton
                                isChecked={uuid_standard}
                                on:click={() => (uuid_standard = true)}
                                >行业规范</MyRadioButton
                            >
                            <MyRadioButton
                                isChecked={!uuid_standard}
                                on:click={() => (uuid_standard = false)}
                                >自定义</MyRadioButton
                            >
                        </div>
                    </td>
                </tr>
                <tr style:display={uuid_standard ? "none" : ""}>
                    <td style="text-align: right;"
                        ><MyNormalSpan>玩家 UUID</MyNormalSpan></td
                    >
                    <td
                        ><MyTextInput
                            placeholder="请输入 UUID"
                            style_in="width: 160px; height: 24px; margin-left: 4px"
                            title="应为 32 位 16 进制字符串，不含连字符。"
                            value={useruuid}
                            on:blur={useruuidInput}
                        /></td
                    >
                </tr>
                <tr>
                    <td colspan="2">
                        <div
                            style="display: flex; align-items: center; justify-content: space-around;"
                        >
                            <MyNormalButton
                                style_in="width: 100px; height: 30px"
                                on:click={createAccount}>创建</MyNormalButton
                            >
                            <MyNormalButton
                                style_in="width: 100px; height: 30px"
                                on:click={() => current_account_page.set(true)}
                                >返回</MyNormalButton
                            >
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
        box-shadow: 0 0 6px gray;
        border-radius: 4px;
        margin-bottom: 16px;
    }
    .table {
        width: max-content;
    }
</style>
