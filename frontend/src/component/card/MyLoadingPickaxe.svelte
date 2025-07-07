<script lang="ts">
    import {dark_mode} from "../../logic/changeBody";
    export let loading_text = "正在加载……"
    // false: 加载中、true: 失败
    export let state = false
    export let style_in = ""
    $: dark = $dark_mode ? "#00080fcf" : "#f0f8ffcf"
</script>
<div class="ala" style="{style_in}; --dark-color: {dark}; --font-color: {state ? '#ff1f1fcf' : '#1370F3cf'}; --is-infinite: {state ? '1' : 'infinite'}; --icon-hidden: {state ? '1' : '0'}">
    <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 80 80"
            fill="none"
            class="loading-svg">
        <!--镐子底下的线-->
        <path
                stroke-width="2"
                stroke-linecap="round"
                d="M5 66L29 66" />
        <!--镐子本体-->
        <path
                d="M53.7693 59.6884L53.7763 27.5221C59.2904 27.6694 62.7075 30.1275 67.6568 31.1071C73.6227 32.4731 74.8377 33.7263 71.1429 29.9828C66.9914 25.9769 61.6417 23.1999 55.6855 22.2612C55.1971 20.2186 53.3582 18.7011 51.1671 18.7045C48.976 18.7078 47.138 20.2185 46.6487 22.2612C40.6892 23.2052 35.3415 25.9781 31.1913 29.9828C26.8344 34.236 28.2414 32.9395 34.6703 31.1142C39.6201 30.1326 43.0467 27.6687 48.5579 27.5221L48.5508 59.6884C48.5508 59.7233 48.5795 59.752 48.6145 59.752L53.7056 59.752C53.7434 59.7491 53.7663 59.7262 53.7693 59.6884Z"
                stroke-width="2" class="loading"/>
        <!--镐子底下两点-->
        <path
                d="M19.7416 60.147L21.4566 62.0039L16.8852 64.5056"
                stroke="none"
                fill="#1370F3cf" class="loading-bling-1" />
        <!--镐子底下两点-->
        <path
                d="M11 61.67L12.8992 60.0019L15.2859 64.6345"
                stroke="none"
                fill="#1370F3cf" class="loading-bling-2" />
        <!--镐子红叉-->
        <path
                d="M10.2929 49.2929C9.90237 49.6834 9.90237 50.3166 10.2929 50.7071L20.2929 60.7071C20.6834 61.0976 21.3166 61.0976 21.7071 60.7071C22.0976 60.3166 22.0976 59.6834 21.7071 59.2929L11.7071 49.2929C11.3166 48.9024 10.6834 48.9024 10.2929 49.2929Z"
                stroke="none"
                fill="#ff1f1fcf" class="loading-icon" />
        <!--镐子红叉-->
        <path
                d="M10.2929 60.7071C10.6834 61.0976 11.3166 61.0976 11.7071 60.7071L21.7071 50.7071C22.0977 50.3166 22.0976 49.6835 21.7071 49.2929C21.3166 48.9024 20.6834 48.9024 20.2929 49.2929L10.2929 59.2929C9.9024 59.6834 9.9024 60.3166 10.2929 60.7071Z"
                stroke="none"
                fill="#ff1f1fcf"
                fill-rule="evenodd" class="loading-icon" />
    </svg><br>
    { loading_text }
</div>
<style>
    .ala {
        width: max-content;
        height: max-content;
        background-color: var(--dark-color);
        text-align: center;
        font-weight: bold;
        border-radius: 10px;
        stroke: var(--font-color);
        color: var(--font-color);
        transition: all 0.2s;
        padding: 12px 24px;
    }
    .ala:hover {
        box-shadow: 0 0 5px #0077ff;
    }
    .loading-svg {
        width: 80px;
        height: 80px;
    }
    .loading-icon {
        transition: opacity 0.2s;
        opacity: var(--icon-hidden);
    }
    .loading-bling-1 {
        animation: bling-1 1.3s var(--is-infinite);
        opacity: 0;
    }
    .loading-bling-2 {
        animation: bling-2 1.3s var(--is-infinite);
        opacity: 0;
    }
    .loading {
        animation: loading 1.3s linear var(--is-infinite);
        transform: rotate(8deg);
        transform-origin: 51px 59px;
    }
    @keyframes loading {
        0%, 100%{
            transform: rotate(8deg);
        }
        8% {
            transform: rotate(-10deg);
        }
        20% {
            transform: rotate(-65deg);
        }
        36% {
            transform: rotate(-8deg);
        }
        60% {
            transform: rotate(12deg);
        }
        80% {
            transform: rotate(4deg);
        }
    }
    @keyframes bling-1 {
        0%, 20%, 50%, 100% {
            opacity: 0;
        }
        20.1%, 30% {
            opacity: 1;
        }
        20% {
            transform: translate(0, 0);
        }
        50% {
            transform: translate(6px, -6px);
        }
    }
    @keyframes bling-2 {
        0%, 20%, 50%, 100% {
            opacity: 0;
        }
        20.1%, 30% {
            opacity: 1;
        }
        20% {
            transform: translate(0, 0);
        }
        50% {
            transform: translate(-6px, -6px);
        }
    }
</style>