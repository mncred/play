<template>
  <QLayout
    :style="{ 'border-radius': fancyWindow ? (`${config.get?.window?.appearance?.roundness}px` ?? '0') : 'none', 'background-color': config.get?.window?.appearance?.background?.color ?? '#000000FF' }"
    view="lHh lpr lFf" container>
    <QHeader v-if="!config.get?.window?.header?.native" style="background: rgba(0, 0, 0, 0)">
      <WindowTitle draggable :hideMaximise="!config.get?.window?.header?.maximise"
        :backgroundColor="config.get?.window?.header?.background?.color"
        :color="config.get?.window?.header?.text?.color" />
    </QHeader>
    <QPageContainer>
      <RouterView v-slot="{ Component }">
        <transition appear enter-active-class="animated fadeIn" leave-active-class="animated fadeOut">
          <component :is="Component" />
        </transition>
      </RouterView>
    </QPageContainer>
  </QLayout>
</template>

<script lang="ts" setup>
import * as logger from '$app/logger/Logger'
import { computed } from 'vue'
import { RouterView } from 'vue-router'
import { useSystemInfo } from './store/systemInfo'
import { useConfigProvider } from './store/configProvider'
import { QLayout, QHeader, QPageContainer } from 'quasar'
import WindowTitle from '@/components/molecules/WindowTitle/WindowTitle.vue'

logger.Info('front-end loaded');

const systemInfo = useSystemInfo()
const config = useConfigProvider()

// Disable right clicking
document.addEventListener('contextmenu', e => {
  e.stopPropagation()
  e.preventDefault()
})
document.addEventListener('mousedown', e => {
  if (e.button === 2) {
    e.stopPropagation()
    e.preventDefault()
  }
})

const fancyWindow = computed(() => ['darwin', 'windows'].includes(systemInfo.platform))
</script>

<style lang="scss">
#app {
  height: 100vh;
  user-select: none;
  -webkit-user-select: none;
}
</style>
