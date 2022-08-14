<template>
  <header>
    <div class="header-content">
      <h1 class="content-logo">
        <img src="@/assets/logo.svg" @click="gotoHome" />
      </h1>
      <div class="bell-user">
        <div class="content-BellFilled">
          <el-icon :size="25" color="#8A919F"><BellFilled /></el-icon>
        </div>
        <router-link to="/userinfo"
          ><el-avatar id="avatar"> user </el-avatar></router-link
        >
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { BellFilled } from "@element-plus/icons-vue";
// Router
import { useRouter } from "vue-router";
const router = useRouter();
function gotoHome() {
  router.push("/");
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
  margin: 20px 900px 20px 0px;
  cursor: pointer;
}

.content-logo img {
  height: 100%;
}

.bell-user {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 80px;
}
</style>
