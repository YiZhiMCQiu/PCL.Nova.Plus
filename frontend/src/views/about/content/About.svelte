<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import {OpenCustomURL} from "../../../store/functions";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import MyNormalLabel from "../../../component/input/MyNormalLabel.svelte";
    // 导入图片
    import xphost from '../../../assets/images/Heads/xphost.jpg'
    import FireDragon0659 from '../../../assets/images/Heads/firedragon0659.jpg'
    import T3Project from '../../../assets/images/Heads/3T.png'
    import NovaPlus from '../../../assets/images/NovaPlus.png'
    import LtCatt from '../../../assets/images/Heads/ltcatt.jpg'
    import MoYuan from '../../../assets/images/Heads/moyuan.jpg'
    import Hill233 from '../../../assets/images/Heads/hill233.jpg'
    import MyCardButton from "../../../component/button/MyCardButton.svelte";
    import {dark_mode} from "../../../store/changeBody";
    export let slide = null
    export let after_leave = null
    interface People{
        avatar: string,
        title: string,
        desc: string,
        button: [string, () => void, string]
    }
    const dev: People[] = [
        {
            avatar: xphost,
            title: "xphost008",
            desc: "Plain Craft Launcher II: Nova Plus 的现作者之一",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/xphost008")}, `<svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
        {
            avatar: FireDragon0659,
            title: "FireDragon0659",
            desc: "Plain Craft Launcher II: Nova Plus 的现作者之一",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/firedragon0659")}, `<svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
        {
            avatar: T3Project,
            title: "T3 Project",
            desc: "T3 Project 开发团队",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/3TProject")}, `
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
        {
            avatar: NovaPlus,
            title: "Plain Craft Launcher II: Nova Plus",
            desc: "当前版本：Dev 0.0.13",
            button: ["查看源代码", () => {OpenCustomURL("https://github.com/3TProject/PCL.Nova.Plus")}, `
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
    ]
    const credit: People[] = [
        {
            avatar: LtCatt,
            title: "龙腾猫跃",
            desc: "Plain Craft Launcher II 的原作者",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/MoYuan-CN")}, `
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
        {
            avatar: MoYuan,
            title: "墨渊",
            desc: "Plain Craft Launcher II: Nova 的原作者",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/MoYuan-CN")}, `
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        },
        {
            avatar: Hill233,
            title: "Hill233",
            desc: "Plain Craft Launcher II: Neo 的原作者",
            button: ["Github 主页", () => {OpenCustomURL("https://github.com/Hill23333")}, `
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? '#ffffff' : '#000000'}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                            </svg>`]
        }
    ]
</script>
<div
        class="component-about"
        in:slide
        out:slide
        on:outroend={after_leave}
>
    <MySelectCard title="开发人员">
        <div class="proc">
            {#each dev as d}
                <MyCardButton
                    image={d.avatar}
                    title={d.title}
                    desc={d.desc}
                    image_style="border-radius: 50%;"
                    button_style="width: 50px; height: 50px; top: 0;"
                    buttons={[
                        {
                            title: d.button[0],
                            click: d.button[1],
                            icon: d.button[2],
                        }
                    ]}
                >
                </MyCardButton>
            {/each}
        </div>
    </MySelectCard>
    <MySelectCard title="鸣谢人员">
        <div class="proc">
            {#each credit as d}
                <MyCardButton
                        image={d.avatar}
                        title={d.title}
                        desc={d.desc}
                        image_style="border-radius: 50%;"
                        button_style="width: 50px; height: 50px; top: 0;"
                        buttons={[
                        {
                            title: d.button[0],
                            click: d.button[1],
                            icon: d.button[2],
                        }
                    ]}
                >
                </MyCardButton>
            {/each}
        </div>
    </MySelectCard>
    <MySelectCard title="协议" isExpand={true}>
        <div class="proc">
            <div style="margin: 0 20px">
                <MyNormalLabel>PCL.Nova.Plus 开源项目及其官方衍生程序均遵循 GNU General Public License 3.0 (GPL v3) 协议开源。本项目主导组织为 T3 Project 以及 PCL Community，任何人都可以提 Pull Request！</MyNormalLabel><br><br>
                <MyNormalButton style_in="width: 150px; height: 40px">查看源代码</MyNormalButton>
                <MyNormalButton style_in="width: 150px; height: 40px; margin-left: 20px">查看许可证</MyNormalButton>
            </div>
        </div>
    </MySelectCard>
    <MySelectCard title="许可与版权声明" isExpand={true}>
        <div class="proc">
            <div style="margin: 0 20px; width: calc(100% - 40px)">
                <MyNormalLabel>PCL.Nova.Plus</MyNormalLabel><br>
                <span class="gray">GNU-GPL v3</span>&nbsp;<span class="code">https://github.com/3TProject/PCL.Nova.Plus</span><br>
                <MyNormalLabel>Wails</MyNormalLabel><br>
                <span class="gray">MIT</span>&nbsp;<span class="code">https://github.com/wailsapp/wails</span><br>
                <MyNormalLabel>Golang</MyNormalLabel><br>
                <span class="gray">BSD 3</span>&nbsp;<span class="code">https://github.com/golang/go</span><br>
                <MyNormalLabel>Svelte</MyNormalLabel><br>
                <span class="gray">MIT</span>&nbsp;<span class="code">https://github.com/sveltejs/svelte</span><br>
                <MyNormalLabel>TypeScript</MyNormalLabel><br>
                <span class="gray">Apache License 2.0</span>&nbsp;<span class="code">https://github.com/microsoft/TypeScript</span><br>
            </div>
        </div>
    </MySelectCard>
</div>

<style>
    .component-about {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .all {
        height: 60px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        width: 100%;
    }
    .all img {
        width: 50px;
        height: 50px;
        border-radius: 50%;
        margin: 0 10px 5px 20px;
    }
    .desc p{
        padding: 0;
        margin: 0;
    }
    .info {
        display: flex;
        flex-direction: row;
    }
    .buttons {
        display: flex;
        flex-direction: row;
    }
    .desc {
        display: flex;
        flex-direction: column;
    }
    .desc p:last-child {
        color: gray;
    }
    .gray {
        color: gray;
        font-size: 14px;
    }
    .code {
        box-shadow: 0 0 3px black;
        background-color: lightgray;
        border-radius: 4px;
        color: black;
        padding: 0 5px;
        user-select: all;
        font-size: 14px;
    }
</style>