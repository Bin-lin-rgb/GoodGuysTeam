<template>
  <CommHeader />
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header">
        <div class="go-back-header">
          <el-icon id="arrow-left"><ArrowLeftBold /></el-icon>
          <div class="go-back-item" @click="GoBack">返回个人主页</div>
        </div>
      </el-header>
      <el-container class="container-aside-main">
        <el-aside width="250px" class="aside">
          <div class="aside-wrap">
            <el-menu
              default-active="1"
              class="el-menu-vertical-demo"
              @open="handleOpen"
              @close="handleClose"
            >
              <el-menu-item index="1" @click="GotoProfile">
                <el-icon><User /></el-icon>
                <span>个人资料</span>
              </el-menu-item>
              <el-menu-item index="2" @click="GotoAccount">
                <el-icon><Setting /></el-icon>
                <span>账号管理</span>
              </el-menu-item>
              <el-menu-item index="3">
                <el-icon><Collection /></el-icon>
                <span>简历管理</span>
              </el-menu-item>
              <el-menu-item index="4"
                ><el-icon><Warning /></el-icon>
                <span>屏蔽管理</span>
              </el-menu-item>
            </el-menu>
          </div>
        </el-aside>
        <el-main class="main"> <router-view></router-view> </el-main
      ></el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import CommHeader from "../../components/common/CommHeader.vue";
import {
  User,
  Collection,
  Setting,
  Warning,
  ArrowLeftBold,
} from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
const router = useRouter();

function GoBack() {
  router.push("/userinfo");
}
function GotoProfile() {
  router.push("/userinfo/setting/profile");
}
function GotoAccount() {
  router.push("/userinfo/setting/account");
}
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
</script>

<style lang="less" scoped>
.common-layout {
  margin-top: 80px !important;
  .container {
    width: 1280px;
    .header {
      margin-top: 10px;
      height: 40px;
      background: rgb(255, 255, 255);
      .go-back-header {
        height: 100%;
        width: 10%;
        display: flex;
        .go-back-item {
          text-align: center;
          line-height: 40px;
        }

        #arrow-left {
          margin: 13px;
        }
      }
      .go-back-header:hover {
        text-decoration: none;
        color: rgb(64, 158, 255);
        transition: 0.4s;
        cursor: pointer;
      }
    }
    .container-aside-main {
      padding-top: 20px;
      .aside {
        padding-right: 20px;
        .aside-wrap {
          padding: 10px;
          height: 600px;
          background: rgb(255, 255, 255);
        }
      }
      .main {
        width: 60%;
        padding: 20px;
        background: rgb(255, 255, 255);
      }
    }
  }
}
</style>
