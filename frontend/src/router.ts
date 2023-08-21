import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/', name: 'welcome', component: () => import('./views/WelcomePage.vue') },
        { path: '/auth', name: 'auth', component: () => import('./views/AuthPage.vue') },
    ]
})

export default router