// Vue Core
import App from './App.vue'
import { createApp } from 'vue'
// Plugins
import pluginI18n from './plugins/i18n'
import pluginLogger from './plugins/logger'
import pluginPinia from './plugins/pinia'
import pluginRouter from './plugins/router'
import pluginVuetify from './plugins/vuetify'

// Async app installation context.
// Some plugins requires async actions to be done before install.
(async () => {
    createApp(App)
        .use(await pluginLogger())
        .use(await pluginI18n())
        .use(await pluginPinia())
        .use(await pluginRouter())
        .use(await pluginVuetify())
        .mount('#app')
})()
