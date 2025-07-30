<script lang="ts">
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import {current_about} from "../../store/changeBody";
    export let width = "144px"
    export let slide = null
    export let after_leave = null
</script>
<div
    class="component"
    style:width={width}
    in:slide={{ x: Number(width.replace("px", "")) }}
    out:slide={{ x: Number(width.replace("px", "")) }}
    on:outroend={after_leave}
>
    <MySidebarButton isChecked={$current_about === "Help"} in_style="margin-top: 8px" click={() => {current_about.set("Help")}}>
        <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                fill="none" class="side-icon">
            <g>
                <circle cx="12" cy="12" r="10"/>
                <path d="M9 10c0 -1.66 1.34 -3 3 -3c1.66 0 3 1.34 3 3c0 0.98 -0.47 1.85 -1.2 2.4c-0.73 0.55 -1.3 0.6 -1.8 1.6" />
                <path d="M12 17v0.01" />
            </g>
        </svg>
        <span>帮助</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_about === "About"} click={() => {current_about.set("About")}}>
        <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" class="side-icon">
            <path stroke-width="2" d="M12,11 L12,17"/>
            <line stroke-width="2" x1="12" y1="8" x2="12" y2="8"/>
            <circle cx="12" cy="12" r="10"/>
        </svg>
        <span>关于与鸣谢</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_about === "Treasure-Box"} click={() => {current_about.set("Treasure-Box")}}>
        <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" class="side-icon">
            <title id="coloursIconTitle">Colours</title>
            <circle cx="12" cy="9" r="5"/>
            <circle cx="9" cy="14" r="5"/>
            <circle cx="15" cy="14" r="5"/>
        </svg>
        <span>百宝箱</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_about === "Feedback"} click={() => {current_about.set("Feedback")}}>
        <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" class="side-icon">
            <rect width="20" height="14" x="2" y="5"/>
            <path stroke-linecap="round" d="M2 5l10 9 10-9"/>
            <path stroke-linecap="round" d="M2 19l6.825-7.8"/>
            <path stroke-linecap="round" d="M22 19l-6.844-7.822"/>
        </svg>
        <span>反馈</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_about === "Discussion"} click={() => {current_about.set("Discussion")}}>
        <svg
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 48 48"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="4"
                fill="none" class="side-icon">
            <g>
                <path d="M17 18H4v24h13z"/>
                <path d="M30 6H17v36h13z"/>
                <path d="M43 26H30v16h13z"/>
            </g>
        </svg>
        <span>新功能投票</span>
    </MySidebarButton>
    <details class="grid"><summary>小游戏</summary></details>
    <div>
        <MySidebarButton isChecked={$current_about === "Minesweeper"} click={() => {current_about.set("Minesweeper")}}>
            <svg
                    role="img"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-width="1"
                    stroke-linecap="square"
                    stroke-linejoin="miter"
                    fill="none" class="side-icon">
                <path d="M5 20v-5c1.333-1.333 3-2 5-2 3 0 4 2 6 2 1 0 2-.333 3-1V5c-1 .667-2 1-3 1-2 0-3-2-6-2-2 0-3.667.667-5 2v9"/>
            </svg>
            <span>扫雷</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_about === "P2048"} click={() => {current_about.set("P2048")}}>
            <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    fill="none" class="side-icon">
                <path d="M8 8a4 4 0 1 1 8 0c0 1.098-.564 2.025-1.159 2.815L8 20h8"/>
            </svg>
            <span>2048</span>
        </MySidebarButton>
    </div>
</div>
<style>
    .component {
        height: 100%;
        display: flex;
        flex-direction: column;
    }
    .component span {
        vertical-align: middle;
        margin-left: 9px;
    }
    .grid {
        margin-top: 25px;
    }
    .grid summary {
        font-size: 12px;
        color: gray;
        margin-bottom: 7px;
        margin-left: 20px;
    }
    .grid + div {
        max-height: 0;
        transition: max-height 0.5s;
        margin: 0 0 1rem;
        overflow: hidden;
    }
    .grid[open] + div {
        max-height: 80px;
    }
</style>