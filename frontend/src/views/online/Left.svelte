<script lang="ts">
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import {current_online} from "../../store/changeBody";

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
    <MySidebarButton in_style="margin-top: 8px" isChecked={$current_online === "IPv6"} on:click={() => { current_online.set("IPv6") }}>
        <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5" class="side-icon">
            <path d="m9.143 10.691l.207-.207a5.067 5.067 0 1 1 7.166 7.166l-2.866 2.866a5.067 5.067 0 1 1-7.166-7.166l.464-.464"/>
            <path d="m17.052 11.114l.464-.464a5.067 5.067 0 1 0-7.166-7.166L7.484 6.35a5.067 5.067 0 1 0 7.166 7.166l.207-.207"/>
        </svg>
        <span>IPv6 联机</span>
    </MySidebarButton>
</div>
<style>
    .component {
        height: 100%;
        display: flex;
        flex-direction: column;
    }
</style>