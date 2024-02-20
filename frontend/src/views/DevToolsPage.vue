<template>
    <VContainer fluid style="overflow: scroll; height: 100%;">
        <VCard variant="text">
            <VCardTitle class="d-flex align-center ga-2 mb-4">
                <VBtn :prepend-icon="IconChevronLeft" variant="text" @click="router.back()">back</VBtn>
                <VIcon size="32px" :icon="IconCogFilledLoop" />
                <span class="text-h5">DevTools</span>
            </VCardTitle>
            <VCardText>
                <VExpansionPanels>
                    <VExpansionPanel>
                        <template #title>
                            <div class="d-flex align-center ga-2">
                                <VIcon :icon="IconComputer" />
                                Build Options (ldflags)
                            </div>
                        </template>
                        <template #text>
                            <code
                                class="selectable"><div v-for="(val, key) in opts" :key="key" density="compact">{{ key }}: {{ val }}</div></code>
                        </template>
                    </VExpansionPanel>
                </VExpansionPanels>
            </VCardText>
        </VCard>
    </VContainer>
</template>

<script lang="ts" setup async>
import * as options from '$app/opts/Options'
import * as models from '$app/models'
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { IconCogFilledLoop, IconChevronLeft, IconComputer } from '@iconify-prerendered/vue-line-md'

const router = useRouter()
const opts = ref<models.opts.Options | undefined>();

(async () => {
    opts.value = await options.Get()
})()
</script>

<style lang="scss"></style>