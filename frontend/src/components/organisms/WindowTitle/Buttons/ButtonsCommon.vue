<template>
    <div class="d-flex justify-center align-center">
        <VBtn :disabled="props.disableMinimize" size="x-small" density="compact" style="height: 100%;"
            @click="emit('minimize')">
            <VIcon size="x-large" :icon="IconMinimize" />
        </VBtn>
        <VBtn :disabled="props.disableMaximize" size="x-small" density="compact" style="height: 100%;"
            @click="emit('maximize')">
            <VIcon size="x-large" :icon="IconSquareOutline" />
        </VBtn>
        <VBtn :disabled="props.disableClose" size="x-small" density="compact" style="height: 100%;" @click="emit('close')">
            <VIcon size="x-large" :icon="IconClose" />
        </VBtn>
    </div>
</template>

<script lang="ts" setup async>
import { IconMinimize, IconSquareOutline, IconClose } from '@iconify-prerendered/vue-material-symbols-light'

const props = withDefaults(defineProps<{
    disableMinimize?: boolean
    disableMaximize?: boolean
    disableClose?: boolean
}>(), {
    disableMinimize: false,
    disableMaximize: false,
    disableClose: false,
})

const emit = defineEmits<{
    (e: 'minimize'): void
    (e: 'maximize'): void
    (e: 'close'): void
}>()
</script>

<style lang="scss" scoped></style>