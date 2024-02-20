import { createRouter, createWebHashHistory } from 'vue-router'

import AuthPage from '@/views/AuthPage.vue'
import DevToolsPage from '@/views/DevToolsPage.vue'
import WelcomePage from '@/views/WelcomePage.vue'

export enum Page {
    DevTools = 'devtools',
    Welcome = 'welcome',
    Auth = 'auth',
}

const pluginRouter = async () => {
    const router = createRouter({
        history: createWebHashHistory(),
        routes: [
            {
                path: '/',
                name: Page.Welcome,
                component: WelcomePage,
            },
            {
                path: '/devtools',
                name: Page.DevTools,
                component: DevToolsPage,
            },
            {
                path: '/auth',
                name: Page.Auth,
                component: AuthPage,
            },
        ]
    })

    return router
}

export default pluginRouter