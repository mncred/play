<template>
    <VContainer>
        <VBtn @click="router.push({ name: Page.Welcome })">welcome</VBtn>
    </VContainer>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { Page } from '@/plugins/router'

const router = useRouter()
</script>

<style lang="scss"></style>@/plugins/router