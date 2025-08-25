<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import MyTextInput from "../../../component/input/MyTextInput.svelte";
    // import MyComboBox from "../../../component/card/MyComboBox.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {dark_mode, ModrinthSearchURL} from "../../../store/changeBody";
    import {HttpGet} from "../../../../wailsjs/go/launcher/Network";
    import {onDestroy, onMount} from "svelte";
    import {search_mod_cache} from "../../../store/mc";
    import {slide_up, slide_opacity} from "../../../store/functions";
    import MyLoadingPickaxe from "../../../component/card/MyLoadingPickaxe.svelte";
    import MySearchButton from "../../../component/button/MySearchButton.svelte";

    export let slide = null
    export let after_leave = null
    let maxHeight = 0
    let searchText = ""
    let isTransitioning = true
    let loading_text = ""
    let loading_state = false
    function control_leave() {
        isTransitioning = true
    }
    function onSearchBoxInput(e: CustomEvent) {
        searchText = e.detail.value
    }
    let f = true
    const unsubscribe_search_mod_cache = search_mod_cache.subscribe(value => {
        if(f) {
            f = false
            isTransitioning = true
        } else {
            isTransitioning = !isTransitioning
        }
    })
    function onSearchButtonClick() {
        search_mod_cache.set([])
        onSearch()
    }
    async function onSearch() {
        loading_text = "正在搜索 Mod 中~"
        loading_state = false
        let cg = await HttpGet(ModrinthSearchURL + `&query=${encodeURIComponent(searchText)}`, "")
        if(cg != "") {
            let json = JSON.parse(cg).hits
            if(json.length <= 0) {
                loading_text = "未能找到你所需要的 Mod！"
                loading_state = true
                return
            }
            for(let i = 0; i < json.length; i++) {
                search_mod_cache.update(value => {
                    value.push({
                        image_url: json[i].icon_url,
                        mod_name: json[i].title,
                        mod_id: json[i].slug,
                        libraries: json[i].categories,
                        description: json[i].description,
                        mod_type: json[i].categories[0] || "",
                        download_count: json[i].downloads > 10000 ? Math.floor(json[i].downloads / 10000) + "万" : json[i].downloads,
                        update_date: json[i].date_modified,
                        mod_source: "Modrinth"
                    })
                    return value
                })
            }
        }else{
            loading_text = "搜索失败！请检查网络配置！"
            loading_state = true
        }
        maxHeight = $search_mod_cache.length * 80 + 20
    }
    onMount(() => {
        if($search_mod_cache.length <= 0) {
            onSearch()
        }
    })
    function clickAnyMod(index: number) {

    }
    onDestroy(unsubscribe_search_mod_cache)
</script>
<div
        class="component-extension_mod"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <MySelectCard title="">
        <div class="search-box">
            <MyTextInput style_in="border: none; outline: none; flex: 1; height: 50px; font-size: 20px;" on:blur={onSearchBoxInput} placeholder="搜索 Mod~"/>
            <MyNormalButton style_in="width: 100px; height: 50px; border: none; cursor: pointer" on:click={onSearchButtonClick}>
                <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        style="width: 30px; height: 30px; vertical-align: middle;" style:stroke={$dark_mode ? '#f6f6f6cf' : '#1a1a1acf'}>
                    <path d="M15 9.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z M20 20l-6-6"/>
                </svg>
            </MyNormalButton>
        </div>
    </MySelectCard>
    {#if isTransitioning && $search_mod_cache.length > 0}
        <div in:slide_up out:slide_up on:outroend={control_leave}>
            <MySelectCard title="" maxHeight={maxHeight}>
                <div class="version-all">
                    {#each $search_mod_cache as mod, i}
                        <MySearchButton
                            image_url={mod.image_url}
                            mod_name={mod.mod_name}
                            mod_id={mod.mod_id}
                            libraries={mod.libraries}
                            description={mod.description}
                            mod_type={mod.mod_type}
                            download_count={mod.download_count}
                            update_date={mod.update_date}
                            mod_source={mod.mod_source}
                            on:click={() => clickAnyMod(i)}
                        />
                    {/each}
                </div>
            </MySelectCard>
        </div>
    {:else if isTransitioning}
        <div in:slide_opacity out:slide_opacity on:outroend={control_leave} class="search-loading-content">
            <MyLoadingPickaxe loading_text={loading_text} state={loading_state}/>
        </div>
    {/if}
</div>
<style>
    .component-extension_mod {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .search-box {
        width: 100%;
        height: 50px;
        display: flex;
        align-items: center;
    }
    .search-loading-content {
        width: 100%;
        height: calc(100% - 65px);
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>