<template>
    <QBar
        :style="{ '--wails-draggable': props.draggable ? 'drag' : 'none', 'background-color': props.backgroundColor ?? 'transparent' }">
        <template v-if="design === 'macos'">
            <QBtn :disable="props.hideClose" dense flat round icon="mdi-circle" size="8px"
                :color="props.hideClose ? 'gray' : 'red'" @click="onClose" />
            <QBtn dense flat round icon="mdi-circle" size="8px" color="yellow" @click="onMinimize" />
            <QBtn :disable="props.hideMaximise" dense flat round icon="mdi-circle" size="8px"
                :color="props.hideMaximise ? 'gray' : 'green'" @click="onMaximize" />
            <QSpace />
            <div :style="{ 'color': props.color }">{{ title }}</div>
            <QSpace />
        </template>
        <template v-else>
            <div :style="{ 'color': props.color }">{{ title }}</div>
            <QSpace />
            <QBtn dense flat icon="mdi-window-minimize" @click="onMinimize" />
            <QBtn v-if="!props.hideMaximise" dense flat icon="mdi-window-maximize" @click="onMaximize" />
            <QBtn v-if="!props.hideClose" dense flat icon="mdi-window-close" @click="onClose" />
        </template>
    </QBar>
</template>

<script lang="ts" setup>
import * as runtime from '$runtime'
import { useSystemInfo } from '@/store/systemInfo'
import { computed, ref, Ref } from 'vue'
import {
    QBar, QIcon, QSpace, QBtn
} from 'quasar'

type Design = 'common' | 'macos'

const props = withDefaults(defineProps<{
    /** Allow window dragging by this element */
    draggable: boolean,
    /** Defines custom text instead of default document.title */
    title?: string,
    /** Background color for header. Transparent if not specified */
    backgroundColor?: string,
    /** Title text color */
    color?: string,
    /** Appearance preference like 'macos' or something. Automatically if not specified. */
    appearance?: Design
    /** Hide window close button */
    hideClose?: boolean
    /** Hide window maximise button */
    hideMaximise?: boolean
}>(), {
    draggable: false,
    title: undefined,
    backgroundColor: undefined,
    color: '#000000',
    appearance: undefined,
    hideClose: false,
    hideMaximise: false,
})

const systemInfo = useSystemInfo()

const title = computed(() => props.title ?? document.title)

const design = computed(() => {
    if (props.appearance) {
        return props.appearance
    } else {
        return systemInfo.platform === 'darwin' ? 'macos' : 'common'
    }
})

const onClose = async () => runtime.Quit()
const onMaximize = async () => runtime.WindowToggleMaximise()
const onMinimize = async () => {
    /* MacOS can't handle window minimization */
    runtime.WindowMinimise()
    // await new Promise(r => setTimeout(r, 1000))
    // runtime.WindowUnminimise()
}
</script>

<style lang="scss"></style>