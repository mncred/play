<template>
    <template v-if="props.type === 'mac'">
        <ButtonsMac />
    </template>
    <template v-else>
        <ButtonsCommon />
    </template>
</template>

<script lang="ts" setup async>
import ButtonsCommon from './ButtonsCommon.vue'
import ButtonsMac from './ButtonsMac.vue'

const props = withDefaults(defineProps<{
    type?: 'common' | 'mac'
}>(), {
    type: 'common'
})
</script>

<style lang="scss" scoped></style>