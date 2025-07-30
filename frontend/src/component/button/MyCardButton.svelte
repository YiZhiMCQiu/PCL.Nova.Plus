<script lang="ts">
    import MyNormalLabel from "../input/MyNormalLabel.svelte";
    // 图片
    export let image = null
    // 悬停提示
    export let hint = ""
    // 大标题
    export let title = ""
    // 描述
    export let desc = ""
    // 点击事件
    export let click: (() => void | null) = null
    // 右侧icon按钮组【内部必须有（title、icon、click）三个字段，其中icon是一个必须有的svg标签的字符串】
    export let buttons = []
    // 图片样式【只能调整圆角，宽高度最大50px，如果超过请自行在源码里做适配】
    export let image_style = ""
    // 按钮样式（调整按钮的高度、宽度等。。必要的时候需要调整top参数，图标样式可以直接内联写~）
    export let button_style = ""
</script>
<div class={click === null ? 'comp' : 'comp-click'} on:click={() => {
                if(click !== null) {
                    click()
                }
            }} on:keydown|preventDefault title={hint}>
    <div class="info">
        <img src={image} alt="图片" style={image_style}>
        <div class="desc">
            <p><MyNormalLabel style_in="font-size: 17px; cursor: pointer">{title}</MyNormalLabel></p>
            <p style="font-size: 14px">{desc}</p>
        </div>
    </div>
    <div class="buttons">
        {#each buttons as button}
            <button title={button.title} on:click|stopPropagation={() => {
                if(button.click !== null) {
                    button.click()
                }
            }} style="{button_style}">
                <!-- 请记住，这个icon只能填入svg，并且千万不要加width和height俩参数！ -->
                {@html button.icon}
            </button>
        {/each}
    </div>
</div>
<style>
    .comp,
    .comp-click{
        width: 100%;
        height: 50px;
        transition: all 0.2s;
        border-radius: 10px;
        display: flex;
        flex-direction: row;
        align-items: center;
    }
    .comp-click {
        cursor: pointer;
    }
    .comp-click:hover {
        background-color: rgba(100, 100, 156, 0.2);
    }
    .comp-click:active {
        transform: scale(98%);
        background-color: rgba(100, 100, 156, 0.3);
    }
    .info {
        display: flex;
        align-items: center;
        height: 50px;
        width: max-content;
    }
    .info img {
        width: 40px;
        height: 40px;
        margin-left: 10px;
    }
    .desc {
        display: flex;
        flex-direction: column;
        margin-left: 8px;
        pointer-events: none;
    }
    .desc p {
        padding: 0;
        margin: 0;
    }
    .desc p:last-child {
        color: gray;
    }
    .buttons {
        width: 120px;
        height: 50px;
        margin-right: 10px;
        flex: 1;
    }
    .buttons button {
        position: relative;
        top: 8px;
        width: 34px;
        height: 34px;
        background-color: transparent;
        border: 0;
        transition: all 0.2s;
        cursor: pointer;
        float: right;
    }
    .comp-click:hover .buttons button :global(svg) {
        stroke: #00aaff;
    }
    .comp-click:hover .buttons button:hover :global(svg) {
        stroke: #0080ff;
    }
    .buttons button :global(svg),
    .buttons button :global(img) {
        width: 22px;
        height: 22px;
        vertical-align: middle;
        stroke: transparent;
        transition: all 0.2s;
        position: relative;
        left: -1px;
        top: 1px;
    }
</style>