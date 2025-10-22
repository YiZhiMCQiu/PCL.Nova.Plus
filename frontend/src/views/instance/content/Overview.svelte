<script lang="ts">
    export let slide = null;
    export let after_leave = null;
    import MyCardButton from "../../../component/button/MyCardButton.svelte";
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import { onMount } from "svelte";
    import Grass from "../../../assets/images/Blocks/Grass.png";
    import { current_mc_version_path, game_log } from "../../../store/mc";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {
        messagebox,
        MSG_ERROR,
        MSG_INFO,
        showHint,
    } from "../../../store/messagebox";
    import { LaunchGame } from "../../../../wailsjs/go/launcher/LaunchMethod";
    import { GetArgsDir } from "../../../../wailsjs/go/launcher/MainMethod";
    import {
        GetCurrentExeDir,
        OpenExplorer,
    } from "../../../../wailsjs/go/launcher/ReaderWriter";
    onMount(() => {
        // Initialize or perform any necessary actions here
    });
    function loadVersionName(): string {
        let path = $current_mc_version_path;
        let i1 = path.lastIndexOf("/");
        let i2 = path.lastIndexOf("\\");
        if (i1 < 0 && i2 < 0) {
            return "Unknown";
        }
        if (i1 >= 0) {
            return path.substring(i1 + 1);
        } else if (i2 >= 0) {
            return path.substring(i2 + 1);
        }
    }
    function loadVersionPath(): string {
        return $current_mc_version_path;
    }
    async function exportLaunchArgs() {
        let isExportAccessToken = true;
        if (
            (await messagebox(
                "是否为 Access Token 打码？",
                "请选择是否要为 Access token 打码？<br>如果不选择打码，你可以随意分发你的 启动参数，但是无法使用正版登录！<br>但是如果你选择不打码，只是想做个快捷启动放到桌面，那么你可以选择不打码，但是这样的话！你绝对不可以将你的 启动参数 发给任何一个人！否则他们将拥有使用你的正版账号的权限！",
                MSG_INFO,
                ["yes", "no"],
            )) == 0
        ) {
            isExportAccessToken = false;
        } else {
            isExportAccessToken = true;
        }
        showHint("正在导出启动参数中~");
        let l = await LaunchGame(true, isExportAccessToken);
        if (!l.status) {
            await messagebox(
                "导出失败！",
                "导出启动参数时出现了错误！<br>错误信息：" +
                    l.message +
                    "<br>错误描述" +
                    l.data,
                MSG_ERROR,
            );
        } else {
            showHint("导出参数成功！");
            await OpenExplorer(await GetArgsDir());
        }
        game_log.set("");
    }
</script>

<div class="component-overview" in:slide out:slide on:outroend={after_leave}>
    <MySelectCard title="">
        <div class="version-all">
            <MyCardButton
                image={Grass}
                title={loadVersionName()}
                desc={loadVersionPath()}
            />
        </div>
    </MySelectCard>
    <MySelectCard title="高级管理">
        <div class="version-all">
            <MyNormalButton
                style_in="width: 130px; height: 35px;"
                on:click={exportLaunchArgs}>导出启动参数</MyNormalButton
            >
        </div>
    </MySelectCard>
</div>

<style>
    .component-overview {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
</style>
