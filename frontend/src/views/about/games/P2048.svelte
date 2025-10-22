<script lang="ts">
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    import { onMount } from "svelte";
    import { current_view, current_about } from "../../../store/changeBody";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    export let slide = null;
    export let after_leave = null;

    let start = false;
    let score = 0;
    let step = 0;

    // 棋盘
    let tds = [
        [2, 32, 16, 2048],
        [256, 4, 64, 32],
        [64, 512, 8, 128],
        [256, 128, 1024, 16],
    ];

    let success = "";

    function genRandomChess(bc: boolean): boolean {
        let count = 0;
        for (let c = 0; c < 4; c++) {
            for (let d = 0; d < 4; d++) {
                if (tds[c][d] == 0) count++;
            }
        }
        if (count == 0) return false;
        let rand = Math.floor(Math.random() * count);
        count = 0;
        for (let c = 0; c < 4; c++) {
            for (let d = 0; d < 4; d++) {
                if (tds[c][d] == 0) {
                    if (count == rand) {
                        if (bc) {
                            let rand2 = Math.floor(Math.random() * 3);
                            switch (rand2) {
                                case 0:
                                    tds[c][d] = 256;
                                    break;
                                case 1:
                                    tds[c][d] = 512;
                                    break;
                                case 2:
                                    tds[c][d] = 1024;
                                    break;
                            }
                        } else {
                            tds[c][d] = 2;
                        }
                        return true;
                    }
                    count++;
                }
            }
        }
        return false;
    }

    function compArray(input: number[]): number[] {
        let arr = input.concat();
        let n = arr.length;
        for (let i = 0; i < n; ) {
            if (arr[i] == 0) {
                i++;
                continue;
            }
            let merged = false;
            if (i + 1 < n && arr[i + 1] == arr[i]) {
                //首先找一遍是否相邻
                score += arr[i];
                arr[i] *= 2;
                if (arr[i] == 2048) {
                    success = "你赢了！";
                    start = false;
                }
                arr[i + 1] = 0;
                i += 2;
                merged = true;
            } else {
                //如果不相邻
                let j = i + 1;
                while (j < n && arr[j] == 0) j++;
                if (j < n && arr[j] == arr[i]) {
                    let canMerge = true;
                    for (let k = i + 1; k < j; k++) {
                        if (arr[k] != 0) {
                            canMerge = false;
                            break;
                        }
                    }
                    if (canMerge) {
                        score += arr[i];
                        arr[i] *= 2;
                        if (arr[i] == 2048) {
                            success = "你赢了！";
                            start = false;
                        }
                        arr[j] = 0;
                        i = j + 1;
                        merged = true;
                    }
                }
            }
            if (!merged) i++;
        }
        //合并数字
        let w = 0;
        for (let i = 0; i < n; i++) {
            if (arr[i] != 0) arr[w++] = arr[i];
        }
        while (w < n) arr[w++] = 0;
        return arr;
    }
    function up() {
        if (start) {
            for (let c = 0; c < 4; c++) {
                let arr = [];
                for (let d = 0; d < 4; d++) {
                    arr.push(tds[d][c]);
                }
                let p = compArray(arr);
                for (let d = 0; d < 4; d++) {
                    tds[d][c] = p[d];
                }
            }
            if (genRandomChess(false)) {
                step++;
                if (success != "你赢了！") {
                    success = "";
                }
            } else {
                success = "无法往上";
            }
        }
    }
    function down() {
        if (start) {
            for (let c = 0; c < 4; c++) {
                let arr = [];
                for (let d = 3; d >= 0; d--) {
                    arr.push(tds[d][c]);
                }
                let p = compArray(arr);
                for (let d = 0, e = 3; d < 4; d++, e--) {
                    tds[e][c] = p[d];
                }
            }
            if (genRandomChess(false)) {
                step++;
                if (success != "你赢了！") {
                    success = "";
                }
            } else {
                success = "无法往下";
            }
        }
    }
    function left() {
        if (start) {
            for (let c = 0; c < 4; c++) {
                let arr = [];
                for (let d = 0; d < 4; d++) {
                    arr.push(tds[c][d]);
                }
                let p = compArray(arr);
                for (let d = 0; d < 4; d++) {
                    tds[c][d] = p[d];
                }
            }
            if (genRandomChess(false)) {
                step++;
                if (success != "你赢了！") {
                    success = "";
                }
            } else {
                success = "无法往左";
            }
        }
    }
    function right() {
        if (start) {
            for (let c = 0; c < 4; c++) {
                let arr = [];
                for (let d = 3; d >= 0; d--) {
                    arr.push(tds[c][d]);
                }
                let p = compArray(arr);
                for (let d = 0, e = 3; d < 4; d++, e--) {
                    tds[c][e] = p[d];
                }
            }
            if (genRandomChess(false)) {
                step++;
                if (success != "你赢了！") {
                    success = "";
                }
            } else {
                success = "无法往右";
            }
        }
    }
    function cheatGen() {
        if (start) {
            if (genRandomChess(true)) {
                success = "作弊成功";
            } else {
                success = "无法生成";
            }
        }
    }

    function startGame() {
        score = 0;
        step = 0;
        tds = [
            [0, 0, 0, 0],
            [0, 0, 0, 0],
            [0, 0, 0, 0],
            [0, 0, 0, 0],
        ];
        start = true;
        success = "";
        genRandomChess(false);
        genRandomChess(false);
    }

    onMount(() => {
        document.addEventListener("keyup", (e: KeyboardEvent) => {
            if ($current_view === "about" && $current_about === "P2048") {
                if (
                    (e.code as string) === "ArrowUp" ||
                    (e.code as string) === "KeyW"
                ) {
                    up();
                } else if (
                    (e.code as string) === "ArrowLeft" ||
                    (e.code as string) === "KeyA"
                ) {
                    left();
                } else if (
                    (e.code as string) === "ArrowRight" ||
                    (e.code as string) === "KeyD"
                ) {
                    right();
                } else if (
                    (e.code as string) === "ArrowDown" ||
                    (e.code as string) === "KeyS"
                ) {
                    down();
                }
            }
        });
    });
</script>

<div class="component-p2048" in:slide out:slide on:outroend={after_leave}>
    <div class="score">
        <MyNormalButton
            style_in="width: calc(100% - 20px); height: 100px; font-size: 30px; font-weight: bold; margin: 10px;"
            on:click={startGame}>点我开始</MyNormalButton
        >
        <MyNormalSpan
            style_in="font-size: 30px; font-weight: bold; margin: 10px;"
            >分数：{score}</MyNormalSpan
        >
        <MyNormalSpan
            style_in="font-size: 30px; font-weight: bold; margin: 10px;"
            >步数：{step}</MyNormalSpan
        >
        <MyNormalButton
            style_in="width: calc(100% - 20px); height: 100px; font-size: 30px; font-weight: bold; margin: 10px;"
            on:click={cheatGen}>作弊生成</MyNormalButton
        >
        <MyNormalSpan
            style_in={"font-size: 30px; font-weight: bold; margin: 10px;" +
                (success == "你赢了！" ? "color: green;" : "color: red;")}
            >{success}</MyNormalSpan
        >
    </div>
    <div class="chess">
        <div class="chess-content">
            <table>
                <tbody>
                    <tr>
                        <td></td>
                        <td
                            class="chess-direction chess-direction-tb"
                            on:click={up}
                            on:keydown|preventDefault
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                            >
                                <path d="M12 20V4m0 0l6 6m-6-6l-6 6" />
                            </svg>
                        </td>
                        <td></td>
                    </tr>
                    <tr>
                        <td
                            class="chess-direction chess-direction-lr"
                            on:click={left}
                            on:keydown|preventDefault
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                            >
                                <path d="M20 12H4m0 0l6-6m-6 6l6 6" />
                            </svg>
                        </td>
                        <td class="chess-content-center font-pcl">
                            <table>
                                <tbody>
                                    {#each tds as td1}
                                        <tr>
                                            {#each td1 as td}
                                                <td
                                                    style={"background-color: " +
                                                        (td == 0
                                                            ? "#cdc1b3"
                                                            : td == 2
                                                              ? "#efe5db"
                                                              : td == 4
                                                                ? "#eddcbe"
                                                                : td == 8
                                                                  ? "#f3b079"
                                                                  : td == 16
                                                                    ? "#f7925c"
                                                                    : td == 32
                                                                      ? "#f57656"
                                                                      : td == 64
                                                                        ? "#f4522c"
                                                                        : td ==
                                                                            128
                                                                          ? "#edce71"
                                                                          : td ==
                                                                              256
                                                                            ? "#e6d151"
                                                                            : td ==
                                                                                512
                                                                              ? "#1200ff"
                                                                              : td ==
                                                                                  1024
                                                                                ? "#cc00ff"
                                                                                : td ==
                                                                                    2048
                                                                                  ? "#000000"
                                                                                  : "white") +
                                                        "; font-size: " +
                                                        (td == 2 ||
                                                        td == 4 ||
                                                        td == 8
                                                            ? "50px"
                                                            : td == 16 ||
                                                                td == 32 ||
                                                                td == 64
                                                              ? "40px"
                                                              : td == 128 ||
                                                                  td == 256 ||
                                                                  td == 512
                                                                ? "30px"
                                                                : td == 1024 ||
                                                                    td == 2048
                                                                  ? "25px"
                                                                  : "1000px") +
                                                        "; color: " +
                                                        (td == 2 ||
                                                        td == 4 ||
                                                        td == 512 ||
                                                        td == 128
                                                            ? "#7a6d65"
                                                            : td == 8 ||
                                                                td == 16 ||
                                                                td == 32 ||
                                                                td == 64 ||
                                                                td == 256 ||
                                                                td == 1024 ||
                                                                td == 2048
                                                              ? "#ffffff"
                                                              : "#000000")}
                                                    >{td === 0 ? "" : td}</td
                                                >
                                            {/each}
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </td>
                        <td
                            class="chess-direction chess-direction-lr"
                            on:click={right}
                            on:keydown|preventDefault
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="32"
                                height="32"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                            >
                                <path d="M4 12h16m0 0l-6-6m6 6l-6 6" />
                            </svg>
                        </td>
                    </tr>
                    <tr>
                        <td></td>
                        <td
                            class="chess-direction chess-direction-tb"
                            on:click={down}
                            on:keydown|preventDefault
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                            >
                                <path d="M12 4v16m0 0l6-6m-6 6l-6-6" />
                            </svg>
                        </td>
                        <td></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</div>

<style>
    .component-p2048 {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        overflow: hidden;
    }
    .score {
        flex-shrink: 0;
        width: 200px;
        height: 100%;
        border-right: 2px solid gray;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .chess {
        height: 100%;
        flex: 1;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    .chess-content {
        width: calc(100vh - 100px);
        height: calc(100vh - 100px);
    }
    .chess-direction {
        stroke: gray;
        cursor: pointer;
        transition: stroke 0.2s;
    }
    .chess-direction:hover {
        stroke: dimgray;
    }
    .chess-direction-tb {
        width: 80%;
        height: 10%;
    }
    .chess-direction-lr {
        width: 10%;
        height: 80%;
    }
    .chess-direction svg {
        width: 32px;
        height: 32px;
    }
    .chess-content > table {
        table-layout: fixed;
        width: 100%;
        height: 100%;
    }
    .chess-content > table tr td {
        text-align: center;
    }
    .chess-content-center {
        width: 80%;
        height: 80%;
        border: 2px solid black;
        background-color: rgb(192, 175, 157);
        border-radius: 10px;
        position: relative;
        padding: 3px;
    }
    .chess-content-center > table {
        table-layout: fixed;
        width: 100%;
        height: 100%;
    }
    .chess-content-center > table tr td {
        border-radius: 10px;
        font-weight: bold;
        width: 24%;
        height: 24%;
    }
</style>
