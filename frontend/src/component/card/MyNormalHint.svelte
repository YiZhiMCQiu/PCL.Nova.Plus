<script lang="ts">
    import { h_content, h_level } from "../../store/messagebox";
    import { quadInOut } from "svelte/easing";
    $: backColor =
        $h_level == 0
            ? "rgba(0, 128, 255, 207)"
            : $h_level == 1
              ? "#c7ad2acf"
              : $h_level == 2
                ? "#ff4c4ccf"
                : "#10e830cf";
    $: content = $h_content;
    function slide(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `
                    transform: translateX(-${100 - t * 100}%);
                `;
            },
        };
    }
</script>

{#if content !== ""}
    <div in:slide out:slide id="hint" style="--back-color: {backColor}">
        {content}
    </div>
{/if}

<style>
    #hint {
        position: fixed;
        bottom: 20px;
        left: 0;
        height: 20px;
        font-size: 13px;
        border-radius: 0 4px 4px 0;
        display: flex;
        padding: 2px 5px 2px 10px;
        align-items: center;
        transition: all 0.2s;
        background-color: var(--back-color);
        color: white;
    }
</style>
