<script lang="ts">
    import {OpenCustomURL} from "../../../../store/functions";
    import MySelectCard from "../../../../component/card/MySelectCard.svelte";
    import MyNormalLabel from "../../../../component/input/MyNormalLabel.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {fabric_list} from "../../../../store/mc";
    import {quadInOut} from "svelte/easing";
    import {onDestroy, onMount} from "svelte";
    import MyLoadingPickaxe from "../../../../component/card/MyLoadingPickaxe.svelte";
    import {HttpGet} from "../../../../../wailsjs/go/launcher/Network";
    import MyCardButton from "../../../../component/button/MyCardButton.svelte";
    import Fabric from '../../../../assets/images/Blocks/Fabric.png'
    let loading_text = ""
    let loading_state = false
    export let slide = null
    export let after_leave = null
    onMount(async () => {
        if($fabric_list.length <= 1) {
            loading_text = "正在获取 Fabric 列表"
            loading_state = false
            let meta = await HttpGet("https://meta.fabricmc.net/v2/versions/installer", "")
            if(meta == "") {
                loading_text = "网络环境不佳，请稍后重试，或使用 VPN 以改善网络环境"
                loading_state = true
                return
            }
            let json = JSON.parse(meta)
            for(let i = 0; i < json.length; i++) {
                fabric_list.update((value) => [...value, {
                    url: json[i].url,
                    version: json[i].version,
                    stable: json[i].stable
                }])
            }
        }
    })
    let isTransitioning = true
    function control_leave() {
        isTransitioning = true
    }
    let f = false
    const unsubscribe_fabric_list = fabric_list.subscribe((value) => {
        if (!f) {
            f = true
            isTransitioning = true
        } else {
            isTransitioning = !isTransitioning
        }
    })
    function slide_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `opacity: ${t};`;
            }
        };
    }
    function slide_up(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number, n: number) => {
                return `
                    transform: translateY(${-50 * n}%);
                    opacity: ${t};
                `
            }
        }
    }
    onDestroy(unsubscribe_fabric_list)
    function slide_button_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `
                    transform: scale(${t});
                `
            }
        }
    }
    let scrollStep = 100
    let scrollTop = 0
    function onDivScroll(e: Event) {
        let target = e.target as HTMLElement
        scrollTop = target.scrollTop
    }
    function onTopDiv() {
        let target = document.getElementsByClassName("component-fabric")[0] as HTMLElement
        let s = setInterval(function() {
            target.scrollTop -= scrollStep
            scrollTop -= scrollStep
            if(scrollTop <= scrollStep) {
                scrollTop = 0
                target.scrollTop = 0
                clearInterval(s)
            }
        }, 10)
    }
</script>
<div
        class="component-fabric"
        in:slide
        out:slide
        on:outroend={after_leave}
        on:scroll={onDivScroll}
>
    <div style="width: 100%; height: max-content">
        <MySelectCard title="Fabric 简介" maxHeight={150}>
            <div class="proc" style="align-items: start; margin-left: 40px">
                <div>
                    <MyNormalLabel>Fabric Loader 是新版 Minecraft 下的轻量化 Mod 加载器，你需要先安装它才能安装各种 Fabric 模组</MyNormalLabel><br>
                    <MyNormalLabel>本页面提供 Fabric 安装器下载，在下载后你需要手动打开安装器进行安装。</MyNormalLabel>
                </div>
                <div style="margin-top: 10px"><MyNormalButton style_in="width: 150px; height: 40px; border: 1px solid skyblue" click={() => {OpenCustomURL("https://www.fabricmc.net/")}}>打开官网</MyNormalButton></div>
            </div>
        </MySelectCard>
    </div>
    {#if $fabric_list.length <= 1 && isTransitioning}
        <div style="display: flex; align-items: center; justify-content: center;" in:slide_opacity out:slide_opacity on:outroend={control_leave}>
            <MyLoadingPickaxe loading_text={loading_text} state={loading_state}/>
        </div>
    {:else if isTransitioning}
        <div
                in:slide_up
                out:slide_up
                on:outroend={control_leave}>
            <MySelectCard title="版本列表 ({$fabric_list.length})">
                <div class="version-all">
                    {#each $fabric_list as list}
                        <MyCardButton
                                image={Fabric}
                                title={list.version}
                                desc={list.stable ? '正式版' : '测试版'}
                                click={() => {}}
                                buttons={[]}
                        />
                    {/each}
                </div>
            </MySelectCard>
        </div>
        {#if scrollTop >= 1000}
            <button id="topButton" on:click={onTopDiv} in:slide_button_opacity out:slide_button_opacity>
                <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="3">
                    <path d="M5 3l15 0M12 21l0 -13.5M12 7l7 7M12 7l-7 7" />
                </svg>
            </button>
        {/if}
    {/if}
</div>
<style>
    #topButton {
        position: fixed;
        right: 20px;
        bottom: 20px;
        width: 50px;
        height: 50px;
        border-radius: 50%;
        background-color: #00aaffcf;
        border: 0;
        transition: all 0.2s;
    }
    #topButton:hover {
        background-color: #0077ffcf;
    }
    #topButton svg {
        stroke: white;
        width: 25px;
        height: 25px;
        vertical-align: middle;
    }
    .component-fabric {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .component-fabric {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .component-fabric > div {
        width: 100%;
        height: calc(100% - 203px);
    }
    .component-fabric > div:last-child {
        height: calc(100% - 223px);
    }
</style>