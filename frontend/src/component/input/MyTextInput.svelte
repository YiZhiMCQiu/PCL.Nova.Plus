<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    import { createEventDispatcher } from "svelte";

    export let placeholder = "";
    export let style_in = "";
    const dispatch = createEventDispatcher();
    export let title = "";
    export let value = "";
    export let password = false;

    function onInput(e: Event) {
        dispatch("input", {
            value: (e.target as HTMLInputElement).value,
        });
    }

    function onBlur(e: Event) {
        dispatch("blur", {
            value: (e.target as HTMLInputElement).value,
        });
    }

    function updateInputType(node: HTMLInputElement, isPassword: boolean) {
        node.type = isPassword ? "password" : "text";
        return {
            update(isPassword: boolean) {
                node.type = isPassword ? "password" : "text";
            },
        };
    }
    $: ({ light, dark } = $dark_mode
        ? { light: "#e6e6e6cf", dark: "#1a1a1acf" }
        : { light: "#1a1a1acf", dark: "#e6e6e6cf" });
</script>

<input
    use:updateInputType={password}
    bind:value
    {title}
    {placeholder}
    class="text-input font-pcl"
    style="{style_in}; --light-color: {light}; --dark-color: {dark};"
    on:input={onInput}
    on:blur={onBlur}
/>

<style>
    .text-input {
        border: 1px solid var(--light-color);
        border-radius: 5px;
        background-color: var(--dark-color);
        color: var(--light-color);
        padding: 0 10px;
        transition: all 0.2s;
        min-width: 100px;
    }
</style>
