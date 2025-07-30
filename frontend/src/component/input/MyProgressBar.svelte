<script lang="ts">
    export let max = 100
    export let min = 0
    export let value = 0
    export let onDragging: ((value: number) => void) | null = null
    // 获取两个 DOM 结构，一个是 整个滑动条 的 DOM，一个是 滑块 的 DOM
    let progressDOM: (HTMLDivElement | null) = null
    let thumbDOM: (HTMLDivElement | null) = null
    // 设置一些值
    // 当前滑块所处的 left 位置
    let calThumb = 0
    // 当前滑动条的百分比
    let calPercent = 0
    // 当前是否正在滑动
    let isDragging = false
    // 当前滑块左侧相对鼠标点击的位置
    let oX = 0
    // 更新 滑动条信息
    function updateBar() {
        calThumb = ((value - min) / (max - min)) * progressDOM!.offsetWidth - 5;
        calPercent = value / max * 100
    }
    // 当滑动条任意位置点击时触发的事件
    // ⚠️这部分和以下部分用了些许AI，请注意甄别~
    function ProgressClick(e: MouseEvent) {
        if(e.target == thumbDOM) return
        const tRect = progressDOM.getBoundingClientRect()
        const aWidth = tRect.width - thumbDOM.offsetWidth
        let mX = e.clientX - tRect.left - thumbDOM.offsetWidth / 2
        mX = Math.max(0, Math.min(mX, aWidth))
        const step = aWidth / (max - min)
        value = min + Math.round(mX / step)
        if(onDragging != null) {
            onDragging(value)
        }
        updateBar()
    }
    // 当滑动条开始拖拽时的事件
    function onDrag(e: MouseEvent) {
        if(!isDragging) return
        const tRect = progressDOM.getBoundingClientRect()
        const aWidth = tRect.width - thumbDOM.offsetWidth
        let mX = e.clientX - tRect.left - oX
        mX = Math.max(0, Math.min(mX, aWidth))
        const step = aWidth / (max - min)
        value = min + Math.round(mX / step)
        if(onDragging != null) {
            onDragging(value)
        }
        updateBar()
    }
    // 当鼠标抬起，不滑动了的事件
    function onStopDrag(e: MouseEvent) {
        isDragging = false
        document.body.style.cursor = 'default'
        document.removeEventListener('mousemove', onDrag)
        document.removeEventListener('mouseup', onStopDrag)
    }
    // 当鼠标按住滑块开始拖动了的事件
    function thumbSlide(e: MouseEvent) {
        isDragging = true
        const tl = thumbDOM!.getBoundingClientRect()
        oX = e.clientX - tl.left
        document.body.style.cursor = 'grabbing'
        document.addEventListener('mousemove', onDrag)
        document.addEventListener('mouseup', onStopDrag)
    }
</script>
<!-- 用了一点非常巧妙的方式实现了响应式更新 value（ -->
<div id="progress-bar" style="--hov-percent-left: {value / max * 100}%; --hov-percent-right: {1 - value / max * 100}%; --thumb-left: {Math.floor(value / max * 100)}%;" on:click={ProgressClick} on:keydown|preventDefault>
    <div id="progress" bind:this={progressDOM} on:mousedown|stopPropagation={thumbSlide}>
        <div id="bar" bind:this={thumbDOM}></div>
    </div>
</div>
<style>
    #progress-bar {
        width: 100%;
        display: flex;
        align-items: center;
        height: 30px;
    }
    #progress-bar:hover{
        cursor: pointer;
    }
    #progress-bar:hover #progress {
        background: linear-gradient(to right, rgb(12, 92, 180) var(--hov-percent-left), rgb(185, 201, 218) var(--hov-percent-right));
    }
    #progress-bar:hover #bar {
        cursor: grab;
        background-color: rgb(12, 92, 180);

    }
    #progress-bar:active #bar {
        cursor: grabbing;
        transform: scale(1.2);
    }
    #progress {
        position: relative;
        width: 100%;
        height: 2px;
        border-radius: 5px;
        background: linear-gradient(to right, rgb(18, 122, 225) var(--hov-percent-left), rgb(215, 231, 248) var(--hov-percent-right));
        transition: all 0.2s;
    }
    #bar {
        width: 10px;
        height: 10px;
        border-radius: 50%;
        background-color: rgb(18, 122, 225);
        position: absolute;
        top: -4px;
        left: var(--thumb-left);
        transition: transform, background-color 0.2s;
    }
</style>