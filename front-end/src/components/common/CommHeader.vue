<template>
  <header>
    <div class="header-content">
      <h1 class="content-logo">
        <img src="@/assets/logo.svg" @click="gotoHome" />
      </h1>
      <div class="content-nav">
        <ul>
          <router-link class="link" to="/">首页</router-link>
          <li>课程</li>
          <li>直播</li>
          <li>活动</li>
        </ul>
      </div>
      <div class="search-login">
        <div class="content-search">
          <input type="" placeholder="请输入要搜索的内容" />
          <el-icon color="#808080" :size="22">
            <search style="width: 22px; height: 22px" />
          </el-icon>
        </div>
        <div class="content-create">
          <button @click="gotoCreator">创作者中心</button>
        </div>
        <div class="content-BellFilled">
          <el-icon :size="25" color="#8A919F"><BellFilled /></el-icon>
        </div>
        <div class="content-login" v-if="!isLogin">
          <router-link to="/login">登录 / 注册</router-link>
        </div>
        <!-- 划过头像显示  -->
        <el-dropdown v-else>
          <span class="el-dropdown-link">
            <el-avatar
              @mouseenter="isShow = true"
              src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          /></span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="gotoWrite">写文章</el-dropdown-item>
              <el-dropdown-item @click="gotoUserinfo"
                >我的主页</el-dropdown-item
              >
              <el-dropdown-item>设置</el-dropdown-item>
              <el-dropdown-item @click="Logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { Search, BellFilled } from "@element-plus/icons-vue";
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

function gotoHome() {
  router.push("/");
}
function gotoCreator() {
  router.push("/creator");
}
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

function Logout() {
  ElMessageBox.confirm("确定退出登录吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      const userStore = useUserStore();
      userStore.clearToken();
      router.go(0);
      ElMessage({
        type: "success",
        message: "退出成功！",
      });
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消",
      });
    });
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
</script>

<style scoped>
header {
  position: fixed;
  top: 0;
  left: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 70px;
  background: white;
  box-shadow: 0px 5px 6px rgba(0, 0, 0, 0.16);
  opacity: 1;
}

.header-content {
  display: flex;
  justify-content: space-around;
  width: 1700px;
}

.content-logo {
  width: 100px;
  height: 30px;
  margin: 20px 0;
  cursor: pointer;
}

.content-logo img {
  height: 100%;
}

.content-nav {
  width: 300px;
  height: 75px;
}

.content-nav ul {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
  height: 70px;
}

.content-nav ul li {
  font-size: 16px;
  color: #323232;
}

.content-nav ul .link {
  font-size: 16px;
  color: #323232;
}

.search-login {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 650px;
}

.content-search {
  display: flex;
  align-items: center;
  padding: 5px 10px;
  width: 350px;
  height: 35px;
  border-radius: 8px;
  background: #f0f2f4;
}

.content-search input {
  padding: 0 10px;
  width: 430px;
  height: 40px;
  border: 0;
  border-radius: 8px;
  color: #808080;
  background: #f0f2f4;
  font-size: 16px;
  outline: none;
}
.content-create {
  /* background: #1e80ff; */
  height: 38px;
  margin: auto;
  width: 105px;
  /* color: #ffffff; */
}

.content-BellFilled {
  /* font-size: 23px; */
  margin-right: 23px;
}

/* From uiverse.io by @kirzin */
button {
  text-decoration: none;
  position: absolute;
  border: none;
  font-size: 14px;
  font-family: inherit;
  color: #fff;
  width: 8em;
  height: 2.2em;
  line-height: 2em;
  text-align: center;
  background: linear-gradient(90deg, #03a9f4, #f441a5, #ffeb3b, #03a9f4);
  background-size: 300%;
  border-radius: 30px;
  z-index: 1;
}

button:hover {
  animation: ani 8s linear infinite;
  border: none;
}

@keyframes ani {
  0% {
    background-position: 0%;
  }

  100% {
    background-position: 400%;
  }
}

button:before {
  content: "";
  position: absolute;
  top: -5px;
  left: -5px;
  right: -5px;
  bottom: -5px;
  z-index: -1;
  background: linear-gradient(90deg, #03a9f4, #f441a5, #ffeb3b, #03a9f4);
  background-size: 400%;
  border-radius: 35px;
  transition: 1s;
}

button:hover::before {
  filter: blur(20px);
}

button:active {
  background: linear-gradient(32deg, #03a9f4, #f441a5, #ffeb3b, #03a9f4);
}

.content-login {
  font-size: 18px;
  color: #808080;
  text-align: center;
  cursor: pointer;
}

#avatar {
  position: fixed;
  right: 20px;
}
</style>
