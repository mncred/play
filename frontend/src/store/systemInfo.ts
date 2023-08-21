import * as systemInfo from '$app/system/Info'

import { Ref, ref } from 'vue'

import { defineStore } from 'pinia'

const platforms = ['unknown', 'windows', 'darwin', 'linux'] as const
export type Platform = typeof platforms[number]

const arches = ['unknown', '386', 'amd64', 'arm', 'arm64'] as const
export type Arch = typeof arches[number]

export const useSystemInfo = defineStore('systemInfo', () => {
    const ready = ref(false)
    const fetchers = [] as Promise<void>[]

    const platform = ref('unknown') as Ref<Platform>
    fetchers.push((async () => {
        const actualPlatform = await systemInfo.Platform()
        if (platforms.includes(actualPlatform as Platform)) {
            platform.value = actualPlatform as Platform
        } else {
            platform.value = 'unknown'
        }
    })())

    const arch = ref('unknown') as Ref<Arch>
    fetchers.push((async () => {
        const actualArch = await systemInfo.Arch()
        if (arches.includes(actualArch as Arch)) {
            arch.value = actualArch as Arch
        } else {
            arch.value = 'unknown'
        }
    })())

    Promise.all(fetchers).then(() => {
        ready.value = true
        console.log('[STORE] systemInfo ready', platform.value, arch.value)
    })

    return { ready, platform, arch }
})