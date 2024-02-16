<template>
    <template v-if="props.appearance === 'macos'">
        <VAppBar :elevation="0" :height="28" :style="{ '--wails-draggable': props.draggable ? 'drag' : 'none' }">
            <div style="position: absolute; width: 100%; display: flex; justify-content: center;">
                <VIcon v-if="props.icon" class="mr-1" size="small" :icon="props.icon" :color="props.iconColor" />
                <span style="font-weight: bold; font-size: .8rem; opacity: .7;">
                    {{ buildInfo.windowTitle }}
                </span>
            </div>
            <div class="fill-height d-flex ga-2 mt-3 ml-2">
                <VBtn :disabled="!!buildInfo.windowDisableClose" size="x-small" density="compact"
                    :color="!!buildInfo.windowDisableClose ? 'gray' : 'red'" :icon="IconCircle" @click="onClose" />
                <VBtn :disabled="!!buildInfo.windowDisableMinimize" size="x-small" density="compact"
                    :color="!!buildInfo.windowDisableMinimize ? 'gray' : 'yellow'" :icon="IconCircle" @click="onMinimize" />
                <VBtn :disabled="!!buildInfo.windowDisableMaximize" size="x-small" density="compact"
                    :color="!!buildInfo.windowDisableMaximize ? 'gray' : 'green'" :icon="IconCircle" @click="onMaximize" />
            </div>
            <VSpacer />
            <Breadcrumbs style="font-weight: bold; font-size: .8rem; opacity: .7;" />
        </VAppBar>
    </template>
    <template v-else>
        <VAppBar :elevation="0" :height="32" :style="{ '--wails-draggable': props.draggable ? 'drag' : 'none' }">
            <VIcon v-if="props.icon" class="ml-2" size="small" :icon="props.icon" :color="props.iconColor" />
            <div class="ml-2 d-flex align-center ga-4" style="font-weight: bold; font-size: .8rem; opacity: .7;">
                <div>{{ buildInfo.windowTitle }}</div>
                <Breadcrumbs />
            </div>
            <VSpacer />
            <div class="fill-height d-flex">
                <VBtn :disabled="!!buildInfo.windowDisableMinimize" size="x-small" density="compact" style="height: 100%;"
                    @click="onMinimize">
                    <VIcon size="x-large" :icon="IconMinimize" />
                </VBtn>
                <VBtn :disabled="!!buildInfo.windowDisableMaximize" size="x-small" density="compact" style="height: 100%;"
                    @click="onMaximize">
                    <VIcon size="x-large" :icon="IconSquareOutline" />
                </VBtn>
                <VBtn :disabled="!!buildInfo.windowDisableClose" size="x-small" density="compact" style="height: 100%;"
                    @click="onClose">
                    <VIcon size="x-large" :icon="IconClose" />
                </VBtn>
            </div>
        </VAppBar>
    </template>
</template>

<script lang="ts" setup async>
import * as runtime from '$runtime'
import * as build from '$app/build/Build'
import { computed, ref, Ref, Component } from 'vue'
import { IconCircle, IconMinimize, IconSquareOutline, IconClose } from '@iconify-prerendered/vue-material-symbols-light'
import Breadcrumbs from './Breadcrumbs.vue'

const buildInfo = await build.Info()

type Design = 'common' | 'macos'

const props = withDefaults(defineProps<{
    /** Allow window dragging by this element */
    draggable: boolean,
    /** Appearance preference like 'macos' or something. Automatically if not specified. */
    appearance?: Design
    /** Show application icon */
    icon?: any
    /** Application icon color */
    iconColor?: string
}>(), {
    draggable: false,
    backgroundColor: undefined,
    color: '#000000',
    appearance: undefined,
})

const onClose = async () => runtime.Quit()
const onMaximize = async () => runtime.WindowToggleMaximise()
const onMinimize = async () => runtime.WindowMinimise()
</script>

<style lang="scss" scoped></style>