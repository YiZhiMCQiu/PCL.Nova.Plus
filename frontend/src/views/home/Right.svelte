<script lang="ts">
    import {onDestroy, onMount} from "svelte";
    import HomePage from "./content/HomePage.svelte";
    import {parseXmlToJson, getHomePage} from "../../store/functions";
    import {messagebox, MSG_ERROR} from "../../store/messagebox";
    import {
        homepage_cache,
    } from "../../store/changeBody";
    // 检测 主页缓存 的变化。
    const unsubscribe_homepage_cache = homepage_cache.subscribe(async (value) => {
        if(value == "") {
            let hp = await getHomePage()
            if (hp != "none" && await recursion(hp)) {
                homepage_cache.set(hp)
            }
        }
    })
    onDestroy(unsubscribe_homepage_cache)
    export let slide = null
    export let after_leave = null
    let structures = []
    async function recursion(xml: string): Promise<boolean> {
        try{
            structures = parseXmlToJson(`<?xml version="1.0" encoding="UTF-8"?><NovaHomePage>${xml.replaceAll('<?xml version="1.0" encoding="UTF-8"?>', '')}</NovaHomePage>`)[0].children
            if(!structures) {
                structures = []
            }
            return true
        }catch(e: any) {
            await messagebox("自定义主页解析错误！", "自定义主页解析错误，信息：" + e.message, MSG_ERROR)
            return false
        }
    }
    // 强制刷新一次监测，同时强制刷新一次主页（
    onMount(async () => {
        if($homepage_cache == "") {
            homepage_cache.set("1")
            homepage_cache.set("")
        }else{
            await recursion($homepage_cache)
        }
    })
</script>
<div
        id="home"
        class="component-right"
        in:slide
        out:slide
        on:outroend={after_leave}>
    {#each structures as structure}
        <HomePage control={structure}></HomePage>
    {/each}
</div>
<style>
    .component-right {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>