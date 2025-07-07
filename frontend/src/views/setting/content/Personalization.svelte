<script lang="ts">
    import MyToggleSwitch from "../../../component/button/MyToggleSwitch.svelte";

    export let slide = null
    export let after_leave = null
    import { theme_mode, dark_mode } from "../../../logic/changeBody";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import MyNormalLabel from "../../../component/input/MyNormalLabel.svelte";
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import MyRadioButton from "../../../component/button/MyRadioButton.svelte";
    import {OpenCustomURL} from "../../../logic/functions";
    import {GetConfigIniPath, WriteConfig} from "../../../../wailsjs/go/launcher/ReaderWriter";

    async function changeTheme(index: number) {
        theme_mode.set(index)
        await WriteConfig(await GetConfigIniPath(), "Misc", "ThemeMode", $theme_mode.toString())
    }
    async function toggleDark(is: boolean) {
        dark_mode.set(!is)
        await WriteConfig(await GetConfigIniPath(), "Misc", "DarkMode", $dark_mode ? '1' : '0')
    }
    let isSelect = false
</script>
<div
        class="component-person"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <MySelectCard title="主题" isExpand={false}>
        <div class="proc" style="position: relative;">
            <div id="mask">
                <MyNormalLabel style_in="font-weight: bold">请支持官方版本以使用主题功能</MyNormalLabel>
                <MyNormalButton style_in="width: 200px; height: 40px" click={() => OpenCustomURL('https://afdian.com/a/LTCat')}>点我进入龙猫的爱发电</MyNormalButton>
            </div>
            <div class="line">
                <MyRadioButton isChecked={$theme_mode === 1} click={() => changeTheme(1)}>龙猫蓝</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 2} click={() => changeTheme(2)}>甜柠青</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 3} click={() => changeTheme(3)}>小草绿</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 4} click={() => changeTheme(4)}>菠萝黄</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 5} click={() => changeTheme(5)}>橡木棕</MyRadioButton>
            </div>
            <div class="line">
                <MyRadioButton isChecked={$theme_mode === 6} click={() => changeTheme(6)}>玄素黑</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 7} click={() => changeTheme(7)}>滑稽彩</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 8} click={() => changeTheme(8)}>铁杆粉</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 9} click={() => changeTheme(9)}>神秘紫</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 10} click={() => changeTheme(10)}>欧皇彩</MyRadioButton>
            </div>
            <div class="line">
                <MyRadioButton isChecked={$theme_mode === 11} click={() => changeTheme(11)}>秋仪金</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 12} click={() => changeTheme(12)}>活跃橙</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 13} click={() => changeTheme(13)}>跳票红</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 14} click={() => changeTheme(14)}>极客蓝</MyRadioButton>
                <MyRadioButton isChecked={$theme_mode === 15} click={() => changeTheme(15)}>自定义</MyRadioButton>
            </div>
        </div>
    </MySelectCard>
    <MySelectCard title="个性化" isExpand={true}>
        <div class="proc">
            <div id="dark">
                <MyNormalLabel style_in="font-weight: bold">暗色模式</MyNormalLabel>
                <MyToggleSwitch isSelect={$dark_mode} click={() => toggleDark($dark_mode)} />
            </div>
        </div>
    </MySelectCard>
</div>
<style>
    .component-person {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .line {
        width: 100%;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: space-around;
    }
    #mask {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        backdrop-filter: blur(2px);
        z-index: 2;
    }
    #dark {
        width: 100%;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
</style>