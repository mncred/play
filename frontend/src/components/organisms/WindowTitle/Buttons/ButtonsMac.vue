<template>
    <div class="d-flex justify-center align-center ga-2">
        <VBtn :disabled="props.disableClose" size="x-small" density="compact" :color="props.disableClose ? 'gray' : 'red'"
            :icon="IconCircle" @click="emit('close')" />
        <VBtn :disabled="props.disableMinimize" size="x-small" density="compact"
            :color="props.disableMinimize ? 'gray' : 'yellow'" :icon="IconCircle" @click="emit('minimize')" />
        <VBtn :disabled="props.disableMaximize" size="x-small" density="compact"
            :color="props.disableMaximize ? 'gray' : 'green'" :icon="IconCircle" @click="emit('maximize')" />
    </div>
</template>

<script lang="ts" setup async>
import { IconCircle } from '@iconify-prerendered/vue-material-symbols-light'

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