<script setup lang="ts">
import { RouterView } from "vue-router";
import CommHeader from "./components/common/CommHeader.vue";
</script>

<template>
  <CommHeader />
  <RouterView />
</template>

<style>
@import "./assets/css/reset.css";
</style>
