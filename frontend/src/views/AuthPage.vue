<template>
    AuthPage
    <VBtn @click="router.push({ name: Page.Welcome })">welcome</VBtn>
</template>

<script lang="ts" setup async>
import { useRouter } from 'vue-router'
import { Page } from '@/plugins/router'

const router = useRouter()
</script>

<style lang="scss"></style>@/plugins/router