<script lang="ts">
    import {dark_mode} from "../../store/changeBody";
    $: light = $dark_mode ? "#e6e6e6cf" : "#1a1a1acf"
    export let style_in = ""
    export let title = ""
</script>
<span class="font-pcl" title={title} style="font-size: 16px; color: {light}; stroke: {light}; {style_in};" >
    <slot />
</span>