import '@quasar/extras/roboto-font/roboto-font.css'
import '@quasar/extras/mdi-v7/mdi-v7.css'
import 'quasar/src/css/index.sass'
// Animations
import '@quasar/extras/animate/fadeIn.css'
import '@quasar/extras/animate/fadeOut.css'

import App from './App.vue'
import { Quasar } from 'quasar'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import i18n from './assets/locales'
import router from './router'

createApp(App)
    .use(i18n)
    .use(createPinia())
    .use(router)
    .use(Quasar, {})
    .mount('#app')
