<template>
    <VBreadcrumbs :items="items" density="compact">
        <template #title="{ item }">
            <template v-if="item.pageName">
                <span class="text">
                    {{ $t(`pages.${item.title}.title`) }}
                </span>
            </template>
            <template v-else>
                <span class="text">
                    {{ item.title }}
                </span>
            </template>
        </template>
        <template #divider>
            <VIcon :size="20" :icon="IconChevronRightRounded" style="margin: 0 0 2px 0;" />
        </template>
    </VBreadcrumbs>
</template>

<script lang="ts" setup async>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { IconChevronRightRounded } from '@iconify-prerendered/vue-material-symbols-light'

const router = useRouter()
const route = useRoute()

const items = computed(() => {
    return route.matched.map(r => {
        return {
            title: r.name?.toString() || '',
            pageName: r.name?.toString(),
            disabled: false,
        }
    })
})
</script>

<style lang="scss" scoped>
.text {
    font-weight: normal;
}
</style>