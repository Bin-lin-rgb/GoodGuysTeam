<template>
  <CommHeader />
  <h1>testData</h1>
  <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
    <li v-for="i in count" :key="i" class="infinite-list-item">{{ i }}</li>
  </ul>
</template>

<script setup lang="ts">
import CommHeader from "@/components/common/CommHeader.vue";
// api
import { getUserInfo } from "@/utils/api/login";
//element
import { ElMessage } from "element-plus";
import { ref } from "vue";
const count = ref(0);
const load = () => {
  count.value += 2;
};
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
