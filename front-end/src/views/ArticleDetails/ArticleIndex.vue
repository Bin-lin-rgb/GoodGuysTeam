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
              :marked-heading-id="generateId"
              @onGetCatalog="GetCatalog"
            />
            <a-space
              >标签 :
              <a-tag checkable color="arcoblue" :default-checked="true">
                {{ data.community_name }}</a-tag
              >
            </a-space>
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
            <el-divider />
            <div class="catalog">
              <a-affix :offsetTop="90" class="a-affix">
                <a-anchor class="a-anchor">
                  <a-anchor-link
                    v-for="item in catalog"
                    :key="item.id"
                    :href="`#heading-${item.id}`"
                  >
                    {{ item.text }}
                    <!-- <template #sublist>
                    <a-anchor-link
                      v-for="item in level_2"
                      :key="item.id"
                      :href="`#heading-${item.id}`"
                      >{{ item.text }}
                      <template #sublist>
                        <a-anchor-link
                          v-for="item in level_2"
                          :key="item.id"
                          :href="`#heading-${item.id}`"
                          >{{ item.text }}</a-anchor-link
                        >
                      </template></a-anchor-link
                    >
                  </template> -->
                  </a-anchor-link>
                </a-anchor></a-affix
              >
            </div>
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

// 加载内容
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
};

onBeforeMount(() => {
  loadBlog();
});

// 目录相关
const catalog = ref();
// const level_1 = ref();
// const level_2 = ref();
// const level_3 = ref();
const generateId = (_text: any, _level: any, index: any) => `heading-${index}`;
const GetCatalog = (list: []) => {
  // list 一个对象数组
  // 0: {text: '上手使用', level: 1}
  // 1: {text: '生产环境 Tree Sharking', level: 2}
  // length: 2
  catalog.value = list;
  for (let index = 1; index <= catalog.value.length; index++) {
    // 增加id属性
    catalog.value[index - 1].id = index;
    // 根据level分级,只考虑到第三级
    // switch (catalog.value[index - 1].level) {
    //   case 1:
    //     level_1.value = [...level_1.value, ...catalog.value[index - 1]];
    //     console.log("---", level_1);
    //     break;
    //   case 2:
    //     level_2.value = [...level_2.value, ...catalog.value[index - 1]];
    //     break;
    //   case 3:
    //     level_3.value = [...level_3.value, ...catalog.value[index - 1]];
    //     break;
    //   default:
    //     break;
    // }
  }
};

MdEditor.config({
  markedRenderer(renderer) {
    renderer.heading = (text, level, _r, _s, index) => {
      const id = generateId(text, level, index);
      return `<h${level} id="${id}">${text}</h${level}>`;
    };
    return renderer;
  },
});
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

          .aside-user-bar {
            background: rgb(255, 255, 255);
            margin: 0 auto;
            text-align: center;
            .aside-avatar {
              padding-top: 2rem;
              padding-bottom: 2rem;
            }

            .aside-user-name {
              padding-bottom: 2rem;
            }
          }

          .catalog {
            .a-affix {
              .a-anchor {
                background: rgb(255, 255, 255);
                width: 250px;
              }
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
