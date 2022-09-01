<template>
  <CommHeader />
  <div class="common-layout">
    <el-container class="container">
      <el-main class="main">
        <div class="major-area">
          <div class="user-info-block">
            <el-avatar :size="90" :src="circleUrl" class="avatar" />
            <el-row class="info-box" type="flex">
              <el-col>
                <div class="info-username">{{ userinfo.username }}</div>
              </el-col>
            </el-row>
            <div class="action-box">
              <el-row>
                <el-col>
                  <div>
                    <el-button type="primary" plain @click="GotoEdit"
                      >编辑个人资料</el-button
                    >
                  </div>
                </el-col>
              </el-row>
            </div>
          </div>
          <el-divider />
          <div class="badge-wall">
            <div class="badge-left">
              <div>
                <el-row class="badge-left">
                  <div class="badge-img"></div>
                  <div>获得勋章</div>
                </el-row>
              </div>
            </div>
            <div class="badge-list">
              <el-icon :size="20">
                <ArrowRight />
              </el-icon>
            </div>
          </div>
        </div>
        <div class="list-block">
          <div class="list-header">
            <el-menu default-active="2" class="el-menu-demo" mode="horizontal">
              <el-menu-item index="1">动态</el-menu-item>
              <el-menu-item index="2" @click="GotoArticleList"
                >文章</el-menu-item
              >
              <el-menu-item index="3">专栏</el-menu-item>
              <el-menu-item index="4">沸点</el-menu-item>
              <el-menu-item index="5">收藏集</el-menu-item>
              <el-menu-item index="6">关注</el-menu-item>
              <el-menu-item index="7">赞</el-menu-item>
            </el-menu>
          </div>
          <div class="list-body">
            <!-- <vue-router></vue-router> -->
          </div>
        </div>
      </el-main>
      <el-aside width="250px" class="aside">
        <div class="aside-wrap"></div>
      </el-aside>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, reactive, toRefs, onBeforeMount } from "vue";
import { ArrowRight } from "@element-plus/icons-vue";
import { getUserInfo } from "@/utils/api/login";
const router = useRouter();

const state = reactive({
  circleUrl:
    "https://p6-passport.byteacctimg.com/img/mosaic-legacy/3791/5070639578~300x300.image",
});

const { circleUrl } = toRefs(state);

function GotoEdit() {
  router.push("/userinfo/setting");
}

function GotoArticleList() {
  router.push("/userinfo/list");
}

const userinfo = ref({
  username: "",
});
const loadUserInfo = async () => {
  const res: any = await getUserInfo();
  if (res.code === 1000) {
    userinfo.value = res.data;
  } else {
    return;
  }
};

onBeforeMount(() => {
  loadUserInfo();
});
</script>

<style scoped lang="less">
.common-layout {
  margin-top: 80px !important;

  .container {
    width: 1100px;
    .main {
      height: 800px;
      .major-area {
        width: 100%;
        background: rgb(255, 255, 255);
        .user-info-block {
          width: 100%;
          padding: 25px 25px 25px 35px;
          display: flex;
          .avatar {
            text-align: center;
            line-height: 100%;
            // margin: 0 auto;
          }
          .info-box {
            width: 307.21px;
            position: relative;
            left: 5%;
            justify-content: flex-start;
          }

          .action-box {
            width: 118px;
            position: relative;
            left: 25%;
            display: flex;
            align-items: flex-end;
          }
        }
        .badge-wall {
          width: 708px;
          position: relative;

          .badge-left {
            width: 106.13px;
            height: 28px;
            position: relative;
            margin-left: 20px;
            .badge-img {
              background-image: url("https://lf3-cdn-tos.bytescm.com/obj/static/xitu_juejin_web/img/badge-count-icon.0586ac4.png");
              background-size: contain;
              width: 20px;
              height: 20px;
              margin: 0 5px;
            }
          }

          .badge-list {
            width: 21px;
            height: 28px;
            position: absolute;
            top: 0px;
            right: 5%;
            justify-content: flex-end;
          }
        }
        .list-block {
          margin-top: 10px;
        }
      }
    }
    .aside {
      padding: 20px 20px 20px 0;
      .aside-wrap {
        width: 100%;
        height: 100%;
        background: rgb(255, 255, 255);
      }
    }
  }
}
</style>
