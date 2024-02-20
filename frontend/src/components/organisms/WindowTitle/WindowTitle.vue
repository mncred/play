<template>
    <template v-if="appearance === 'mac'">
        <VAppBar :elevation="0" :height="28" :style="{ '--wails-draggable': props.draggable ? 'drag' : 'none' }">
            <div style="position: absolute; width: 100%; display: flex; justify-content: center;">
                <VIcon v-if="props.icon" class="mr-1" size="small" :icon="props.icon" :color="props.iconColor" />
                <span style="font-weight: bold; font-size: .8rem; opacity: .7;">
                    {{ opts.windowTitle }}
                </span>
            </div>
            <Buttons class="fill-height ml-2" style="position: relative; top: -2px;" type="mac"
                :disableMinimize="disableMinimize" :disableMaximize="disableMaximize" :disableClose="disableClose"
                @minimize="onMinimize" @maximize="onMaximize" @close="onClose" />
            <VSpacer />
            <Breadcrumbs style="font-weight: bold; font-size: .8rem; opacity: .7;" />
        </VAppBar>
    </template>
    <template v-else>
        <VAppBar :elevation="0" :height="32" :style="{ '--wails-draggable': props.draggable ? 'drag' : 'none' }">
            <VIcon v-if="props.icon" class="ml-2" size="small" :icon="props.icon" :color="props.iconColor" />
            <div class="ml-2 d-flex align-center ga-4" style="font-weight: bold; font-size: .8rem; opacity: .7;">
                <div>{{ opts.windowTitle }}</div>
                <Breadcrumbs />
            </div>
            <VSpacer />
            <Buttons class="fill-height" type="common" :disableMinimize="disableMinimize" :disableMaximize="disableMaximize"
                :disableClose="disableClose" @minimize="onMinimize" @maximize="onMaximize" @close="onClose" />
        </VAppBar>
    </template>
</template>

<script lang="ts" setup async>
import * as runtime from '$runtime'
import * as options from '$app/opts/Options'
import { UAParser } from 'ua-parser-js'
import Buttons from './Buttons/Buttons.vue'
import { computed, ref, Ref, Component } from 'vue'
import Breadcrumbs from './Breadcrumbs.vue'

const uaparser = new UAParser()
const opts = await options.Get()

const props = withDefaults(defineProps<{
    /** Allow window dragging by this element */
    draggable?: boolean,
    /** Show application icon */
    icon?: any
    /** Application icon color */
    iconColor?: string
}>(), {
    draggable: true,
})

const appearance = ref(uaparser.getOS().name === 'Mac OS' ? 'mac' : 'common')
const disableMinimize = !!opts.windowDisableMinimize
const disableMaximize = !!opts.windowDisableMaximize
const disableClose = !!opts.windowDisableClose

const onClose = async () => runtime.Quit()
const onMaximize = async () => runtime.WindowToggleMaximise()
const onMinimize = async () => runtime.WindowMinimise()
</script>

<style lang="scss" scoped></style>