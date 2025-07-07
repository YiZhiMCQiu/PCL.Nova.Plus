<script lang="ts">
    import MyNormalButton from "../../component/button/MyNormalButton.svelte";
    import {current_account_page} from "../../logic/changeBody";
    import {quadInOut} from "svelte/easing";
    import {onDestroy} from "svelte";
    import AccountSelect from "./account/AccountSelect.svelte";
    import AddAccount from "./account/AddAccount.svelte";

    export let width = "144px"
    export let slide = null
    export let after_leave = null
    let isTransitioning = true
    let f = false

    function control_leave() {
        isTransitioning = true
    }
    const unsubscribe_current_view = current_account_page.subscribe((value) => {
        if (!f) {
            f = true
            isTransitioning = true;
        }else{
            isTransitioning = !isTransitioning
        }
    })
    function account_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `opacity: ${t};`;
            }
        };
    }
    onDestroy(unsubscribe_current_view)
</script>
<div
        class="component"
        style:width={width}
        in:slide={{ x: Number(width.replace("px", "")) }}
        out:slide={{ x: Number(width.replace("px", "")) }}
        on:outroend={after_leave}
>
    <div style="display: flex; flex-direction: column; height: 100%">
        <div id="middle">
            {#if $current_account_page && isTransitioning}
                <AccountSelect opacity={account_opacity} after_leave={control_leave} />
            {:else if !$current_account_page && isTransitioning}
                <AddAccount opacity={account_opacity} after_leave={control_leave} />
            {/if}
        </div>
        <div id="bottom">
            <MyNormalButton style_in="height: 75px; width: calc(100% - 52px); border: 1px solid #216fbd; margin-top: 6px" isDisabled={!$current_account_page}>
                <span id="launch-title">启动游戏</span><br>
                <span id="launch-version">测试客户端</span>
            </MyNormalButton>
            <div id="setting">
                <MyNormalButton style_in="width: calc(50% - 4px); height: 40px;" isDisabled={!$current_account_page}>
                    选择核心
                </MyNormalButton>
                <MyNormalButton style_in="width: calc(50% - 4px); height: 40px;" isDisabled={!$current_account_page}>
                    核心设置
                </MyNormalButton>
            </div>
        </div>
    </div>
</div>
<style>
    .component {
        height: 100%;
    }
    #middle {
        width: 100%;
        flex: 1;
    }
    #bottom {
        width: 100%;
        height: 156px;
        flex-shrink: 0;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    #launch-title {
        background-image: linear-gradient(to right, rgb(63, 207, 255), rgb(96, 96, 255));
        color: transparent;
        background-clip: text;
        font-size: 20px;
    }
    #launch-version {
        font-size: 12px;
    }
    #setting {
        height: 54px;
        width: calc(100% - 52px);
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
    }
</style>