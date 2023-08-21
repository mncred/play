import * as configProvider from '$app/config/Provider'

import { Ref, ref } from 'vue'

import { config } from '$app/models'
import { defineStore } from 'pinia'

export const useConfigProvider = defineStore('configProvider', () => {
    const ready = ref(false)
    let cfg = ref({}) as Ref<config.Config>
    configProvider.Get().then(conf => {
        cfg.value = conf
        ready.value = true
        console.log('[STORE] config ready', conf)
    })

    return { get: cfg }
})