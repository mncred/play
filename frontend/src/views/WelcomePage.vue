<template>
    <QPage class="q-pa-md">
        <QBtn :to="{ name: 'auth' }">NEXT</QBtn>
        {{ systemInfo.platform }}
    </QPage>
</template>

<script lang="ts" setup>
import { QBtn, QPage } from 'quasar'
import { useSystemInfo } from '@/store/systemInfo'
const systemInfo = useSystemInfo()
</script>

<style lang="scss"></style>../store/cpuload