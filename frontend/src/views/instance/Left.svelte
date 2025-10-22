<script lang="ts">
    import MySidebarButton from "../../component/button/MySidebarButton.svelte";
    import { current_instance } from "../../store/changeBody";
    import { onMount } from "svelte";

    export let width = "144px";
    export let slide = null;
    export let after_leave = null;
    onMount(() => {
        current_instance.set("overview");
    });
</script>

<div
    class="component"
    style:width
    in:slide={{ x: Number(width.replace("px", "")) }}
    out:slide={{ x: Number(width.replace("px", "")) }}
    on:outroend={after_leave}
>
    <MySidebarButton
        isChecked={$current_instance === "overview"}
        in_style="margin-top: 8px"
        on:click={() => current_instance.set("overview")}
    >
        <svg
            role="img"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            class="side-icon"
        >
            <path
                d="m12 3l8 4.5v9L12 21l-8-4.5v-9zm0 9l8-4.5M12 12v9m0-9L4 7.5"
            />
        </svg>
        <span>概览</span>
    </MySidebarButton>
    <MySidebarButton
        isChecked={$current_instance === "setting"}
        on:click={() => current_instance.set("setting")}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="side-icon"
        >
            <path
                d="M14 21h-4l-.551-2.48a7 7 0 0 1-1.819-1.05l-2.424.763l-2-3.464l1.872-1.718a7 7 0 0 1 0-2.1L3.206 9.232l2-3.464l2.424.763A7 7 0 0 1 9.45 5.48L10 3h4l.551 2.48a7 7 0 0 1 1.819 1.05l2.424-.763l2 3.464l-1.872 1.718a7 7 0 0 1 0 2.1l1.872 1.718l-2 3.464l-2.424-.763a7 7 0 0 1-1.819 1.052z"
            />
            <circle cx="12" cy="12" r="3" />
        </svg>
        <span>设置</span>
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
