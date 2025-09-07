<script lang="ts">
  import NavBar from './views/NavBar.svelte'
  import Body from './views/Body.svelte'
  import {dark_mode, dont_click, theme_mode} from "./store/changeBody";
  import {onMount} from "svelte";
  import {GetConfigIniPath, ReadConfig} from "../wailsjs/go/launcher/ReaderWriter";
  import {GetBackgroundImage} from "../wailsjs/go/launcher/MainMethod";
  import {DarkAndThemeToMain} from "./store/functions";
  import MyMessageBox from "./component/card/MyMessageBox.svelte";
  import MyNormalHint from "./component/card/MyNormalHint.svelte";
  import MyInputBox from "./component/card/MyInputBox.svelte";
  function ConvertDarkToRGB(dark: boolean, theme: number) {
    const d = DarkAndThemeToMain(dark, theme)
    return {
      dark1: d.substring(d.indexOf('bottom, ') + 8, d.indexOf('), rgb(') + 1),
      dark2: d.substring(d.indexOf('), rgb(') + 3, d.length - 1)
    }
  }
  $: ({dark1, dark2} = ConvertDarkToRGB($dark_mode, $theme_mode))
  function getRotate(dont: number) {
    return {
      rotateX: dont == 1 || dont == 3 ? '180deg' : '0',
      rotateY: dont == 1 || dont == 2 ? '180deg' : '0'
    }
  }
  $: ({rotateX, rotateY} = getRotate($dont_click))
  let backImage = ""
  onMount(async () => {
    let d = await ReadConfig(await GetConfigIniPath(), "Misc", "DarkMode")
    dark_mode.set(d == "1" ? true : d == "0" ? false : window.matchMedia("(prefers-color-scheme: light)").matches)
    let t = Number(await ReadConfig(await GetConfigIniPath(), "Misc", "ThemeMode"))
    theme_mode.set(Number.isNaN(t) || t < 1 || t > 14 ? 1 : t)
    let back = await GetBackgroundImage(-1)
    if (back.length != 0) {
      backImage = `url('data:image/${back[1]};base64,${back[0]}')`
    }else{
      backImage = ``
    }
    document.addEventListener("contextmenu", (e) => {
      e.preventDefault()
    })
  })
</script>

<div id="all" style="--rotate-y: {rotateY}; --rotate-x: {rotateX}">
  <NavBar />
  <!--背景颜色，在下层-->
  <div id="back" style="--dark-1: {dark1}; --dark-2: {dark2}">
    <!--背景图片，在上层-->
    <div id="main" style="--back-image: {backImage}">
      <Body />
    </div>
  </div>
  <MyMessageBox />
  <MyNormalHint />
  <MyInputBox />
</div>

<style>
  #all {
    position: absolute;
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
    overflow: hidden;
    transform: rotateY(var(--rotate-y)) rotateX(var(--rotate-x));
    transition: transform 5s linear;
  }
  #back {
    width: 100%;
    background: linear-gradient(to left bottom, var(--dark-1), var(--dark-2));
    height: calc(100% - 56px);
    transition: all 0.2s;
  }
  #main {
    position: relative;
    width: 100%;
    height: 100%;
    transition: all 0.2s;
    background-size: cover;
    background-repeat: no-repeat;
    background-position: 50% 50%;
    background-image: var(--back-image);
  }
  @keyframes LineAni {
    0% {
      background-position: 0 50%;
    }
    50% {
      background-position: 100% 50%;
    }
    100% {
      background-position: 200% 50%;
    }
  }
</style>
