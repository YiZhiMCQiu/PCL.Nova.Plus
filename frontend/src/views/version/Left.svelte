<script lang="ts">
    import { onMount } from "svelte";
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import {
        GetConfigIniPath,
        OpenDirectoryDialog,
        ReadConfig,
        WriteConfig,
    } from "../../../wailsjs/go/launcher/ReaderWriter";
    import { messagebox, MSG_ERROR, showHint } from "../../store/messagebox";
    import { current_view } from "../../store/changeBody";
    import { inputbox } from "../../store/messagebox.js";
    import { current_mc_index, select_mc } from "../../store/mc";
    import {
        GetMCVersionConfig,
        SetMCVersionConfig,
    } from "../../../wailsjs/go/launcher/LaunchMethod";
    import { launcher } from "../../../wailsjs/go/models";

    export let width = "144px";
    export let slide = null;
    export let after_leave = null;

    onMount(async () => {
        select_mc.set([]);
        let v = await GetMCVersionConfig();
        if (!v.status) {
            await messagebox(
                "JSON 文件有误",
                "你擅自修改了 MCJson.json 文件，请立刻恢复原样！如果你不知道如何恢复原样，请尝试删除该文件后重试！错误信息：" +
                    v.message,
                MSG_ERROR,
            );
            current_view.set("home");
            return;
        }
        for (let i = 0; i < v.data.mc.length; i++) {
            select_mc.set([
                ...$select_mc,
                {
                    path: v.data.mc[i].path,
                    name: v.data.mc[i].name,
                },
            ]);
        }
        let m = await ReadConfig(await GetConfigIniPath(), "MC", "SelectMC");
        if (m == "" || parseInt(m) < 0) {
            return;
        }
        let num = parseInt(m);
        current_mc_index.set(num);
    });
    async function setMCVersionConfig(num: number) {
        await WriteConfig(
            await GetConfigIniPath(),
            "MC",
            "SelectMC",
            num.toString(),
        );
        current_mc_index.set(num);
    }
    async function openMCSelectFile() {
        let dirpath = await OpenDirectoryDialog("请选择 MC 根路径");
        if (dirpath == "") {
            return;
        }
        let dirname = await inputbox(
            "输入显示名称",
            "输入该文件夹在左边栏列表中显示的名称。",
            0,
            "请输入显示名称",
        );
        if (dirname == "") {
            return;
        }
        select_mc.set([
            ...$select_mc,
            {
                path: dirpath,
                name: dirname,
            },
        ]);
        await SetMCVersionConfig(
            launcher.MCConfigs.createFrom({ mc: $select_mc }),
        );
    }
</script>

<div
    class="component"
    style:width
    in:slide={{ x: Number(width.replace("px", "")) }}
    out:slide={{ x: Number(width.replace("px", "")) }}
    on:outroend={after_leave}
>
    <div class="grid">文件夹列表</div>
    <div class="file-list">
        <MySidebarButton
            isChecked={$current_mc_index === 0}
            on:click={() => {
                setMCVersionConfig(0);
            }}
        >
            <div style="margin-left: 10px;">当前文件夹</div>
        </MySidebarButton>
        {#each $select_mc as f, i}
            <MySidebarButton
                isChecked={$current_mc_index === i + 1}
                on:click={() => {
                    setMCVersionConfig(i + 1);
                }}
            >
                <div style="margin-left: 10px;">{f.name}</div>
            </MySidebarButton>
        {/each}
    </div>
    <div class="grid" style="margin-top: 15px">添加或导入</div>
    <div class="list">
        <MySidebarButton
            isChecked={false}
            on:click={() => {
                openMCSelectFile();
            }}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                fill="none"
                class="side-icon"
            >
                <path
                    d="M12 8v8m4-4H8m14 0c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10"
                />
            </svg>
            <span>添加已有文件夹</span>
        </MySidebarButton>
        <MySidebarButton
            isChecked={false}
            on:click={() => {
                showHint("目前导入整合包暂时还没有做好😭，请敬请期待吧！");
            }}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-linejoin="round"
                stroke-width="1.5"
                fill="none"
                class="side-icon"
            >
                <path
                    d="M10.691 4.562a6.19 6.19 0 0 1 6.545-1.42c.378.141.45.62.165.906l-2.787 2.787a1.037 1.037 0 0 0 0 1.467l1.084 1.084a1.037 1.037 0 0 0 1.467 0L19.953 6.6c.285-.285.764-.212.905.165a6.187 6.187 0 0 1-7.696 8.058c-.396-.128-.84-.054-1.134.24L6.481 20.61a2.186 2.186 0 1 1-3.09-3.09l5.547-5.548c.294-.294.368-.738.24-1.134a6.19 6.19 0 0 1 1.513-6.276Z"
                />
            </svg>
            <span>导入整合包</span>
        </MySidebarButton>
    </div>
</div>

<style>
    .component {
        height: 100%;
        display: flex;
        flex-direction: column;
    }
    .grid {
        margin-top: 5px;
        font-size: 12px;
        color: gray;
        margin-bottom: 7px;
        margin-left: 20px;
    }
    .list {
        width: 100%;
        height: max-content;
    }
</style>
