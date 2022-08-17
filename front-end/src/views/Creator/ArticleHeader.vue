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
    <div class="content-login" v-if="!isLogin">
      <router-link to="/login">登录 / 注册</router-link>
    </div>
    <!-- 划过头像显示  -->
    <el-dropdown v-else>
      <span class="el-dropdown-link">
        <el-avatar
          class="avatar"
          @mouseenter="isShow = true"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      /></span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item @click="gotoWrite">写文章</el-dropdown-item>
          <el-dropdown-item @click="gotoUserinfo">我的主页</el-dropdown-item>
          <el-dropdown-item>设置</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
// pinia
import { useUserStore } from "@/stores/user";
// Router
import { useRouter } from "vue-router";
// api
import { getUserInfo } from "@/utils/api/login";
//element
import { ElMessage, ElMessageBox } from "element-plus";

const router = useRouter();
//用户是否是登录状态
const isLogin = ref(false);
//显示用户更多数据
const isShow = ref(false);
const userInfo = ref();

function gotoUserinfo() {
  router.push("/userinfo");
}
function gotoWrite() {
  const routeData = router.resolve({
    // query: {id:deviceId},
    path: "/write",
  });
  window.open(routeData.href, "_blank");
}

onBeforeMount(() => {
  getUserInfo().then((res: any) => {
    // console.log(res);
    if (res.code == 1000) {
      isLogin.value = true;
      userInfo.value = res.data[0];
    }
  });
});

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
  .avatar {
    margin: 15px 45px 15px;
  }
}
</style>
