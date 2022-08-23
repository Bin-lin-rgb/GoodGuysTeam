<template>
  <div class="w">
    <CommHeader />
    <div class="common-layout">
      <el-container class="container">
        <el-main class="main">
          <div class="main-wrap">
            <h1 class="title">title</h1>
            <div class="userinfo-bar">
              <div class="avatar">
                <el-avatar
                  src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                />
              </div>
              <div class="userinfo-box">
                <div class="user-name">{{ data.author_name }}</div>
                <div class="created-at">{{ data.created_at2 }}</div>
              </div>
            </div>
            <md-editor
              v-model="text"
              :previewOnly="true"
              class="md"
              @onGetCatalog="GetCatalog"
            />
          </div>
        </el-main>
        <el-aside width="300px" class="aside">
          <div class="aside-wrap">
            <div class="aside-user-bar">
              <div class="aside-avatar">
                <el-avatar
                  src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                />
              </div>
              <div class="aside-user-name">{{ data.author_name }}</div>
            </div>
            <el-affix :offset="120">
              <el-button type="primary">Offset top 120px</el-button>
              <div class="aside-catalog"></div>
            </el-affix>
          </div>
        </el-aside>
      </el-container>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CommHeader from "@/components/common/CommHeader.vue";
import { ref, onBeforeMount, reactive } from "vue";
import { GetPostDetails } from "@/utils/api/article";
// md-editor-v3
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
// router
import { useRoute } from "vue-router";
const route = useRoute();

const text = ref("Hello Editor!");
let data = reactive({
  author_name: "ddff",
  community_name: "Android",
  content: "content test content",
  created_at2: "2022-08-21 22:15",
  title: "test",
});

const loadBlog = async () => {
  const res = await GetPostDetails({ id: route.query.id });
  data = res.data;
  text.value = res.data.content;
  console.log(data);
};

onBeforeMount(() => {
  loadBlog();
});

const GetCatalog = (h: string) => {
  console.log(h);
  // 0: {text: '上手使用', level: 1}
  // 1: {text: '生产环境 Tree Sharking', level: 2}
  // length: 2
};
</script>

<style scoped lang="less">
.w {
  .common-layout {
    margin-top: 90px !important;
    .container {
      width: 1200px;
      .main {
        padding: 10px;
        .main-wrap {
          padding: 32px;
          background: rgb(255, 255, 255);
          .title {
            font-size: 3rem;
            font-weight: 700;
            margin-bottom: 1rem;
          }
          .userinfo-bar {
            display: flex;
            box-sizing: border-box;
            height: 50px;
            margin: 10px 10px 10px 0;
            .avatar {
              margin-right: 10px;
            }
            .userinfo-box {
              min-height: 40px;
              color: #515767;
              .created-at {
                color: #818797;
              }
            }
          }

          // .md {}
        }
      }
      .aside {
        padding: 10px;

        .aside-wrap {
          padding: 10px;
          background: rgb(255, 255, 255);
          .aside-user-bar {
            margin: 0 auto;
            text-align: center;
            .aside-avatar {
              // width: 100%;
              margin-top: 2rem;
            }

            .aside-user-name {
              margin-bottom: 2rem;
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
