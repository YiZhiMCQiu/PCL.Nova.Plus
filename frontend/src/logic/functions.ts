import {BrowserOpenURL} from "../../wailsjs/runtime";

export function DarkAndThemeToConst(darkMode: boolean, themeMode: number): string {
    let co = [[
        "linear-gradient(to left, rgb(10, 114, 189), rgb(11, 129, 213), rgb(10, 114, 189))",//-6, +8, -7, -7, +10, -8
        "linear-gradient(to left, rgb(40, 147, 132), rgb(45, 166, 148), rgb(40, 147, 132))",
        "linear-gradient(to left, rgb(58, 156, 78), rgb(66, 180, 73), rgb(58, 156, 78))",
        "linear-gradient(to left, rgb(117, 155, 42), rgb(131, 176, 47), rgb(117, 155, 42))",
        "linear-gradient(to left, rgb(144, 129, 74), rgb(161, 146, 85), rgb(144, 129, 74))",
        "url('/src/assets/images/Themes/BlackDark.png')",
        "url('/src/assets/images/Themes/FoolRainbowDark.png')",
        "url('/src/assets/images/Themes/PinkDark.png')",
        "url('/src/assets/images/Themes/PurpleDark.png')",
        "url('/src/assets/images/Themes/LuckyRainbowDark.png')",
        "url('/src/assets/images/Themes/GoldDark.png')",
        "url('/src/assets/images/Themes/OrangeDark.png')",
        "url('/src/assets/images/Themes/MojangRedDark.png')",
        "url('/src/assets/images/Themes/HackBlueDark.png')",
    ], [
        "linear-gradient(to left, rgb(16, 106, 196), rgb(18, 119, 221), rgb(16, 106, 196))",
        "linear-gradient(to left, rgb(46, 139, 139), rgb(52, 156, 156), rgb(46, 139, 139))",
        "linear-gradient(to left, rgb(64, 148, 85), rgb(73, 170, 81), rgb(64, 148, 85))",
        "linear-gradient(to left, rgb(123, 147, 49), rgb(138, 166, 55), rgb(123, 147, 49))",
        "linear-gradient(to left, rgb(150, 121, 81), rgb(168, 136, 93), rgb(150, 121, 81))",
        "url('/src/assets/images/Themes/Black.png')",
        "url('/src/assets/images/Themes/FoolRainbow.png')",
        "url('/src/assets/images/Themes/Pink.png')",
        "url('/src/assets/images/Themes/Purple.png')",
        "url('/src/assets/images/Themes/LuckyRainbow.png')",
        "url('/src/assets/images/Themes/Gold.png')",
        "url('/src/assets/images/Themes/Orange.png')",
        "url('/src/assets/images/Themes/MojangRed.png')",
        "url('/src/assets/images/Themes/HackBlue.png')",
    ]]
    return co[darkMode ? 0 : 1][themeMode - 1] ?? "linear-gradient(to left, rgb(16, 106, 196), rgb(18, 119, 221), rgb(16, 106, 196))"
}
export function DarkAndThemeToMain(darkMode: boolean, themeMode: number): string {
    let co = [[
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
    ], [
        "linear-gradient(to left bottom, rgb(150, 212, 235), rgb(179, 196, 241))",
        "linear-gradient(to left bottom, rgb(214, 237, 232), rgb(216, 233, 239))",
        "linear-gradient(to left bottom, rgb(223, 238, 222), rgb(220, 237, 227))",
        "linear-gradient(to left bottom, rgb(236, 237, 213), rgb(229, 239, 217))",
        "linear-gradient(to left bottom, rgb(238, 230, 226), rgb(235, 233, 223))",
        "linear-gradient(to left bottom, rgb(230, 232, 234), rgb(231, 231, 234))",
        "linear-gradient(to left bottom, rgb(255, 250, 254), rgb(255, 126, 143))",
        "linear-gradient(to left bottom, rgb(245, 219, 231), rgb(246, 223, 224))",
        "linear-gradient(to left bottom, rgb(234, 224, 242), rgb(239, 219, 240))",
        "linear-gradient(to left bottom, rgb(235, 223, 241), rgb(240, 220, 238))",
        "linear-gradient(to left bottom, rgb(243, 233, 215), rgb(238, 239, 207))",
        "linear-gradient(to left bottom, rgb(246, 226, 219), rgb(244, 232, 212))",
        "linear-gradient(to left bottom, rgb(246, 222, 224), rgb(245, 227, 219))",
        "linear-gradient(to left bottom, rgb(221, 229, 243), rgb(227, 226, 245))",
    ]]
    return co[darkMode ? 0 : 1][themeMode - 1] ?? "linear-gradient(to left bottom, rgb(150, 212, 235), rgb(179, 196, 241))"
}
export function OpenCustomURL(url: string) {
    BrowserOpenURL(url)
}