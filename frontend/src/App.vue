<template>
  <VApp id="window" theme="dark">
    <VLayout>
      <Suspense>
        <WindowTitle />
      </Suspense>
      <VMain>
        <RouterView v-slot="{ Component }">
          <Transition name="page">
            <Suspense>
              <KeepAlive>
                <component class="page" :is="Component" />
              </KeepAlive>
            </Suspense>
          </Transition>
        </RouterView>
      </VMain>
    </VLayout>
  </VApp>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { RouterView, useRouter } from 'vue-router'
import { Page } from './plugins/router'
import WindowTitle from '@/components/organisms/WindowTitle/WindowTitle.vue'

const router = useRouter()

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

// Launcher Shortcuts
document.addEventListener('keydown', e => {
  console.log(e)
  if (e.metaKey || e.ctrlKey) {
    if (e.code === 'KeyR') {
      location.reload()
    }
    if (e.code === 'KeyI') {
      router.push({ name: Page.DevTools })
    }
  }
})
</script>

<style lang="scss" scoped async>
#window {
  width: 100vw;
  height: 100vh;
  border-radius: 8px;
  transition: .3s ease-in-out;
  border: 1px solid rgba(255, 255, 255, .1);
}

.page {
  position: absolute;
}

.page-enter-active,
.page-leave-active {
  transition: .3s ease-in-out;
}

.page-enter-from,
.page-leave-to {
  filter: opacity(0);
}

.page-enter-to,
.page-leave-from {
  filter: opacity(1);
}
</style>
