import {writable} from "svelte/store";

export let b_title = writable("")
export let b_content = writable("")
export let b_level = writable(0)
export let b_button = writable(["ok"])
export let b_show_all = writable(false)
export let b_resolve = writable<((value: number) => void) | null>(null)

export function messagebox(title: string, content: string, level: number = 0, button: string[] = ["ok"]): Promise<number> {
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
