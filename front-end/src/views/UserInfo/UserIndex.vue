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
          <a-tabs default-active-key="2">
            <a-tab-pane key="1" title="动态"> 动态 </a-tab-pane>
            <a-tab-pane key="2" title="文章" @click="GotoArticleList">
              <router-view></router-view>
            </a-tab-pane>
            <a-tab-pane key="3" title="专栏">
              Content of Tab Panel 2
            </a-tab-pane>
            <a-tab-pane key="4" title="沸点">
              Content of Tab Panel 2
            </a-tab-pane>
            <a-tab-pane key="5" title="收藏集">
              Content of Tab Panel 2
            </a-tab-pane>
            <a-tab-pane key="6" title="关注">
              Content of Tab Panel 2
            </a-tab-pane>
            <a-tab-pane key="7" title="赞"> Content of Tab Panel 2 </a-tab-pane>
          </a-tabs>
        </div>
      </el-main>
      <el-aside width="250px" class="aside">
        <div class="aside-wrap">
          <div class="star-block">
            <div class="block-title">个人成就</div>
            <a-divider />
            <div class="block-body">
              <div class="star-item">
                <el-icon><StarFilled /></el-icon>
                <span class="item-text">文章被点赞 2</span>
              </div>
              <div class="star-item">
                <el-icon><DocumentChecked /></el-icon>
                <span class="item-text">文章被阅读 726</span>
              </div>
              <div class="star-item">
                <el-icon><DataLine /></el-icon>
                <span class="item-text">掘力值 199</span>
              </div>
            </div>
          </div>
          <div class="follow-block"></div>
          <a-divider />
          <div class="more-block">
            <div class="more-item">
              <div class="item-left">收藏集</div>
              <span>3</span>
            </div>
            <a-divider />
            <div class="more-item">
              <div class="item-left">关注标签</div>
              <span>21</span>
            </div>
            <a-divider />
            <div class="more-item">
              <div class="item-left">加入于</div>
              <div>{{ userinfo.CreatedAt }}</div>
            </div>
          </div>
        </div>
      </el-aside>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, reactive, toRefs, onBeforeMount } from "vue";
import {
  ArrowRight,
  StarFilled,
  DataLine,
  DocumentChecked,
} from "@element-plus/icons-vue";
import { getUserInfo } from "@/utils/api/login";
const router = useRouter();

const state = reactive({
  circleUrl:
    "https://p6-passport.byteacctimg.com/img/mosaic-legacy/3791/5070639578~300x300.image",
});

const { circleUrl } = toRefs(state);

function GotoEdit() {
  router.push("/userinfo/setting/profile");
}

function GotoArticleList() {
  router.push("/userinfo/list");
}

const userinfo = ref({
  username: "",
  CreatedAt: "",
});

const loadUserInfo = async () => {
  const res: any = await getUserInfo();
  if (res.code === 1000) {
    userinfo.value = res.data;
    userinfo.value.CreatedAt = res.created_at;
  } else {
    return;
  }
};

onBeforeMount(() => {
  loadUserInfo();
  GotoArticleList();
});
</script>

<style scoped lang="less">
.common-layout {
  margin-top: 90px !important;

  .container {
    width: 1100px;
    .main {
      height: 800px;
      padding: 0 0 20px 20px;
      .major-area {
        width: 100%;
        background: rgb(255, 255, 255);
        padding-bottom: 12px;
        .user-info-block {
          width: 100%;
          padding: 25px 25px 25px 35px;
          display: flex;
          .avatar {
            text-align: center;
            line-height: 100%;
          }
          .info-box {
            width: 307.21px;
            position: relative;
            left: 5%;
            justify-content: flex-start;
            .info-username {
              font-size: large;
            }
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
      }
      .list-block {
        margin-top: 10px;
        background: rgb(255, 255, 255);
      }
    }
    .aside {
      padding: 0 20px 20px 10px;
      .aside-wrap {
        width: 100%;
        height: 100%;
        .star-block {
          background: rgb(255, 255, 255);
          padding: 10px;
          margin-bottom: 15px;
          .block-body {
            font-size: large;
          }
          .block-body {
            font-size: medium;
            .star-item {
              .item-text {
                margin-left: 5px;
              }
            }
          }
        }
        .follow-block {
          background: rgb(255, 255, 255);
          padding: 10px;
        }
        .more-block {
          padding: 10px;
          .more-item {
            width: 100%;
            display: inline-flex;
            .item-left {
              margin-right: auto;
              font-size: medium;
            }
          }
        }
      }
    }
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
