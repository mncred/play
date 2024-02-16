import ViteYaml from '@modyfi/vite-plugin-yaml'
import { defineConfig } from 'vite'
import path from "path"
import vue from '@vitejs/plugin-vue'

const __dirname = './'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ViteYaml(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, "./src/"),
      '$runtime': path.resolve(__dirname, "./wailsjs/runtime/"),
      '$app': path.resolve(__dirname, "./wailsjs/go/"),
    },
  }
})
