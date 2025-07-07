<script lang="ts">
  import NavBar from './views/NavBar.svelte'
  import Body from './views/Body.svelte'
  import {dark_mode, dont_click, theme_mode} from "./logic/changeBody";
  import {onMount} from "svelte";
  import {GetConfigIniPath, ReadConfig} from "../wailsjs/go/launcher/ReaderWriter";
  import {GetBackgroundImage} from "../wailsjs/go/launcher/MainMethod";
  import {DarkAndThemeToMain} from "./logic/functions";
  import MyMessageBox from "./component/card/MyMessageBox.svelte";
  function ConvertDarkToRGB(dark: boolean, theme: number) {
    const d = DarkAndThemeToMain(dark, theme)
    return {
      dark1: d.substring(d.indexOf('bottom, ') + 8, d.indexOf('), rgb(') + 1),
      dark2: d.substring(d.indexOf('), rgb(') + 3, d.length - 1)
    }
  }
  $: ({dark1, dark2} = ConvertDarkToRGB($dark_mode, $theme_mode))
  $: rotate = $dont_click == 1 ? '180deg' : '0'
  let backImage = ""
  onMount(async () => {
    dark_mode.set(await ReadConfig(await GetConfigIniPath(), "Misc", "DarkMode") === "1")
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

<div id="all" style="--rotate: {rotate}">
  <NavBar />
  <!--背景颜色，在下层-->
  <main id="back" style="--dark-1: {dark1}; --dark-2: {dark2}">
    <!--背景图片，在上层-->
    <main id="main" style="--back-image: {backImage}">
      <Body />
    </main>
  </main>
  <MyMessageBox />
</div>

<style>
  #all {
    position: absolute;
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
    overflow: hidden;
    transform: rotate(var(--rotate));
    transition: transform 5s linear;
  }
  #back {
    width: 100%;
    background: linear-gradient(to left bottom, var(--dark-1), var(--dark-2));
    height: calc(100% - 56px);
    transition: all 10s;
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
