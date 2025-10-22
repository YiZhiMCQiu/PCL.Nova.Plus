<script lang="ts">
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {
        current_account_page,
        select_account,
    } from "../../../../store/changeBody";
    import {
        copyToClipBoard,
        OpenCustomURL,
        sleep,
    } from "../../../../store/functions";
    import MyCustomBox from "../../../../component/card/MyCustomBox.svelte";
    import MyLoadingPickaxe from "../../../../component/card/MyLoadingPickaxe.svelte";
    import {
        GetAccessToken,
        GetUserCode,
        GetUserHeadSkin,
        SetAccountConfig,
    } from "../../../../../wailsjs/go/launcher/AccountMethod";
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyTextButton from "../../../../component/button/MyTextButton.svelte";
    import { launcher } from "../../../../../wailsjs/go/models";
    import { HNT_PASS, showHint } from "../../../../store/messagebox";
    // import AccountType from launcher.AccountType;

    export let opacity = null;
    export let after_leave = null;
    let isOpen = false;
    let loading_text = "";
    let loading_state = false;
    let buttonShow = false;
    let inputShow = "";
    async function startLogin() {
        loading_text = "正在获取 User Code 中……";
        loading_state = false;
        isOpen = true;
        buttonShow = false;
        let user_codes = await GetUserCode();
        if (!user_codes.status) {
            loading_text =
                "获取 User Code 失败，错误信息：" + user_codes.message;
            loading_state = true;
            buttonShow = true;
            inputShow = "";
            return;
        }
        loading_text = "正在登录中~";
        let user_code = user_codes.data[0];
        let device_code = user_codes.data[1];
        inputShow = user_code;
        copyToClipBoard(user_code);
        let accessToken = null as launcher.AccountType;
        for (let i = 0; i < 36; i++) {
            let acc = await GetAccessToken(device_code);
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
        loading_text = "登录成功！正在获取用户大头像~";
        let head_skin = await GetUserHeadSkin(accessToken.uuid);
        if (!head_skin.status) {
            loading_text = "获取大头像失败！错误信息：" + head_skin.message;
            loading_state = true;
            buttonShow = true;
            inputShow = "";
            return;
        }
        accessToken.head_skin = head_skin.data;
        select_account.set([...$select_account, accessToken]);
        await SetAccountConfig(
            launcher.AccountList.createFrom({ accounts: $select_account }),
        );
        current_account_page.set(true);
        isOpen = false;
        showHint(`添加成功😀！玩家名称：${accessToken.name}！`, HNT_PASS);
    }
</script>

<div id="component" in:opacity out:opacity on:outroend={after_leave}>
    <div id="center">
        <svg
            role="img"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
            fill="none"
            id="login-img"
        >
            <path
                d="M4,20 C4,17 8,17 10,15 C11,14 8,14 8,9 C8,5.667 9.333,4 12,4 C14.667,4 16,5.667 16,9 C16,14 13,14 14,15 C16,17 20,17 20,20"
            />
        </svg>
        <div id="button-group">
            <MyNormalButton
                style_in="width: 100px; height: 30px; border: 1px solid skyblue;"
                on:click={startLogin}
            >
                添加新账号
            </MyNormalButton>
            <MyNormalButton
                style_in="width: 80px; height: 30px; border: 1px solid skyblue;"
                on:click={() => {
                    OpenCustomURL(
                        "https://www.xbox.com/zh-cn/games/store/minecraft-java-bedrock-edition-for-pc/9nxp44l49shj",
                    );
                }}
            >
                购买正版
            </MyNormalButton>
            <MyNormalButton
                style_in="width: 50px; height: 30px; border: 1px solid skyblue;"
                on:click={() => {
                    current_account_page.set(true);
                }}
                >返回
            </MyNormalButton>
        </div>
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
                    >{inputShow}
                </MyNormalSpan>
                <MyNormalSpan>请打开该链接：</MyNormalSpan>
                <MyTextButton
                    on:click={() =>
                        OpenCustomURL("https://www.microsoft.com/link")}
                >
                    <MyNormalSpan style_in="text-decoration: underline;">
                        https://www.microsoft.com/link
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
                    >返回
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
