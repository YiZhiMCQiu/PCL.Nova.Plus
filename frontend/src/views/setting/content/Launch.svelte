<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import MyToggleSwitch from "../../../component/button/MyToggleSwitch.svelte";
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    import MyTextInput from "../../../component/input/MyTextInput.svelte";
    import MyRadioButton from "../../../component/button/MyRadioButton.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import { onDestroy, onMount } from "svelte";
    import { current_java_index, select_java } from "../../../store/mc";
    import {
        HNT_PASS,
        messagebox,
        MSG_ERROR,
        showHint,
    } from "../../../store/messagebox";
    import {
        GetJavaConfig,
        GetJavaInfo,
        SetJavaConfig,
    } from "../../../../wailsjs/go/launcher/LaunchMethod";
    import Java from "../../../assets/images/Icons/Java.png";
    import {
        GetConfigIniPath,
        OpenFileDialog,
        ReadConfig,
        WriteConfig,
    } from "../../../../wailsjs/go/launcher/ReaderWriter";
    import { slide_opacity, slide_up } from "../../../store/functions";
    import MyLoadingPickaxe from "../../../component/card/MyLoadingPickaxe.svelte";
    import { launcher } from "../../../../wailsjs/go/models";
    import MyProgressBar from "../../../component/input/MyProgressBar.svelte";
    import { GetAvailableMemory } from "../../../../wailsjs/go/launcher/MainMethod.js";
    import {
        GetJavaExecutableFileName,
        GetTotalMemory,
    } from "../../../../wailsjs/go/launcher/MainMethod";
    import { current_setting, current_view } from "../../../store/changeBody";
    import MyCheckButton from "../../../component/button/MyCheckButton.svelte";

    export let slide = null;
    export let after_leave = null;
    let loading_state = false;
    let loading_text = "";
    let isTransitioning = true;
    let f = false;
    const unsubscribe_select_java = select_java.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_select_java);

    function control_leave() {
        isTransitioning = true;
    }

    let availableMemory = 0;
    let totalMemory = 0;
    let currentMemory = 0;
    let currentNum = 0;
    let isIsolation = true;
    let customInfo = "";
    let isCheckLibraries = true;
    let winWidth = 854;
    let winHeight = 480;
    let additionalJVM = "";
    let additionalGame = "";

    async function reloadSettings() {
        isIsolation =
            (await ReadConfig(
                await GetConfigIniPath(),
                "Version",
                "SelectIsolation",
            )) == "4";
        customInfo = await ReadConfig(
            await GetConfigIniPath(),
            "Version",
            "CustomInfo",
        );
        let w = parseInt(
            await ReadConfig(
                await GetConfigIniPath(),
                "Document",
                "WindowWidth",
            ),
        );
        if (!Number.isNaN(w) && w >= 854) {
            winWidth = w;
        }
        let h = parseInt(
            await ReadConfig(
                await GetConfigIniPath(),
                "Document",
                "WindowHeight",
            ),
        );
        if (!Number.isNaN(h) && h >= 480) {
            winHeight = h;
        }
        additionalJVM = await ReadConfig(
            await GetConfigIniPath(),
            "Document",
            "AdditionalJVM",
        );
        additionalGame = await ReadConfig(
            await GetConfigIniPath(),
            "Document",
            "AdditionalGame",
        );
        let icl = await ReadConfig(
            await GetConfigIniPath(),
            "Document",
            "IsCheckLibraries",
        );
        if (icl == "") {
            isCheckLibraries = true;
            await WriteConfig(
                await GetConfigIniPath(),
                "Document",
                "IsCheckLibraries",
                "1",
            );
        } else {
            isCheckLibraries = icl == "1";
        }
    }
    async function reloadMemory() {
        if (totalMemory == 0) {
            totalMemory = await GetTotalMemory();
            currentNum = parseInt(
                await ReadConfig(
                    await GetConfigIniPath(),
                    "Document",
                    "MaxMemoryLevel",
                ),
            );
            if (Number.isNaN(currentNum) || currentNum > 32 || currentNum < 0) {
                currentNum = 24;
            }
            currentMemory = Math.round(totalMemory / (32 - currentNum + 1));
        }
    }

    setInterval(async () => {
        if ($current_view == "setting" && $current_setting == "Launch") {
            availableMemory = await GetAvailableMemory();
        }
    }, 1000);

    async function reloadJava() {
        if ($select_java.length <= 0) {
            loading_text = "正在加载 Java 中~";
            loading_state = false;
            let v = await GetJavaConfig();
            if (!v.status) {
                await messagebox(
                    "JSON 文件有误",
                    "你擅自修改了 JavaJson.json 文件，错误信息：" + v.message,
                    MSG_ERROR,
                );
                return;
            }
            for (let i = 0; i < v.data.java.length; i++) {
                select_java.set([
                    ...$select_java,
                    {
                        path: v.data.java[i].path,
                        version: v.data.java[i].version,
                        arch: v.data.java[i].arch,
                        vendor: v.data.java[i].vendor,
                    },
                ]);
            }
        }
        if ($select_java.length <= 0) {
            loading_text =
                "你好像还暂未导入任一 Java，请点击 手动添加 导入一个再来吧~";
            loading_state = true;
        }
        let javaIndex = parseInt(
            await ReadConfig(await GetConfigIniPath(), "Java", "SelectJava"),
        );
        if (
            Number.isNaN(javaIndex) ||
            javaIndex < 0 ||
            javaIndex >= $select_java.length
        ) {
            return;
        }
        current_java_index.set(javaIndex);
    }

    onMount(() => {
        reloadSettings();
        reloadMemory();
        reloadJava();
    });

    async function selectJava(index: number) {
        if ($current_java_index != index) {
            current_java_index.set(index);
            await WriteConfig(
                await GetConfigIniPath(),
                "Java",
                "SelectJava",
                index.toString(),
            );
        }
    }

    async function deleteJava(index: number) {
        if (index < $current_java_index) {
            current_java_index.update((value) => value - 1);
        } else if (index == $current_java_index) {
            current_java_index.set(-1);
        }
        select_java.update((value) => {
            value.splice(index, 1);
            return value;
        });
        await SetJavaConfig(
            launcher.JavaConfigs.createFrom({ java: $select_java }),
        );
        if ($select_java.length <= 0) {
            loading_text =
                "你好像还暂未导入任一 Java，请点击 手动添加 导入一个再来吧~";
            loading_state = true;
        }
    }

    async function addJava() {
        let executable = await GetJavaExecutableFileName();
        let java = await OpenFileDialog("请选择 Java 路径", executable);
        if (java == "") {
            return;
        }
        let javaConfig = await GetJavaInfo(java);
        if (!javaConfig.status) {
            await messagebox(
                "输入的 Java 有误",
                "您输入的 Java 有误，请重新输入！错误信息：" +
                    javaConfig.message,
                MSG_ERROR,
            );
            return;
        }
        select_java.set([...$select_java, javaConfig.data]);
        await SetJavaConfig(
            launcher.JavaConfigs.createFrom({ java: $select_java }),
        );
        showHint("添加成功😀", HNT_PASS);
    }

    async function onBarDragging(event: CustomEvent) {
        let num = event.detail.value;
        currentMemory = Math.round(totalMemory / (33 - num));
        currentNum = num;
        await WriteConfig(
            await GetConfigIniPath(),
            "Document",
            "MaxMemoryLevel",
            num.toString(),
        );
    }

    async function toggleIsolation() {
        isIsolation = !isIsolation;
        await WriteConfig(
            await GetConfigIniPath(),
            "Version",
            "SelectIsolation",
            isIsolation ? "4" : "1",
        );
    }

    async function customInfoInput(event: CustomEvent) {
        customInfo = event.detail.value;
        await WriteConfig(
            await GetConfigIniPath(),
            "Version",
            "CustomInfo",
            customInfo,
        );
    }

    async function widthInput(event: CustomEvent) {
        let v = parseInt(event.detail.value);
        if (!Number.isNaN(v) && v >= 854) {
            winWidth = v;
            await WriteConfig(
                await GetConfigIniPath(),
                "Document",
                "WindowWidth",
                v.toString(),
            );
        }
    }

    async function heightInput(event: CustomEvent) {
        let v = parseInt(event.detail.value);
        if (!Number.isNaN(v) && v >= 480) {
            winHeight = v;
            await WriteConfig(
                await GetConfigIniPath(),
                "Document",
                "WindowHeight",
                v.toString(),
            );
        }
    }
    async function toggleFullScreen() {
        if (additionalGame.indexOf("--fullScreen") >= 0) {
            additionalGame = additionalGame.replaceAll("--fullScreen", "");
        } else {
            additionalGame += " --fullScreen";
        }
        additionalGame = additionalGame.trim();
        await WriteConfig(
            await GetConfigIniPath(),
            "Document",
            "AdditionalGame",
            additionalGame,
        );
    }

    async function chooseIPPreference(index: number) {
        if (index == 1) {
            if (additionalJVM.indexOf("-Djava.preferIPv4Stack=true") >= 0) {
                return;
            }
            additionalJVM = additionalJVM.replaceAll(
                "-Djava.preferIPv6Stack=true",
                "",
            );
            additionalJVM += " -Djava.preferIPv4Stack=true";
        } else if (index == 2) {
            additionalJVM = additionalJVM.replaceAll(
                "-Djava.preferIPv4Stack=true",
                "",
            );
            additionalJVM = additionalJVM.replaceAll(
                "-Djava.preferIPv6Stack=true",
                "",
            );
        } else {
            if (additionalJVM.indexOf("-Djava.preferIPv6Stack=true") >= 0) {
                return;
            }
            additionalJVM = additionalJVM.replaceAll(
                "-Djava.preferIPv4Stack=true",
                "",
            );
            additionalJVM += " -Djava.preferIPv6Stack=true";
        }
        additionalJVM = additionalJVM.trim();
        await WriteConfig(
            await GetConfigIniPath(),
            "Document",
            "AdditionalJVM",
            additionalJVM,
        );
    }
    async function handleAdditionalInput(value: string, num: number) {
        if (num == 1) {
            additionalJVM = value;
            await WriteConfig(
                await GetConfigIniPath(),
                "Document",
                "AdditionalJVM",
                additionalJVM,
            );
        } else {
            additionalGame = value;
            await WriteConfig(
                await GetConfigIniPath(),
                "Document",
                "AdditionalJVM",
                additionalGame,
            );
        }
    }
    async function changeIsCheckLibraries() {
        isCheckLibraries = !isCheckLibraries;
        await WriteConfig(
            await GetConfigIniPath(),
            "Document",
            "IsCheckLibraries",
            isCheckLibraries ? "1" : "0",
        );
    }
</script>

<div class="component-launch" in:slide out:slide on:outroend={after_leave}>
    {#if isTransitioning}
        <div in:slide_up out:slide_up on:outroend={control_leave}>
            <MySelectCard title="启动选项">
                <div class="proc">
                    <div class="settings">
                        <MyNormalSpan
                            style_in="width: 100px"
                            title="如果选中则默认所有版本均隔离，否则均不隔离。（建议在加 Mod 的时候开启全部隔离）"
                        >
                            默认版本隔离
                        </MyNormalSpan>
                        <MyToggleSwitch
                            isSelect={isIsolation}
                            on:click={toggleIsolation}
                        ></MyToggleSwitch>
                    </div>
                    <div class="settings">
                        <MyNormalSpan style_in="width: 100px"
                            >自定义信息</MyNormalSpan
                        >
                        <MyTextInput
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="该信息会显示在 MC 主界面左下角，与游戏中按 F3 调试界面的左上角"
                            placeholder="请输入自定义信息"
                            value={customInfo}
                            on:blur={customInfoInput}
                        />
                    </div>
                    <div class="settings">
                        <MyNormalSpan style_in="width: 100px"
                            >窗口大小
                        </MyNormalSpan>
                        <MyTextInput
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            placeholder="宽"
                            title="宽度"
                            value={winWidth.toString()}
                            on:blur={widthInput}
                        />&nbsp;×&nbsp;
                        <MyTextInput
                            style_in="flex: 1; height: 25px"
                            placeholder="高"
                            title="高度"
                            value={winHeight.toString()}
                            on:blur={heightInput}
                        />
                        <MyNormalSpan
                            style_in="width: 100px; margin-left: 20px"
                            title="任意一个框填入【0】，则默认全屏"
                            >是否全屏
                        </MyNormalSpan>
                        <MyToggleSwitch
                            isSelect={additionalGame.indexOf("--fullScreen") >=
                                0}
                            on:click={toggleFullScreen}
                            title="开启这个即忽略宽高属性，直接全屏启动"
                        ></MyToggleSwitch>
                    </div>
                    <div class="settings">
                        <MyNormalSpan
                            style_in="width: 100px"
                            title={'通过设置 Java 虚拟机参数来设置 MC 的 IP 协议版本偏好。\n一般建议设置为 "Java 默认"，如果在将来更新了联机大厅，或许需要确保大厅正常工作而设置 IPv4 优先。\n如果你目前需要体验 Nova 的 IPv6 检测联机，你需要设置成 IPv6 优先。'}
                        >
                            IP 协议偏好
                        </MyNormalSpan>
                        <MyRadioButton
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="将添加额外 JVM 参数：-Djava.net.preferIPv4Stack=true"
                            isChecked={additionalJVM.indexOf(
                                "-Djava.preferIPv4Stack=true",
                            ) >= 0}
                            on:click={() => chooseIPPreference(1)}
                            >IPv4 优先
                        </MyRadioButton>
                        <MyRadioButton
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="将不会添加任何额外 JVM 参数"
                            isChecked={additionalJVM.indexOf(
                                "-Djava.preferIPv4Stack=true",
                            ) < 0 &&
                                additionalJVM.indexOf(
                                    "-Djava.preferIPv6Stack=true",
                                ) < 0}
                            on:click={() => chooseIPPreference(2)}
                            >Java 默认
                        </MyRadioButton>
                        <MyRadioButton
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="将添加额外 JVM 参数：-Djava.net.preferIPv6Stack=true"
                            isChecked={additionalJVM.indexOf(
                                "-Djava.preferIPv6Stack=true",
                            ) >= 0}
                            on:click={() => chooseIPPreference(3)}
                            >IPv6 优先
                        </MyRadioButton>
                    </div>
                </div>
            </MySelectCard>
            <MySelectCard title="Java 管理">
                <div class="version-all">
                    <MyNormalButton
                        style_in="width: 100px; height: 30px"
                        title="让用户手动添加一个 Java，自主选择 java.exe 或者 javaw.exe"
                        on:click={addJava}
                    >
                        手动添加
                    </MyNormalButton>
                    <MyNormalButton
                        style_in="width: 100px; height: 30px; margin-left: 20px"
                        on:click={() => {
                            showHint(
                                "目前 Java 浅搜索暂时还没有做好😭，请敬请期待吧！",
                            );
                        }}
                        title={"浅层搜索 Java\nNova 只会从以下路径按顺序开始遍历：\n\nWindows：\n注册表\n64 位文件夹的 Java 目录\n32 位文件夹的 Java 目录\n官启目录\nNova 手动安装的目录\n\nMacOS：\n注册表\n/Library/Java\n/usr/local/opt\n~/Library/Java\n\nLinux：\n/usr/lib\n/usr/java\n/usr/local/java\n/opt/java"}
                    >
                        浅搜索
                    </MyNormalButton>
                    <MyNormalButton
                        style_in="width: 100px; height: 30px; margin-left: 20px"
                        on:click={() => {
                            showHint(
                                "目前 Java 深搜索暂时还没有做好😭，请敬请期待吧！",
                            );
                        }}
                        title={"深层搜索 Java\nNova 会尝试遍历你的整个文件系统，以最全面的形式找到你系统里所有可能存在的 Java。\n该举动会导致扫盘，可能会很慢，如果没有必要，请不要使用这个。除非你真的忘记了你的 Java 安装路径。"}
                    >
                        深搜索
                    </MyNormalButton>
                </div>
                <div class="version-all">
                    {#if $select_java.length > 0}
                        {#each $select_java as java, i}
                            <div
                                class="a-java"
                                title="Java 路径：{java.path}&#13;Java 架构：{java.arch}&#13;Java 版本：{java.version}&#13;Java 发行商：{java.vendor}"
                                on:click={() => {
                                    selectJava(i);
                                }}
                                on:keydown|preventDefault
                                style={$current_java_index === i
                                    ? "cursor: default"
                                    : "cursor: pointer"}
                            >
                                <MyRadioButton
                                    isChecked={i === $current_java_index}
                                    style_in="margin-left: 5px"
                                />
                                <img
                                    src={Java}
                                    alt="Java"
                                    class="a-java-icon"
                                />
                                <div class="info" style="pointer-events: none">
                                    <MyNormalSpan>{java.version}</MyNormalSpan>
                                    <div style="font-size: 13px; color: gray">
                                        <span class="code">{java.arch}</span
                                        >&nbsp;<span class="code"
                                            >{java.vendor}</span
                                        >&nbsp;<span>{java.path}</span>
                                    </div>
                                </div>
                                <button
                                    class="a-delete cursor-pointer"
                                    on:click|stopPropagation={() =>
                                        deleteJava(i)}
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 24 24"
                                        fill="none"
                                        stroke-linecap="round"
                                        stroke-width="1.5"
                                    >
                                        <path
                                            d="M9.17 4a3.001 3.001 0 0 1 5.66 0m5.67 2h-17m15.333 2.5l-.46 6.9c-.177 2.654-.265 3.981-1.13 4.79s-2.196.81-4.856.81h-.774c-2.66 0-3.991 0-4.856-.81c-.865-.809-.954-2.136-1.13-4.79l-.46-6.9M9.5 11l.5 5m4.5-5l-.5 5"
                                        />
                                    </svg>
                                </button>
                            </div>
                        {/each}
                    {:else}
                        <div
                            style="width: 100%; height: 100%; display: flex; align-items: center; justify-content: center;"
                            in:slide_opacity
                            out:slide_opacity
                            on:outroend={control_leave}
                        >
                            <MyLoadingPickaxe
                                {loading_text}
                                state={loading_state}
                            />
                        </div>
                    {/if}
                </div>
            </MySelectCard>
            <MySelectCard title="游戏内存">
                <div class="proc">
                    <div style="width: 100%; height: 50px">
                        <MyProgressBar
                            max={32}
                            min={0}
                            on:dragging={onBarDragging}
                            value={currentNum}
                        />
                    </div>
                    <MyNormalSpan
                        >已安装内存：{totalMemory}MB，剩余内存：{availableMemory}MB，游戏分配：{currentMemory}MB
                    </MyNormalSpan>
                </div>
            </MySelectCard>
            <MySelectCard title="高级启动设置" isExpand={true} canExpand={true}>
                <div class="version-all">
                    <div class="settings">
                        <MyNormalSpan style_in="width: 120px"
                            >额外 JVM 参数</MyNormalSpan
                        >
                        <MyTextInput
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="启动 Minecraft 时使用的额外 JVM 参数，除非有确定把握，否则请不要修改。"
                            placeholder="请输入额外 JVM 参数"
                            value={additionalJVM}
                            on:blur={(e) =>
                                handleAdditionalInput(e.detail.value, 1)}
                        />
                    </div>
                    <div class="settings">
                        <MyNormalSpan style_in="width: 120px"
                            >额外游戏参数</MyNormalSpan
                        >
                        <MyTextInput
                            style_in="flex: 1; margin-left: 20px; height: 25px"
                            title="文本框中的内容将会被直接拼合在启动参数的末尾。&#13;例如，输入 --demo 则会以试玩模式启动游戏。"
                            placeholder="请输入额外游戏参数"
                            value={additionalGame}
                            on:blur={(e) =>
                                handleAdditionalInput(e.detail.value, 2)}
                        />
                    </div>
                    <div class="normal-setting">
                        <MyCheckButton
                            isChecked={isCheckLibraries}
                            on:click={changeIsCheckLibraries}
                            title="如果这个复选框打开，则默认会校验 Libraries，反之则不会校验（"
                            >是否校验 Libraries</MyCheckButton
                        >
                    </div>
                </div>
            </MySelectCard>
        </div>
    {/if}
</div>

<style>
    .component-launch {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }

    .settings {
        width: 100%;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .a-java {
        position: relative;
        transition: all 0.2s;
        border-radius: 10px;
        height: 50px;
        width: calc(100% - 2px);
        flex-shrink: 0;
        display: flex;
        align-items: center;
    }

    .a-java-icon {
        width: 40px;
        height: 40px;
        image-rendering: pixelated;
        border-radius: 10px;
        box-shadow: 0 0 5px gray;
        margin-left: 5px;
    }

    .a-java:hover {
        background-color: rgba(128, 128, 128, 0.5);
    }

    .a-java:active {
        transform: scale(99%);
    }

    .info {
        display: flex;
        flex-direction: column;
        margin-left: 10px;
    }

    .a-delete {
        position: absolute;
        top: 10px;
        right: 30px;
        width: 30px;
        height: 30px;
        background-color: transparent;
        border: 0;
        border-radius: 50%;
        stroke: gray;
        transition: all 0.2s;
    }

    .a-delete:hover {
        background-color: rgba(0, 0, 0, 0.2);
        stroke: red;
    }

    .a-delete svg {
        width: 20px;
        height: 20px;
        vertical-align: middle;
    }

    .code {
        font-size: 12px;
        background-color: rgba(0, 0, 0, 0.2);
        padding: 0 5px;
        border-radius: 4px;
    }
    .normal-setting {
        width: 100%;
        height: 40px;
    }
</style>
