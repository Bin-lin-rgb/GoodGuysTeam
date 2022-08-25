<template>
  <div class="info-main">
    <div class="main-title">{{ articleListItem.title }}</div>
    <div class="main-info">{{ articleListItem.created_at2 }}</div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  articleListItem: any;
}>();
</script>

<style lang="less" scoped>
.info-main {
  height: 80px;
  padding: 10px;
  border-bottom: solid #e5e6eb 1px;
  .main-title {
    font-size: 22px;
  }
}
</style>
