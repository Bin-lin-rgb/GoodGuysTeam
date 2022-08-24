<template>
  <CommHeader />
  <h1>testData</h1>
  <a-space>
    <a-button type="primary">Primary</a-button>
    <a-button>Secondary</a-button>
    <a-button type="dashed">Dashed</a-button>
    <a-button type="outline">Outline</a-button>
    <a-button type="text">Text</a-button>
  </a-space>
</template>

<script setup lang="ts">
import CommHeader from "@/components/common/CommHeader.vue";
// api
import { getUserInfo } from "@/utils/api/login";
//element
import { ElMessage } from "element-plus";
// import { ref } from "vue";

getUserInfo().then((res: any) => {
  console.log(res);
  if (res.code == 1000) {
    ElMessage({
      message: `Hello,${res.data.username}`,
      type: "success",
    });
  }
});
</script>

<style scoped>
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>
