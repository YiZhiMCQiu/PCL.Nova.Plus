<script lang="ts">
    import {i_content, i_level, i_resolve, i_show_all, i_title, i_placeholder} from '../../store/messagebox'
    import {dark_mode} from '../../store/changeBody'
    import {quadInOut} from "svelte/easing";
    import MyTextInput from "../input/MyTextInput.svelte";
    import MyNormalLabel from "../input/MyNormalLabel.svelte";
    import MyNormalButton from "../button/MyNormalButton.svelte";
    $: bg = $dark_mode ? "#282828cf" : "#f0f8ffcf"
    function getColors(level: number) {
        return {
            font_color: level == 0 ? "#3142b7cf" : level == 1 ? "#c7ad2acf" : "#ff4c4ccf",
            back_color: level == 0 ? '#0000005f' : level == 1 ? '#4f4f005f' : '#4f00005f',
        }
    }
    $: ({font_color, back_color} = getColors($i_level))
    let m_resolve = ""
    function traLeave() {
        i_title.set("")
        i_content.set("")
        i_level.set(0)
        $i_resolve!(m_resolve)
        i_resolve.set(null)
        m_resolve = ""
    }
    function textInput(e: string) {
        m_resolve = e
    }
    function buttonClick(num: number) {
        if(num == 0){
            m_resolve = ""
        }
        i_show_all.set(false)
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
{#if $i_show_all}
    <div id="back" class={i_show_all ? 'back-class' : 'back-class-hide'} in:back_anim out:back_anim style="--m-back-color: {back_color};"></div>
    <div class="content-box-class" in:slide_anim out:slide_anim on:outroend={traLeave} style="--m-font-color: {font_color}; --bg-color: {bg}">
        <div id="content-title">{$i_title}</div>
        <div id="content">
            <MyNormalLabel>{@html $i_content}</MyNormalLabel>
        </div>
        <div id="input-box">
            <MyTextInput placeholder={$i_placeholder} style_in="width: 100%; height: 30px" handleInput={textInput}/>
        </div>
        <MyNormalButton style_in="width: max-content; min-width: 50px; height: 30px; margin: 10px; float: right; font-weight: bold;" click={() => {buttonClick(1)}}>确认</MyNormalButton>
        <MyNormalButton style_in="width: max-content; min-width: 50px; height: 30px; margin: 10px; float: right; font-weight: bold;" click={() => {buttonClick(0)}}>取消</MyNormalButton>
    </div>
{/if}
<style>
    #content-title {
        margin: 10px;
        font-size: 30px;
        font-weight: bold;
        border-bottom: 3px solid var(--m-font-color);
        padding: 5px;
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
        background-color: var(--bg-color);
        border-radius: 20px;
        border: 1px solid black;
        z-index: 1001;
        transition: box-shadow 0.2s;
    }
    .content-box-class:hover {
        box-shadow: 0 0 10px var(--bg-color);
    }
    #input-box {
        margin: 5px 30px 5px 10px;
    }
</style>