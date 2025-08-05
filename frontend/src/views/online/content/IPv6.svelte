<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import {onDestroy, onMount} from "svelte";
    import {select_ipv6} from "../../../store/changeBody";
    import MyLoadingPickaxe from "../../../component/card/MyLoadingPickaxe.svelte";
    import {slide_up, slide_opacity, copyToClipBoard} from "../../../store/functions";
    import {GetAllIPv6} from "../../../../wailsjs/go/launcher/MainMethod";
    import MyCardButton from "../../../component/button/MyCardButton.svelte";
    import IPv6Image from '../../../assets/images/Icons/IPv6.png'
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import {HNT_PASS, showHint} from "../../../store/messagebox";
    let loading_text = ""
    let loading_state = false
    export let slide = null
    export let after_leave = null
    let isTransitioning = true
    function control_leave() {
        isTransitioning = true
    }
    let f = false
    const unsubscribe_select_ipv6 = select_ipv6.subscribe((value) => {
        if (!f) {
            f = true
            isTransitioning = true
        } else {
            isTransitioning = !isTransitioning
        }
    })
    async function reloadIPv6(){
        if($select_ipv6.length <= 0) {
            loading_text = "正在加载 IPv6"
            loading_state = false
            let meta = await GetAllIPv6()
            if(meta.length <= 0) {
                loading_text = "无法加载 IPv6，您可能目前暂未拥有 IPv6！\n如果你确保你本机拥有 IPv6，清尝试使用 sudo 或者管理员权限再次打开次界面！"
                loading_state = true
                return
            }
            select_ipv6.update((value) => {
                Object.assign(value, meta)
                return value
            })
        }
    }
    onDestroy(unsubscribe_select_ipv6)
    onMount(reloadIPv6)
    async function selectOneIPv6(index: number) {
        copyToClipBoard(`[${$select_ipv6[index].ip}]`)
        showHint("复制成功😀", HNT_PASS)
    }
</script>
<div
        class="component-ipv6"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    {#if $select_ipv6.length > 0 && isTransitioning}
        <div in:slide_up out:slide_up on:outroend={control_leave}>
            <MySelectCard title="IPv6 检测">
                <div class="version-all">
                    <MyNormalButton style_in="width: 100px; height: 30px; margin-bottom: 10px" click={() => {
                        select_ipv6.set([])
                        reloadIPv6()
                    }}>
                        重新检测
                    </MyNormalButton>
                    {#each $select_ipv6 as ipv6, i}
                        <MyCardButton
                                image={IPv6Image}
                                title={ipv6.ip}
                                desc={ipv6.success ? '测试通过' : '测试失败'}
                                hint={ipv6.success ? '该 IPv6 测试已通过，请单机该按钮，将 ip 复制到剪切板，随后配上你的 MC 端口，发送给你的好友！' : '啊哦，你的 IPv6 测试失败了！应该是没有 Ping 通的原因，你可以尝试 bing 搜索一下 IPv6 测试，随后查看一下你的 IPv6 是否正常~'}
                                click={() => selectOneIPv6(i)}
                        />
                    {/each}
                </div>
            </MySelectCard>
        </div>
    {:else if isTransitioning}
        <div style="width: 100%; height: 100%; display: flex; align-items: center; justify-content: center;" in:slide_opacity out:slide_opacity
             on:outroend={control_leave}>
            <MyLoadingPickaxe loading_text={loading_text} state={loading_state}/>
        </div>
    {/if}
</div>
<style>
    .component-ipv6 {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
</style>