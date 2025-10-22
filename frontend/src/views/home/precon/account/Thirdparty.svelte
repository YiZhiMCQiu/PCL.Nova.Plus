<script lang="ts">
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyTextInput from "../../../../component/input/MyTextInput.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {
        current_account_page,
        select_account,
    } from "../../../../store/changeBody";
    import {
        HNT_PASS,
        messagebox,
        MSG_ERROR,
        MSG_INFO,
        MSG_WARNING,
        showHint,
    } from "../../../../store/messagebox";
    import {
        copyToClipBoard,
        GetLittleSkinUrl,
        OpenCustomURL,
        sleep,
    } from "../../../../store/functions";
    import MyCustomBox from "../../../../component/card/MyCustomBox.svelte";
    import MyTextButton from "../../../../component/button/MyTextButton.svelte";
    import MyLoadingPickaxe from "../../../../component/card/MyLoadingPickaxe.svelte";
    import {
        GetThirdAPAccessToken,
        GetThirdHeadSkin,
        GetThirdOAuthAccessToken,
        GetUserCodeThirdOAuth,
        SetAccountConfig,
    } from "../../../../../wailsjs/go/launcher/AccountMethod";
    import { launcher } from "../../../../../wailsjs/go/models";

    export let opacity = null;
    export let after_leave = null;
    let isOpen = false;
    let loading_text = "";
    let loading_state = false;
    let buttonShow = false;
    let inputShow = "";
    let urlText = "";
    async function startOAuthLittleSkinLogin() {
        let geoUrl = GetLittleSkinUrl();
        loading_text = "正在获取 User Code 中……";
        loading_state = false;
        isOpen = true;
        buttonShow = false;
        let user_codes = await GetUserCodeThirdOAuth(geoUrl[1]);
        if (!user_codes.status) {
            loading_text =
                "获取 User Code 失败，错误信息：" + user_codes.message;
            loading_state = true;
            buttonShow = true;
            inputShow = "";
            return;
        }
        loading_text = "正在登录中……";
        let user_code = user_codes.data[0];
        let device_code = user_codes.data[1];
        urlText = geoUrl[0] + "?user_code=" + user_code;
        inputShow = user_code;
        copyToClipBoard(user_code);
        let accessToken = null as launcher.AccountType;
        for (let i = 0; i < 36; i++) {
            let acc = await GetThirdOAuthAccessToken(geoUrl[2], device_code);
            if (acc.code == 401) {
                await sleep(5000);
            } else if (acc.status) {
                accessToken = acc.data;
                break;
            } else {
                loading_text = "获取失败，错误信息：" + acc.message;
                loading_state = true;
                buttonShow = true;
                inputShow = "";
                return;
            }
        }
        if (!accessToken) {
            loading_text = "登陆已过期，请重试！";
            loading_state = true;
            buttonShow = true;
            inputShow = "";
            return;
        }
        inputShow = "";
        loading_text = "登陆成功！正在获取用户大头像~";
        let head_skin = await GetThirdHeadSkin(geoUrl[3], accessToken.uuid);
        if (!head_skin.status) {
            loading_text = "获取大头像失败！错误信息：" + head_skin.message;
            loading_state = true;
            buttonShow = true;
            inputShow = "";
            return;
        }
        accessToken.head_skin = head_skin.data;
        accessToken.server = geoUrl[3];
        select_account.set([...$select_account, accessToken]);
        await SetAccountConfig(
            launcher.AccountList.createFrom({ accounts: $select_account }),
        );
        current_account_page.set(true);
        isOpen = false;
        showHint(`添加成功😀！玩家名称：${accessToken.name}！`, HNT_PASS);
    }
    let server = "";
    function serverInput(event: CustomEvent) {
        server = event.detail.value;
    }
    let username = "";
    function usernameInput(event: CustomEvent) {
        username = event.detail.value;
    }
    let password = "";
    function passwordInput(event: CustomEvent) {
        password = event.detail.value;
    }
    async function startLogin() {
        let a = await GetThirdAPAccessToken(server, username, password);
        if (!a.status) {
            await messagebox(
                "账号登陆失败",
                "账号登陆失败，错误信息：" + a.message,
                MSG_ERROR,
            );
            return;
        }
        let accessToken = a.data;
        if (a.data.length == 0) {
            await messagebox(
                "未选择皮肤嗷",
                "你还暂未选择一个皮肤，请去官网上先选择一个皮肤再登陆吧~",
                MSG_WARNING,
            );
            return;
        } else if (a.data.length == 1) {
            showHint("登陆成功！正在获取用户大头像~");
            let userProfile = accessToken[0];
            let head_skin = await GetThirdHeadSkin(server, userProfile.uuid);
            if (!head_skin.status) {
                showHint("获取大头像失败，错误信息：" + head_skin.message);
                return;
            }
            userProfile.head_skin = head_skin.data;
            userProfile.server = server;
            select_account.set([...$select_account, userProfile]);
            await SetAccountConfig(
                launcher.AccountList.createFrom({ accounts: $select_account }),
            );
            current_account_page.set(true);
            showHint(`添加成功😀！玩家名称：${userProfile.name}！`, HNT_PASS);
        } else {
            let m = await messagebox(
                "请选择你需要的登陆的玩家：",
                "请在下方选择你需要登陆的玩家！",
                MSG_INFO,
                [...accessToken.map((item) => item.name)],
            );
            let userProfile = accessToken[m];
            let head_skin = await GetThirdHeadSkin(server, userProfile.uuid);
            if (!head_skin.status) {
                showHint("获取大头像失败，错误信息：" + head_skin.message);
                return;
            }
            userProfile.head_skin = head_skin.data;
            userProfile.server = server;
            select_account.set([...$select_account, userProfile]);
            await SetAccountConfig(
                launcher.AccountList.createFrom({ accounts: $select_account }),
            );
            current_account_page.set(true);
            showHint(`添加成功😀！玩家名称：${userProfile.name}！`, HNT_PASS);
        }
    }
</script>

<div id="component" in:opacity out:opacity on:outroend={after_leave}>
    <div id="center">
        <table>
            <tr>
                <td colspan="2">
                    <div
                        style="display: flex; flex-direction: column; align-items: center; justify-content: center"
                    >
                        <MyNormalButton
                            style_in="width: 200px; height: 30px"
                            on:click={startOAuthLittleSkinLogin}
                            >使用 OAuth 登陆 LittleSkin</MyNormalButton
                        >
                    </div>
                </td>
            </tr>
            <tr>
                <td colspan="2">
                    <h5
                        style="display: flex; justify-content: center; margin: 0"
                    >
                        <MyNormalSpan>下面是手动登陆</MyNormalSpan>
                    </h5>
                </td>
            </tr>
            <tr>
                <td><MyNormalSpan>服务器</MyNormalSpan></td>
                <td
                    ><MyTextInput
                        placeholder="请输入服务器地址"
                        title="后面必须跟着 /api/yggdrasil，否则登录不成功！"
                        style_in="width: 160px; height: 24px;"
                        value={server}
                        on:blur={serverInput}
                    />
                </td>
            </tr>
            <tr>
                <td><MyNormalSpan>账号</MyNormalSpan></td>
                <td
                    ><MyTextInput
                        placeholder="请输入账号"
                        title="多半是邮箱地址"
                        style_in="width: 160px; height: 24px;"
                        on:blur={usernameInput}
                    />
                </td>
            </tr>
            <tr>
                <td><MyNormalSpan>密码</MyNormalSpan></td>
                <td
                    ><MyTextInput
                        password
                        placeholder="请输入密码"
                        title="输入密码即可"
                        style_in="width: 160px; height: 24px;"
                        on:blur={passwordInput}
                    />
                </td>
            </tr>
            <tr>
                <td colspan="2">
                    <div
                        style="display: flex; align-items: center; justify-content: space-around;"
                    >
                        <MyNormalButton
                            style_in="width: 100px; height: 30px"
                            on:click={startLogin}>登录</MyNormalButton
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
    <MyCustomBox {isOpen}>
        <div
            style="width: 600px; height: 400px; display: flex; flex-direction: column; align-items: center; justify-content: center; position: relative;"
        >
            <MyLoadingPickaxe
                {loading_text}
                state={loading_state}
                style_in="max-width: 500px"
            />
            {#if inputShow !== ""}
                <MyNormalSpan>你的用户代码是：</MyNormalSpan>
                <MyNormalSpan
                    style_in="user-select: all;
                    -webkit-user-select: all;
                    -moz-user-select: all;
                    -ms-user-select: all; font-size: 30px"
                    >{inputShow}</MyNormalSpan
                >
                <MyNormalSpan>请打开该链接：</MyNormalSpan>
                <MyTextButton on:click={() => OpenCustomURL(urlText)}>
                    <MyNormalSpan style_in="text-decoration: underline;">
                        {urlText}
                    </MyNormalSpan>
                </MyTextButton>
                <MyNormalSpan>
                    并输入代码以用于验证~验证通过后会自然的为你执行下一步操作！
                </MyNormalSpan>
            {/if}
            {#if buttonShow}
                <MyNormalButton
                    style_in="position: absolute; width: 100px; height: 40px; right: 20px; bottom: 20px"
                    on:click={() => (isOpen = false)}
                >
                    返回
                </MyNormalButton>
            {/if}
        </div>
    </MyCustomBox>
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
