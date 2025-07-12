// 以下是操作 MC 版本数组的类
import {writable} from "svelte/store";

export const mc_list_ok = writable(false)
export const latest_release = writable({})
export const latest_snapshot = writable({})
export const mc_list_release = writable([])
export const mc_list_snapshot = writable([])
export const mc_list_old_beta = writable([])
export const mc_list_old_alpha = writable([])
export const forge_list = writable({})
export const forge_support = writable([])
export const fabric_list = writable([])
export const quilt_list = writable([])
export const neoforge_list = writable([])
export const optifine_list = writable([])
export const liteloader_list = writable([])