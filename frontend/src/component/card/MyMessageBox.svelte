<script lang="ts">
    import {b_button, b_content, b_level, b_resolve, b_show_all, b_title} from '../../store/messagebox'
    import {dark_mode} from '../../store/changeBody'
    import {quadInOut} from "svelte/easing";
    import MyNormalButton from "../button/MyNormalButton.svelte";
    import MyNormalLabel from "../input/MyNormalLabel.svelte";
    $: bg = $dark_mode ? "#282828cf" : "#f0f8ffcf"
    function getColors(level: number) {
        return {
            font_color: level == 0 ? "#3142b7cf" : level == 1 ? "#c7ad2acf" : "#ff4c4ccf",
            back_color: level == 0 ? '#0000005f' : level == 1 ? '#4f4f005f' : '#4f00005f',
        }
    }
    $: ({font_color, back_color} = getColors($b_level))
    let m_resolve = 0
    function traLeave() {
        b_title.set("")
        b_content.set("")
        b_level.set(0)
        b_button.set(["ok"])
        $b_resolve!(m_resolve)
        b_resolve.set(null)
    }
    function buttonClick(index: number) {
        m_resolve = index
        b_show_all.set(false)
    }
    function back_anim(node: HTMLElement) {
        return {
            duration: 330,
            easing: quadInOut,
            css(t: number) {
                return `opacity: ${t}`
            }
        }
    }
    function slide_anim(node: HTMLElement) {
        return {
            duration: 330,
            easing: quadInOut,
            css(t: number) {
                const rotate = -20 * (1 - t);
                if (t < 0.8) {
                    const progress = t / 0.8;
                    return `
            opacity: ${progress * 0.8};
            transform: translate(-50%, ${-20 - 30 * progress}%) scale(${0.9 + 0.2 * progress}) rotate(${rotate}deg);
          `;
                } else {
                    const progress = (t - 0.8) / 0.2;
                    return `
            opacity: ${0.8 + 0.2 * progress};
            transform: translate(-50%, -50%) scale(${1.1 - 0.1 * progress}) rotate(${rotate}deg);
          `;
                }
            }
        }
    }
</script>
{#if $b_show_all}
    <div id="back" class={b_show_all ? 'back-class' : 'back-class-hide'} in:back_anim out:back_anim style="--m-back-color: {back_color};"></div>
    <div class="content-box-class" in:slide_anim out:slide_anim on:outroend={traLeave} style="--m-font-color: {font_color}; --bg-color: {bg}">
        <div id="content-title">{$b_title}</div>
        <div id="content">
            <MyNormalLabel>{@html $b_content}</MyNormalLabel>
        </div>
        {#each $b_button as b, i}
            <MyNormalButton style_in="width: max-content; min-width: 50px; height: 30px; margin: 10px; float: right; font-weight: bold;" click={() => {buttonClick(i)}}>
                { b === "ok" ? "确认" : b === "cancel" ? "取消" : b === "yes" ? "是" : b === "no" ? "否" : b }
            </MyNormalButton>
        {/each}
    </div>
{/if}
<style>
    #content-title {
        margin: 10px;
        font-size: 30px;
        font-weight: bold;
        border-bottom: 3px solid var(--m-font-color);
        color: var(--m-font-color);
    }

    #content {
        margin: 10px;
        padding: 5px;
    }

    #back {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        z-index: 1000;
        --wails-draggable: drag;
    }

    .back-class {
        background-color: var(--m-back-color);
        backdrop-filter: blur(3px);
        transition: all 0.33s;
    }

    .back-class-hide {
        backdrop-filter: blur(0);
        background-color: rgba(0, 0, 0, 0);
        transition: all 0.33s;
    }

    .content-box-class {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        min-width: 200px;
        min-height: 50px;
        width: max-content;
        max-width: 800px;
        background-color: var(--bg-color);
        border-radius: 20px;
        border: 1px solid black;
        z-index: 1001;
        padding: 10px 20px;
        transition: box-shadow 0.2s;
    }
    .content-box-class:hover {
        box-shadow: 0 0 10px var(--bg-color);
    }
</style>