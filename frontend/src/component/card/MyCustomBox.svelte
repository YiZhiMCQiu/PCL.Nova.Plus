<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    import { quadInOut } from "svelte/easing";
    export let isOpen = false;
    $: bg = $dark_mode ? "#282828cf" : "#f0f8ffcf";
    function back_anim(node: HTMLElement) {
        return {
            duration: 330,
            easing: quadInOut,
            css(t: number) {
                return `opacity: ${t}`;
            },
        };
    }
    function slide_anim(node: HTMLElement) {
        return {
            duration: 330,
            easing: quadInOut,
            css(t: number) {
                const rotate = -10 * (1 - t);
                if (t < 0.8) {
                    const progress = t / 0.8;
                    return `
            opacity: ${progress * 0.8};
            transform: translate(-50%, ${-20 - 30 * progress}%) scale(${0.9 + 0.2 * progress}) rotate(${rotate}deg);
          `;
                } else {
                    const progress = (t - 0.8) / 0.2;
                    return `
            opacity: ${0.8 + 0.2 * progress};
            transform: translate(-50%, -50%) scale(${1.1 - 0.1 * progress}) rotate(${rotate}deg);
          `;
                }
            },
        };
    }
</script>

{#if isOpen}
    <div
        id="back"
        class={isOpen ? "back-class" : "back-class-hide"}
        in:back_anim
        out:back_anim
    ></div>
    <div
        class="content-box-class"
        in:slide_anim
        out:slide_anim
        style="--bg-color: {bg}"
    >
        <slot />
    </div>
{/if}

<style>
    #back {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        z-index: 1000;
        --wails-draggable: drag;
    }

    .back-class {
        background-color: var(--m-back-color);
        backdrop-filter: blur(3px);
        -webkit-backdrop-filter: blur(3px);
        transition: all 0.33s;
    }

    .back-class-hide {
        backdrop-filter: blur(0);
        -webkit-backdrop-filter: blur(0);
        background-color: rgba(0, 0, 0, 0);
        transition: all 0.33s;
    }

    .content-box-class {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        min-width: 200px;
        min-height: 50px;
        height: max-content;
        width: max-content;
        max-width: 800px;
        background-color: var(--bg-color);
        border-radius: 20px;
        border: 1px solid black;
        z-index: 1001;
        transition: box-shadow 0.2s;
    }
    .content-box-class:hover {
        box-shadow: 0 0 10px var(--bg-color);
    }
</style>
