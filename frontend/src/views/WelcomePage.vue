<template>
    <VContainer>
        <VBtn @click="router.push({ name: Page.Auth })">AUTH</VBtn>
    </VContainer>
</template>

<script lang="ts" setup>
import { useSystemInfo } from '@/store/systemInfo'
import { useRouter } from 'vue-router'
import { Page } from '@/plugins/router'
const systemInfo = useSystemInfo()

const router = useRouter()
</script>

<style lang="scss"></style>@/plugins/router