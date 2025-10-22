<script lang="ts">
    import { game_log } from "../../store/mc";
    import { EventsOn } from "../../../wailsjs/runtime/runtime";
    import MySelectCard from "../../component/card/MySelectCard.svelte";
    import MyTextArea from "../../component/input/MyTextArea.svelte";

    export let slide = null;
    export let after_leave = null;

    EventsOn("launch_log", (data: any) => {
        game_log.update((value) => (value += data));
    });
</script>

<div class="component-log_view" in:slide out:slide on:outroend={after_leave}>
    <MySelectCard title="游戏日志" maxHeight={10000}>
        <div class="version-all">
            <MyTextArea
                style_in="width: calc(100% - 16px); height: calc(100vh - 180px)"
                value={$game_log}
                readonly
            />
        </div>
    </MySelectCard>
</div>

<style>
    .component-log_view {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
</style>
