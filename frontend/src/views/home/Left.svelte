<script lang="ts">
    import MyNormalButton from "../../component/button/MyNormalButton.svelte";
    import {current_account_page, current_view} from "../../store/changeBody";
    import {quadInOut} from "svelte/easing";
    import {onDestroy, onMount} from "svelte";
    import AccountSelect from "./precon/AccountSelect.svelte";
    import AddAccount from "./precon/AddAccount.svelte";
    import {
        GetCurrentMinecraftDir,
        GetMCAllVersion,
        GetMCVersionConfig, LaunchGame
    } from "../../../wailsjs/go/launcher/LaunchMethod";
    import {GetConfigIniPath} from "../../../wailsjs/go/launcher/ReaderWriter";
    import {ReadConfig} from "../../../wailsjs/go/launcher/ReaderWriter.js";
    import {current_mc_version_path/* , game_log */} from "../../store/mc";
    import {HNT_PASS, messagebox, MSG_ERROR, showHint} from "../../store/messagebox";
    import {EventsOn} from "../../../wailsjs/runtime";

    export let width = "144px"
    export let slide = null
    export let after_leave = null
    let isTransitioning = true
    let f = false
    let version_name = ""

    function control_leave() {
        isTransitioning = true
    }

    const unsubscribe_current_view = current_account_page.subscribe((value) => {
        if (!f) {
            f = true
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning
        }
    })

    function account_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `opacity: ${t};`;
            }
        };
    }

    onDestroy(unsubscribe_current_view)
    onMount(async () => {
        version_name = "暂未选择任一核心~"
        if ($current_mc_version_path != "") {
            let path = $current_mc_version_path
            let i1 = path.lastIndexOf("/")
            let i2 = path.lastIndexOf("\\")
            if (i1 < 0 && i2 < 0) {
                return
            }
            if (i1 >= 0) {
                version_name = path.substring(i1 + 1)
            } else if (i2 >= 0) {
                version_name = path.substring(i2 + 1)
            }
        } else {
            try {
                let v = await GetMCVersionConfig()
                if(!v.status) {
                    return
                }
                let mci = parseInt(await ReadConfig(await GetConfigIniPath(), "MC", "SelectMC"))
                if (Number.isNaN(mci) || mci < 0 || mci > v.data.mc.length) {
                    return
                }
                let rp = mci == 0 ? await GetCurrentMinecraftDir() : v.data.mc[mci - 1].path
                let p2 = await GetMCAllVersion(rp)
                let j = parseInt(await ReadConfig(await GetConfigIniPath(), "MC", "SelectVer"))
                if (Number.isNaN(j) || j < 0 || j > p2.length - 1) {
                    return
                }
                let path = p2[j]
                let i1 = path.lastIndexOf("/")
                let i2 = path.lastIndexOf("\\")
                if (i1 < 0 && i2 < 0) {
                    return
                }
                if (i1 >= 0) {
                    version_name = path.substring(i1 + 1)
                } else if (i2 >= 0) {
                    version_name = path.substring(i2 + 1)
                }
            } catch (_) {
            }
        }
    })
    async function launchGame() {
        showHint("游戏正在启动，请稍后~")
        let l = await LaunchGame()
        if(!l.status) {
            await messagebox("游戏出错了！", "启动游戏时出现了错误！错误信息：" + l.message + "<br>" + "错误描述：" + l.data, MSG_ERROR)
        } else {
            showHint("游戏已结束，玩得愉快！", HNT_PASS)
        }
    }
    EventsOn("launch_success", () => {
        showHint("参数拼接成功啦！正在等待启动游戏嗷~")
    })
</script>
<div
        class="component"
        style:width={width}
        in:slide={{ x: Number(width.replace("px", "")) }}
        out:slide={{ x: Number(width.replace("px", "")) }}
        on:outroend={after_leave}
>
    <div style="display: flex; flex-direction: column; height: 100%">
        <div id="middle">
            {#if $current_account_page && isTransitioning}
                <AccountSelect opacity={account_opacity} after_leave={control_leave}/>
            {:else if !$current_account_page && isTransitioning}
                <AddAccount opacity={account_opacity} after_leave={control_leave}/>
            {/if}
        </div>
        <div id="bottom">
            <MyNormalButton
                    style_in="height: 75px; width: calc(100% - 52px); border: 1px solid #216fbd; margin-top: 6px"
                    isDisabled={!$current_account_page} on:click={launchGame}>
                <span id="launch-title">启动游戏</span><br>
                <span id="launch-version">{version_name}</span>
            </MyNormalButton>
            <div id="setting">
                <MyNormalButton style_in="width: calc(50% - 4px); height: 40px;" isDisabled={!$current_account_page}
                                on:click={() => {
                    current_view.set("version")
                }}>
                    选择核心
                </MyNormalButton>
                <MyNormalButton style_in="width: calc(50% - 4px); height: 40px;" isDisabled={!$current_account_page}
                                on:click={() => {
                    showHint("目前独立核心设置暂时还没有做好😭，请敬请期待吧！")
                }}>
                    核心设置
                </MyNormalButton>
            </div>
        </div>
    </div>
</div>
<style>
    .component {
        height: 100%;
    }

    #middle {
        width: 100%;
        flex: 1;
    }

    #bottom {
        width: 100%;
        height: 156px;
        flex-shrink: 0;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    #launch-title {
        background-image: linear-gradient(to right, rgb(63, 207, 255), rgb(96, 96, 255));
        color: transparent;
        background-clip: text;
        font-size: 20px;
    }

    #launch-version {
        font-size: 13px;
    }

    #setting {
        height: 54px;
        width: calc(100% - 52px);
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
    }
</style>