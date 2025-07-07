<script lang="ts">
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import {current_setting} from "../../logic/changeBody";
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
    <div style="display: flex; flex-direction: column">
        <MySidebarButton isChecked={$current_setting === "Launch"} in_style="margin-top: 8px" click={() => {current_setting.set("Launch")}}>
            <svg
                    role="img"
                    viewBox="0 0 29 28"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round" class="side-icon">
                <path d="M12.7526 9.92418C12.2059 10.2861 11.679 10.7057 11.1924 11.1924C10.4754 11.9093 10.4947 11.9482 9.85359 12.682M12.7526 9.92418C16.178 7.65685 20.3848 7.65685 20.3848 7.65685C20.3848 7.65685 20.3848 11.8636 18.1174 15.289M12.7526 9.92418L18.1174 15.289M18.1174 15.289C17.7555 15.8358 17.3359 16.3626 16.8492 16.8492C16.1323 17.5662 16.0934 17.5469 15.3596 18.188M6.11523 17.429C5.74278 17.9526 5.53552 18.2635 5.53552 18.2635L9.77816 22.5061C9.77816 22.5061 10.0891 22.2988 10.6127 21.9264M6.11523 17.429L2.70709 14.0208L8.36394 11.1924L9.85359 12.682M6.11523 17.429C6.83965 16.4105 8.18898 14.5874 9.85359 12.682M10.6127 21.9264L14.0208 25.3345L16.8492 19.6777L15.3596 18.188M10.6127 21.9264C11.6311 21.202 13.4542 19.8526 15.3596 18.188"/>
                <path d="M5.00003 23C5.35031 21.5825 5.99994 21.0001 6.5 21.5C7.00003 22 6.41751 22.6497 5.00003 23Z"/>
            </svg>
            <span>游戏</span>
        </MySidebarButton>
        <MySidebarButton isChecked={$current_setting === "Personalization"} click={() => {current_setting.set("Personalization")}}>
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
        <MySidebarButton isChecked={$current_setting === "Other"} click={() => {current_setting.set("Other")}}>
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
    </div>
</div>
<style>
    .component {
        height: 100%;
    }
    .component span {
        vertical-align: middle;
        margin-left: 9px;
    }
</style>