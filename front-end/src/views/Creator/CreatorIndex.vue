<template>
  <SimpleHeader />
  <div class="common-layout">
    <el-container class="container">
      <el-aside width="260px" class="aside"
        ><div class="aside-wrap">
          <div class="user-buttom-bar">
            <div class="avatar">
              <el-avatar
                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
              />
            </div>
            <div class="userinfo-box">
              <div class="user-name">
                <!-- {{ data.author_name }} -->
                {{ userStore.username }}
              </div>
            </div>
            <el-divider />
            <div class="buttom-write">
              <button @click="GotoWrite">
                <div class="svg-wrapper-1">
                  <div class="svg-wrapper">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      width="24"
                      height="24"
                    >
                      <path fill="none" d="M0 0h24v24H0z"></path>
                      <path
                        fill="currentColor"
                        d="M1.946 9.315c-.522-.174-.527-.455.01-.634l19.087-6.362c.529-.176.832.12.684.638l-5.454 19.086c-.15.529-.455.547-.679.045L12 14l6-8-8 6-8.054-2.685z"
                      ></path>
                    </svg>
                  </div>
                </div>
                <span>写文章</span>
              </button>
            </div>
          </div>

          <el-menu
            default-active="2-1"
            class="el-menu-vertical-demo aside-menu"
            @open="handleOpen"
            @close="handleClose"
            @select="handleSelect"
          >
            <el-menu-item index="1" @click="GotoIndex">
              <el-icon><House /></el-icon>
              <span>首页</span>
            </el-menu-item>
            <el-sub-menu index="2">
              <template #title>
                <el-icon><Collection /></el-icon>
                <span>内容管理</span>
              </template>
              <el-menu-item index="2-1" @click="GotoArticleList">文章管理</el-menu-item>
              <el-menu-item index="2-2">专栏管理</el-menu-item>
              <el-menu-item index="2-3">沸点管理</el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="3">
              <template #title>
                <el-icon><Histogram /></el-icon>
                <span>数据中心</span>
              </template>
              <el-menu-item index="3-1">内容数据</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </div>
      </el-aside>
      <el-main class="main">
        <div class="main-warp">
          <router-view></router-view>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import SimpleHeader from "@/components/common/SimpleHeader.vue";
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { Histogram, Collection, House } from "@element-plus/icons-vue";

import { useUserStore } from "@/stores/user";
const userStore = useUserStore();

const router = useRouter();
function GotoWrite() {
  router.push("/write");
}
function GotoArticleList() {
  router.push("/creator/list");
}
function GotoIndex() {
  router.push("/creator/index");
}

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

onMounted(() => {
  GotoArticleList();
});
</script>

<style scoped lang="less">
.common-layout {
  // margin-top: 80px !important;
  position: fixed;
  top: 85px;
  left: 10%;
  .container {
    width: 1280px;
    .aside {
      padding: 10px;
      height: 600px;
      .aside-wrap {
        padding: 10px;
        background: rgb(255, 255, 255);
        .aside-menu {
          height: 100%;
        }
        .user-buttom-bar {
          margin: 0 auto;
          text-align: center;
          .userinfo-box {
            color: #515767;
          }
          .avatar {
            margin-top: 20px;
          }
          .buttom-write {
            margin: 5px 46px;
          }
        }
      }
    }
    .main {
      padding: 10px;
      .main-warp {
        background: rgb(255, 255, 255);
      }
    }
  }
}

/* From uiverse.io by @adamgiebl */
button {
  font-family: inherit;
  font-size: 20px;
  background: royalblue;
  color: white;
  padding: 0.7em 1em;
  padding-left: 0.9em;
  display: flex;
  align-items: center;
  border: none;
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.2s;
}

button span {
  display: block;
  margin-left: 0.3em;
  transition: all 0.3s ease-in-out;
}

button svg {
  display: block;
  transform-origin: center center;
  transition: transform 0.3s ease-in-out;
}

button:hover .svg-wrapper {
  animation: fly-1 0.6s ease-in-out infinite alternate;
}

button:hover svg {
  transform: translateX(1.2em) rotate(45deg) scale(1.1);
}

button:hover span {
  transform: translateX(5em);
}

button:active {
  transform: scale(0.95);
}

@keyframes fly-1 {
  from {
    transform: translateY(0.1em);
  }

  to {
    transform: translateY(-0.1em);
  }
}

/* 网页滑动条样式 */
::-webkit-scrollbar {
  width: 5px;
  height: 5px;
}

::-webkit-scrollbar-thumb {
  border-radius: 7px;
  background-color: var(--Yuexing-color);
  transition: all 0.5s;
}

::-webkit-scrollbar-track {
  /* border-radius: 7px; */
  background-color: #d3dce6;
}
</style>
