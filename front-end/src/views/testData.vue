<template>
  <CommHeader />
  <h1>testData</h1>
</template>

<script setup lang="ts">
import CommHeader from "@/components/common/CommHeader.vue";
// api
import { getUserInfo } from "@/utils/api/login";
//element
import { ElMessage } from "element-plus";
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

<style scoped></style>
