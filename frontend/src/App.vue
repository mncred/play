<template>
  <VApp id="window" theme="dark">
    <VLayout>
      <Suspense>
        <!-- ! TODO: It isn't correct to bind to COMPILER OS, need fetch OS of CURRENT machine -->
        <WindowTitle draggable :appearance="uaparser.getOS().name === 'Mac OS' ? 'macos' : 'common'" :icon="mdiCube"
          iconColor="red" />
      </Suspense>
      <VMain>
        <RouterView v-slot="{ Component }">
          <!-- FIXME: Transition doesn't work -->
          <Transition name="page">
            <KeepAlive>
              <Suspense>
                <div>
                  <component :is="Component" />
                </div>
              </Suspense>
            </KeepAlive>
          </Transition>
        </RouterView>
      </VMain>
    </VLayout>
  </VApp>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { UAParser } from 'ua-parser-js'
import { RouterView } from 'vue-router'
import { mdiCube } from '@mdi/js'
import WindowTitle from '@/components/organisms/WindowTitle/WindowTitle.vue'

const uaparser = new UAParser()

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

const fancyWindow = computed(() => ['Mac OS', 'Windows'].includes(uaparser.getOS().name ?? ''))
const borderRadius = computed(() => {
  if (!fancyWindow.value) return '0px'
  return `8px`
})
</script>

<style lang="scss" scoped async>
#window {
  overflow: hidden;
  width: 100vw;
  height: 100vh;
  border-radius: v-bind('borderRadius');
  transition: .3s ease-in-out;
  border: 1px solid rgba(255, 255, 255, .1);
}

.page-enter-active,
.page-leave-active {
  transition: opacity 0.5s ease;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>

<style lang="scss">
:root {
  overflow: hidden;
  margin: 0;
  padding: 0;
  user-select: none;
  -webkit-user-select: none;
  cursor: default;
}
</style>
