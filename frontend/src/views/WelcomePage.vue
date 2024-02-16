<template>
    WelcomePage
    <VBtn @click="router.push({ name: Page.Auth })">AUTH</VBtn>
</template>

<script lang="ts" setup async>
import { useSystemInfo } from '@/store/systemInfo'
import { useRouter } from 'vue-router'
import { Page } from '@/plugins/router'
const systemInfo = useSystemInfo()

const router = useRouter()
</script>

<style lang="scss"></style>@/plugins/router