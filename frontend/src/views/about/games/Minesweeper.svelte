<script lang="ts">
    import MyTextInput from "../../../component/input/MyTextInput.svelte";
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import MyToggleSwitch from "../../../component/button/MyToggleSwitch.svelte";
    export let slide = null;
    export let after_leave = null;

    // 格子接口
    interface Grids {
        // 格子横坐标
        x: number;
        // 格子纵坐标
        y: number;
        // 格子数字（-1：雷、0：周围8格没有雷、1：有一个雷（剩下以此类推）、2、3、4、5、6、7、8）
        p: -1 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8;
        // 格子是否被标记（-1：已被点击、0：未标记、1：旗子、2：问号）
        // 3比较特殊，用于在开启扣分模式后，如果点到了雷则默认将该旗子标记为3，用于表示该格子不可点击，但是却被标了旗子。
        m: -1 | 0 | 1 | 2 | 3;
        // 格子应该被显示成什么（有炸弹、旗子、问号、1~8的数字，其中0不会显示）
        s: string;
        // 是否是当前点击的雷（用于判断点击该雷的时候背景变成红色。）
        c: boolean;
    }

    let score = 0;
    let time = 0;
    let flags = 0;
    let win = 0;

    let width = "";
    let height = "";
    let mines = "";

    // 格子列表
    let grids: Grids[][] = [];

    // 临时宽度
    let tempWidth = 0;
    // 临时高度
    let tempHeight = 0;
    // 临时雷数
    let tempMines = 0;
    // 棋盘记录雷数
    let mine_count = 0;

    // 是否开始游戏
    let start = false;
    // 是否锁住棋盘（无法被点击）
    let locked = false;
    // 扣分模式
    let cheat = false;

    setInterval(() => {
        if (start && !locked) {
            time++;
        }
    }, 1000);
    //初始化数组
    function init_array() {
        grids = [];
        for (let i = 0; i < tempHeight; i++) {
            grids[i] = [];
            for (let j = 0; j < tempWidth; j++) {
                grids[i][j] = {
                    x: j,
                    y: i,
                    p: 0,
                    m: 0,
                    s: "",
                    c: false,
                };
            }
        }
    }
    //生成雷
    function generate_mine() {
        for (let i = 0; i < tempMines; i++) {
            while (true) {
                let x = Math.floor(Math.random() * tempHeight);
                let y = Math.floor(Math.random() * tempWidth);
                if (grids[x][y].p == -1) {
                    continue;
                }
                grids[x][y].p = -1;
                break;
            }
        }
    }
    //初始化数字（为雷的周围生成数字）
    function init_number() {
        for (let i = 0; i < tempHeight; i++) {
            for (let j = 0; j < tempWidth; j++) {
                if (grids[i][j].p == -1) {
                    continue;
                }
                let foo = 0;
                if (i > 0 && j > 0) {
                    if (grids[i - 1][j - 1].p == -1) {
                        foo++;
                    }
                }
                if (i > 0 && j < tempWidth - 1) {
                    if (grids[i - 1][j + 1].p == -1) {
                        foo++;
                    }
                }
                if (i < tempHeight - 1 && j > 0) {
                    if (grids[i + 1][j - 1].p == -1) {
                        foo++;
                    }
                }
                if (i < tempHeight - 1 && j < tempWidth - 1) {
                    if (grids[i + 1][j + 1].p == -1) {
                        foo++;
                    }
                }
                if (i > 0) {
                    if (grids[i - 1][j].p == -1) {
                        foo++;
                    }
                }
                if (i < tempHeight - 1) {
                    if (grids[i + 1][j].p == -1) {
                        foo++;
                    }
                }
                if (j > 0) {
                    if (grids[i][j - 1].p == -1) {
                        foo++;
                    }
                }
                if (j < tempWidth - 1) {
                    if (grids[i][j + 1].p == -1) {
                        foo++;
                    }
                }
                grids[i][j].p = foo as 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8;
            }
        }
    }
    function startGame() {
        let w = parseInt(width);
        let h = parseInt(height);
        let m = parseInt(mines);
        if (
            Number.isNaN(w) ||
            Number.isNaN(h) ||
            Number.isNaN(m) ||
            w < 5 ||
            w > 100 ||
            h < 5 ||
            h > 100 ||
            m < 5 ||
            m > w * h - 10
        ) {
            return;
        }
        tempWidth = w;
        tempHeight = h;
        tempMines = m;
        init_array();
        generate_mine();
        init_number();
        start = false;
        time = 0;
        flags = m;
        win = 0;
        score = 0;
        locked = false;
        mine_count = 0;
    }

    //格子左键点击（揭开旗子）
    function grid_button_click(x: number, y: number) {
        if (locked) return;
        if (!start) start = true;
        if (![-1, 1, 3].includes(grids[y][x].m)) {
            grids[y][x].m = -1;
            grids[y][x].s =
                grids[y][x].p == -1
                    ? "💣"
                    : grids[y][x].p == 0
                      ? ""
                      : grids[y][x].p.toString();
            if (grids[y][x].p == -1) {
                if (!cheat) {
                    locked = true;
                    win = 2;
                    for (let k = 0; k < tempHeight; k++) {
                        for (let l = 0; l < tempWidth; l++) {
                            if (grids[k][l].p == -1) {
                                grids[k][l].s = "💣";
                            }
                        }
                    }
                    grids[y][x].c = true;
                    return;
                } else {
                    grids[y][x].c = true;
                    score -= 50;
                    grids[y][x].m = 3;
                    flags -= 1;
                }
            } else {
                mine_count++;
                score += 10;
            }
            if (mine_count == tempHeight * tempWidth - tempMines) {
                locked = true;
                win = 1;
                return;
            }
            if (grids[y][x].p != 0) return;
            if (x > 0 && y > 0) {
                grid_button_click(x - 1, y - 1);
            }
            if (x > 0 && y < tempHeight - 1) {
                grid_button_click(x - 1, y + 1);
            }
            if (x < tempWidth - 1 && y > 0) {
                grid_button_click(x + 1, y - 1);
            }
            if (x < tempWidth - 1 && y < tempHeight - 1) {
                grid_button_click(x + 1, y + 1);
            }
            if (x > 0) {
                grid_button_click(x - 1, y);
            }
            if (x < tempWidth - 1) {
                grid_button_click(x + 1, y);
            }
            if (y > 0) {
                grid_button_click(x, y - 1);
            }
            if (y < tempHeight - 1) {
                grid_button_click(x, y + 1);
            }
        }
    }

    //格子右键点击（标旗或者揭开【周围】格子）
    function grid_button_right(x: number, y: number) {
        if (locked) return;
        switch (grids[y][x].m) {
            case -1:
                let cc = grids[y][x].p;
                let k = 0;
                if (x > 0 && y > 0)
                    if ([1, 3].includes(grids[y - 1][x - 1].m)) k++;
                if (x > 0 && y < tempHeight - 1)
                    if ([1, 3].includes(grids[y + 1][x - 1].m)) k++;
                if (x < tempWidth - 1 && y > 0)
                    if ([1, 3].includes(grids[y - 1][x + 1].m)) k++;
                if (x < tempWidth - 1 && y < tempHeight - 1)
                    if ([1, 3].includes(grids[y + 1][x + 1].m)) k++;
                if (x > 0) if ([1, 3].includes(grids[y][x - 1].m)) k++;
                if (x < tempWidth - 1)
                    if ([1, 3].includes(grids[y][x + 1].m)) k++;
                if (y > 0) if ([1, 3].includes(grids[y - 1][x].m)) k++;
                if (y < tempHeight - 1)
                    if ([1, 3].includes(grids[y + 1][x].m)) k++;
                // 如果格子本身的数字小于或者等于周围插旗子的数字，则开启
                if (cc > k) return;
                if (x > 0 && y > 0) grid_button_click(x - 1, y - 1);
                if (x > 0 && y < tempHeight - 1)
                    grid_button_click(x - 1, y + 1);
                if (x < tempWidth - 1 && y > 0) grid_button_click(x + 1, y - 1);
                if (x < tempWidth - 1 && y < tempHeight - 1)
                    grid_button_click(x + 1, y + 1);
                if (x > 0) grid_button_click(x - 1, y);
                if (x < tempWidth - 1) grid_button_click(x + 1, y);
                if (y > 0) grid_button_click(x, y - 1);
                if (y < tempHeight - 1) grid_button_click(x, y + 1);
                break;
            case 0:
                if (flags > 0) {
                    grids[y][x].m = 1;
                    grids[y][x].s = "🚩";
                    flags--;
                } else {
                    grids[y][x].m = 2;
                    grids[y][x].s = "❔";
                }
                break;
            case 1:
                grids[y][x].m = 2;
                grids[y][x].s = "❔";
                flags++;
                break;
            case 2:
                grids[y][x].m = 0;
                grids[y][x].s = "";
                break;
            default:
                break;
        }
    }
</script>

<div class="component-minesweeper" in:slide out:slide on:outroend={after_leave}>
    <div class="bar">
        <div style:width="100%">
            <MyNormalSpan>场地宽度</MyNormalSpan>
        </div>
        <MyTextInput
            placeholder="[5, 100] 区间范围"
            value={width}
            on:blur={(e) => (width = e.detail.value)}
            style_in="width: 180px; height: 25px; font-size: 15px; transition: all 0.2s;"
        />
        <div style:width="100%">
            <MyNormalSpan>场地高度</MyNormalSpan>
        </div>
        <MyTextInput
            placeholder="[5, 100] 区间范围"
            value={height}
            on:blur={(e) => (height = e.detail.value)}
            style_in="width: 180px; height: 25px; font-size: 15px; transition: all 0.2s;"
        />
        <div style:width="100%">
            <MyNormalSpan>场地雷数</MyNormalSpan>
        </div>
        <MyTextInput
            placeholder="[5, w*h-10] 区间范围"
            value={mines}
            on:blur={(e) => (mines = e.detail.value)}
            style_in="width: 180px; height: 25px; font-size: 15px; transition: all 0.2s;"
        />
        <div>
            <MyNormalButton
                style_in="width: 95px; height: 25px; margin-top: 5px"
                on:click={() => {
                    width = "9";
                    height = "9";
                    mines = "10";
                }}
            >
                默认条件
            </MyNormalButton>
            <MyNormalButton
                style_in="width: 95px; height: 25px; margin-top: 5px"
                on:click={startGame}
            >
                开始游戏
            </MyNormalButton>
        </div>
        <div
            style="width: 100%; display: flex; align-items: center; height: 30px; justify-content: space-around; margin-top: 10px"
        >
            <MyNormalSpan>扣分模式</MyNormalSpan>
            <MyToggleSwitch
                on:click={() => (cheat = !cheat)}
                isSelect={cheat}
            />
        </div>
        <div class="info">
            <MyNormalSpan style_in="margin-left: 10px; font-size: 24px;"
                >分数：{score}</MyNormalSpan
            >
            <MyNormalSpan style_in="margin-left: 10px; font-size: 24px;"
                >时间：{time}</MyNormalSpan
            >
            <MyNormalSpan style_in="margin-left: 10px; font-size: 24px;"
                >旗子：{flags}</MyNormalSpan
            >
            <MyNormalSpan
                style_in={"font-size: 50px; " +
                    (win == 1
                        ? "color: green;"
                        : win == 2
                          ? "color: red;"
                          : "")}
                >{win == 1 ? "胜利" : win == 2 ? "失败" : ""}</MyNormalSpan
            >
        </div>
    </div>
    <div class="chess">
        {#each grids as grid1}
            {#each grid1 as grid}
                <button
                    class={[-1, 3].includes(grid.m)
                        ? grid.c
                            ? "grid-button-red"
                            : "grid-button-click"
                        : "grid-button"}
                    style={"left: " +
                        grid.x * 25 +
                        "px; top: " +
                        grid.y * 25 +
                        "px; color: " +
                        (grid.p == 1
                            ? "rgb(0, 128, 255)"
                            : grid.p == 2
                              ? "rgb(0, 128, 0)"
                              : grid.p == 3
                                ? "rgb(192, 0, 0)"
                                : grid.p == 4
                                  ? "rgb(0, 0, 96)"
                                  : grid.p == 5
                                    ? "rgb(128, 0, 48)"
                                    : grid.p == 6
                                      ? "rgb(0, 96, 96)"
                                      : grid.p == 7
                                        ? "rgb(10, 10, 10)"
                                        : grid.p == 8
                                          ? "rgb(100, 100, 100)"
                                          : "black") +
                        ";"}
                    on:click={() => grid_button_click(grid.x, grid.y)}
                    on:contextmenu={() => grid_button_right(grid.x, grid.y)}
                >
                    {grid.s}
                </button>
            {/each}
        {/each}
    </div>
</div>

<style>
    .component-minesweeper {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        overflow: hidden;
    }
    .bar {
        width: 200px;
        height: calc(100% - 20px);
        padding: 10px;
        border-right: 2px solid gray;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .info {
        margin-top: 10px;
        display: flex;
        flex-direction: column;
        align-items: start;
        width: 100%;
    }
    .chess {
        width: calc(100% - 222px);
        height: calc(100% - 1px);
        position: relative;
        overflow: auto;
        flex-shrink: 0;
    }
    .grid-button,
    .grid-button-click,
    .grid-button-red {
        position: absolute;
        width: 25px;
        height: 25px;
        border: 1px solid black;
        box-sizing: border-box;
        background-color: lightgray;
        font-size: 16px;
        cursor: pointer;
        font-weight: bold;
    }
    .grid-button-click,
    .grid-button-red {
        background-color: gray;
        cursor: default;
    }
    .grid-button-red {
        background-color: red;
    }
    .grid-button:hover {
        background-color: silver;
    }
    .grid-button:active {
        background-color: darkgray;
    }
</style>
