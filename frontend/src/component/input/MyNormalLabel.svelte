<script lang="ts">
    import {dark_mode} from "../../logic/changeBody";
    $: light = $dark_mode ? "#e6e6e6cf" : "#1a1a1acf"
    export let style_in = ""
</script>
<label class="font-pcl" style="font-size: 16px; color: {light}; stroke: {light}; {style_in};" >
    <slot />
</label>