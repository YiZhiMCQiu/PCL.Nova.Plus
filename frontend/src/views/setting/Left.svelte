<script lang="ts">
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import {current_setting} from "../../store/changeBody";

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
    <MySidebarButton isChecked={$current_setting === "Launch"} in_style="margin-top: 8px"
                     on:click={() => {current_setting.set("Launch")}}>
        <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                fill="none" class="side-icon">
            <path d="M15.59 14.37q.159.666.16 1.38a6 6 0 0 1-6 6v-4.8m5.84-2.58a14.98 14.98 0 0 0 6.16-12.12A14.98 14.98 0 0 0 9.631 8.41m5.96 5.96a14.9 14.9 0 0 1-5.841 2.58m-.119-8.54a6 6 0 0 0-7.381 5.84h4.8m2.581-5.84a14.9 14.9 0 0 0-2.58 5.84m2.699 2.7q-.155.032-.311.06a15 15 0 0 1-2.448-2.448l.06-.312m-2.24 2.39a4.49 4.49 0 0 0-1.757 4.306q.341.054.696.054a4.5 4.5 0 0 0 3.61-1.812M16.5 9a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/>
        </svg>
        <span>启动</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_setting === "Personalization"}
                     on:click={() => {current_setting.set("Personalization")}}>
        <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" class="side-icon">
            <path d="M19 14V22H2.99997V4H13"/>
            <path d="M17.4608 4.03921C18.2418 3.25817 19.5082 3.25816 20.2892 4.03921L20.9608 4.71079C21.7418 5.49184 21.7418 6.75817 20.9608 7.53921L11.5858 16.9142C11.2107 17.2893 10.702 17.5 10.1716 17.5L7.5 17.5L7.5 14.8284C7.5 14.298 7.71071 13.7893 8.08579 13.4142L17.4608 4.03921Z"/>
            <path d="M16.25 5.25L19.75 8.75"/>
        </svg>
        <span>个性化</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_setting === "Other"} on:click={() => {current_setting.set("Other")}}>
        <svg
                role="img"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none" class="side-icon">
            <path d="M3 5H21V21H3V5Z"/>
            <path d="M21 9H3"/>
            <path d="M7 5V3"/>
            <path d="M17 5V3"/>
            <path d="M15 15H9"/>
            <path d="M12 12V18"/>
        </svg>
        <span>启动器</span>
    </MySidebarButton>
    <MySidebarButton isChecked={$current_setting === "HomePage"} on:click={() => {current_setting.set("HomePage")}}>
        <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 14 14"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round" class="side-icon">
            <path
                    d="M13.5 1H.5m13 4H6m7.5 4H6m7.5 4H.5M2 5l2 2l-2 2"/>
        </svg>
        <span>主页</span>
    </MySidebarButton>
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
</style>