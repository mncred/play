import { createRouter, createWebHashHistory } from 'vue-router'

import AuthPage from '@/views/AuthPage.vue'
import WelcomePage from '@/views/WelcomePage.vue'

export enum Page {
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
                path: '/auth',
                name: Page.Auth,
                component: AuthPage,
            },
        ]
    })

    return router
}

export default pluginRouter