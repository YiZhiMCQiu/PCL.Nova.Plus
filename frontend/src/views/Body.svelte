<script lang="ts">
    import {current_view, dark_mode} from "../store/changeBody";
    import HomeLeft from './home/Left.svelte'
    import DownloadLeft from './download/Left.svelte'
    import OnlineLeft from './online/Left.svelte'
    import SettingLeft from './setting/Left.svelte'
    import AboutLeft from './about/Left.svelte'
    import HomeRight from './home/Right.svelte'
    import DownloadRight from './download/Right.svelte'
    import OnlineRight from './online/Right.svelte'
    import SettingRight from './setting/Right.svelte'
    import AboutRight from './about/Right.svelte'
    import {quadInOut} from "svelte/easing";
    import {onDestroy} from "svelte";

    let isTransitioning = true
    let left_width = "330px"
    let right_height = "330px"
    $: dark = $dark_mode ? "rgba(32, 32, 32, 0.6)" : "rgba(255, 255, 255, 0.8)"

    function control_leave() {
        right_height = $current_view === "home" ? "330px" : $current_view === "setting" ? "110px" : "144px"
        isTransitioning = true
    }
    // 用了点奇妙的小魔法解决了一点奇妙的小问题（（
    // 能跑就行，别动以下部分了（除非你非常确保你的逻辑比我的清晰并且比我的更加简便。
    // 最主要的是你的逻辑能正常跑，否则我不建议你动以下部分。
    // 如果你实在有更好的想法，请不仅修改此处，你也要修改所有需要显式声明动画的组件里的模块。【例如
    // 1. home/Left.svelte
    // 2. 其余view组件/Right.svelte
    // 】
    let f = false
    const unsubscribe_current_view = current_view.subscribe((value) => {
        left_width = value === "home" ? "330px" : value === "setting" ? '110px' : '144px'
        if (!f) {
            f = true
            isTransitioning = true;
        }else{
            isTransitioning = !isTransitioning
        }
    })
    function slide_left(node: HTMLElement, param: { x: number }) {
        const x = param.x || 144
        return {
            duration: 150,
            easing: quadInOut,
            css: (t: number, n: number) => {
                return `
                  transform: translateX(${-x * n}px);
                  opacity: ${t};
                `;
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
    onDestroy(unsubscribe_current_view)
</script>
<div>
    <div id="left" style:width={left_width} style:background-color={dark}>
        {#if $current_view === "home" && isTransitioning}
            <HomeLeft width={left_width} slide={slide_left} after_leave={control_leave} />
        {:else if $current_view === "download" && isTransitioning}
            <DownloadLeft width={left_width} slide={slide_left} after_leave={control_leave}/>
        {:else if $current_view === "online" && isTransitioning}
            <OnlineLeft width={left_width} slide={slide_left} after_leave={control_leave}/>
        {:else if $current_view === "setting" && isTransitioning}
            <SettingLeft width={left_width} slide={slide_left} after_leave={control_leave}/>
        {:else if $current_view === "about" && isTransitioning}
            <AboutLeft width={left_width} slide={slide_left} after_leave={control_leave}/>
        {/if}
    </div>
    <div id="right" style:left={right_height} style:width="calc(100% - {right_height})">
        {#if $current_view === "home" && isTransitioning}
            <HomeRight after_leave={control_leave} slide={slide_up}/>
        {:else if $current_view === "download" && isTransitioning}
            <DownloadRight after_leave={control_leave} slide={slide_up}/>
        {:else if $current_view === "online" && isTransitioning}
            <OnlineRight after_leave={control_leave} slide={slide_up}/>
        {:else if $current_view === "setting" && isTransitioning}
            <SettingRight after_leave={control_leave} slide={slide_up}/>
        {:else if $current_view === "about" && isTransitioning}
            <AboutRight after_leave={control_leave} slide={slide_up}/>
        {/if}
    </div>
</div>
<style>
    #left {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        transition: all 0.3s;
    }
    #right {
        position: absolute;
        top: 0;
        height: 100%;
        background-color: transparent;
    }
</style>
