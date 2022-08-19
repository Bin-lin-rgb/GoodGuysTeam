<template>
  <div class="content">
    <el-result
      icon="success"
      title="Success Tip"
      sub-title="Please follow the instructions"
    >
      <template #extra>
        <el-button type="primary" @click="gotoCreator">Back</el-button>
      </template>
    </el-result>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();
function gotoCreator() {
  router.push("/creator");
}
</script>

<style lang="less" scoped>
.content {
  text-align: center;
  justify-content: center;
}
</style>
