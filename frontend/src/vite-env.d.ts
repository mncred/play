/// <reference types="vite/client" />
/// <reference types="@modyfi/vite-plugin-yaml/modules" />

declare module '*.vue' {
    import type { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}
