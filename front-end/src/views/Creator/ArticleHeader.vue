<template>
  <div class="header">
    <el-input
      v-model="input"
      placeholder="输入文章标题..."
      clearable
      maxlength="30"
      show-word-limit
      type="text"
      class="w-50 m-2 input"
      size="large"
    />
    <span class="plainText">文章将自动保存至草稿箱</span>
    <el-button type="primary" plain>草稿箱</el-button>
    <el-button type="primary">发布</el-button>
    <router-link to="/userinfo"
      ><el-avatar id="avatar"> user </el-avatar></router-link
    >
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
const input = ref("");
</script>

<style scoped lang="less">
.header {
  height: 70px;
  width: 100%;
  position: fixed;
  top: 0px;
  left: 0;
  .input {
    margin: 10px 10px 10px;
    width: 60%;
  }
  .plainText {
    margin: 27px 30px 27px 70px;
    color: #afafaf;
    font-size: 16px;
  }
  #avatar {
    margin-left: 35px;
    margin-top: 15px;
  }
}
</style>
