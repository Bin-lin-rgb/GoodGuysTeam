<template>
  <header>
    <div class="header-content">
      <h1 class="content-logo">
        <img src="@/assets/logo.svg" @click="gotoHome"/>
      </h1>
      <div class="content-nav">
        <ul>
          <router-link class="link" to="/">首页</router-link>
          <li>课程</li>
          <li>直播</li>
          <li>活动</li>
          <li>沸点</li>
        </ul>
      </div>
      <div class="search-buy-login">
        <div class="content-search">
          <input type="" placeholder="请输入要搜索的课程" />
          <el-icon color="#808080" :size="22">
            <search style="width: 22px; height: 22px" />
          </el-icon>
        </div>
        <div class="content-create">
          <el-button type="primary" @click="gotoCreator">创作者中心</el-button>
        </div>
        <div class="content-BellFilled">
          <el-icon :size="25" color="#8A919F"><BellFilled /></el-icon>
        </div>
        <div class="content-login">
          <router-link to="/login">登录 / 注册</router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { Search, BellFilled } from "@element-plus/icons-vue";
// Router
import { useRouter } from "vue-router";
const router = useRouter();
function gotoHome() {
  router.push("/");
}
function gotoCreator() {
  router.push("/creator");
}
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
  width: 110px;
  height: 30px;
  margin: 20px 0;
  cursor: pointer;
}

.content-logo img {
  height: 100%;
}

.content-nav {
  width: 400px;
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

.search-buy-login {
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
.create-button {
  background: #1e80ff;
  height: 40px;
  color: #ffffff;
}

/* .content-BellFilled {
  font-size: 23px;
} */

.content-login {
  font-size: 18px;
  color: #808080;
  text-align: center;
  cursor: pointer;
}
</style>
