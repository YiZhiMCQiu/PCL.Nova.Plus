import {writable} from "svelte/store";
// 以下是部分常量
export const MSG_INFO = 0
export const MSG_WARNING = 1
export const MSG_ERROR = 2

export const HNT_INFO = 0
export const HNT_WARNING = 1
export const HNT_ERROR = 2
export const HNT_PASS = 3
// 下列是信息框
export let b_title = writable("")
export let b_content = writable("")
export let b_level = writable(0)
export let b_button = writable(["ok"])
export let b_show_all = writable(false)
export let b_resolve = writable<((value: number) => void) | null>(null)

export function messagebox(title: string, content: string, level: number = MSG_INFO, button: string[] = ["ok"]): Promise<number> {
    b_title.set(title)
    b_content.set(content)
    b_level.set(level)
    b_button.set(button)
    b_show_all.set(false)
    return new Promise((resolve: (value: number) => void) => {
        b_show_all.set(true)
        b_resolve.set(resolve)
    })
}

// 下列是提示框
export let h_content = writable("")
export let h_level = writable(0)
let q = false
export function showHint(content: string, level: number = HNT_INFO): void {
    h_content.set(content)
    h_level.set(level)
    if(!q) {
        q = true
        setTimeout(() => {
            h_content.set("")
            h_level.set(0)
            q = false
        }, 3000)
    }
}

export let i_title = writable("")
export let i_content = writable("")
export let i_level = writable(0)
export let i_show_all = writable(false)
export let i_placeholder = writable("")
export let i_resolve = writable<((value: string) => void) | null>(null)

export function inputbox(title: string, content: string, level: number = MSG_INFO, placeholder: string = ""): Promise<string> {
    i_title.set(title)
    i_content.set(content)
    i_level.set(level)
    i_show_all.set(false)
    i_placeholder.set(placeholder)
    return new Promise((resolve: (value: string) => void) => {
        i_show_all.set(true)
        i_resolve.set(resolve)
    })
}