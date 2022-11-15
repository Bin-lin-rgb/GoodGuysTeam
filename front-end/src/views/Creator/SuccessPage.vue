<template>
  <div class="content">
    <el-result
      icon="success"
      title="发布成功"
      sub-title="继续学习，请保持每天一定的知识输出哦！"
    >
      <template #extra>
        <el-button type="primary" @click="gotoCreator">返回</el-button>
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
