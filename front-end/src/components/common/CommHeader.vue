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
          <!-- <el-button type="primary" @click="gotoCreator">创作者中心</el-button> -->
        </div>
        <div class="content-BellFilled">
          <el-icon :size="25" color="#8A919F"><BellFilled /></el-icon>
        </div>
        <div class="content-login">
          <router-link to="/login">登录 / 注册</router-link>
        </div>
        <router-link to="/userinfo" id="avatar"
          ><el-avatar> user </el-avatar></router-link
        >
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
