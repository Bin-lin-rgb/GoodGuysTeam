<template>
  <CommHeader />
  <SubHeader />
  <div class="common-layout">
    <el-container class="main">
      <el-main>
        <el-row :gutter="20">
          <el-col :span="16">
            <ul
              infinite-scroll-immediate="false"
              v-infinite-scroll="load"
              class="list-wrap"
              style="overflow: auto"
            >
              <li class="item" v-for="item in articleList" :key="item.post_id">
                <div class="info-bar">
                  {{ item.author_name }}
                  <div class="divider"></div>
                  {{ item.created_at }}
                  <div class="divider"></div>
                  {{ item.community_name }}
                </div>

                <div class="info-main">
                  <div class="main-title">{{ item.title }}</div>
                  <div class="main-context">{{ item.content }}</div>
                </div>
              </li>
            </ul>
          </el-col>
          <el-col :span="8"
            ><div class="grid-content aside"></div
          ></el-col> </el-row
      ></el-main>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import CommHeader from "@/components/common/CommHeader.vue";
import SubHeader from "@/components/common/SubHeader.vue";
import { GetPostList, GetPostListWithTime } from "@/utils/api/article";
import { ref, onBeforeMount } from "vue";
import { ElMessage } from "element-plus";
// import { useRouter } from "vue-router";
// const router = useRouter();

const articleList = ref();
const nextTime = ref();
// 监听页面滚动到底，是时候加载数据了
function load() {
  GetPostListWithTime({ nextTime: nextTime.value }).then((res: any) => {
    console.log(res);
    if (res.code != 1000) {
      ElMessage({
        message: res.message,
        type: "error",
      });
      return;
    }
    // message: `获取列表成功！`
    articleList.value = [...articleList.value, ...res.data];
    nextTime.value = res.data.next_time;
  });
}

onBeforeMount(() => {
  GetPostList().then((res: any) => {
    console.log(res);
    if (res.code != 1000) {
      ElMessage({
        message: res.message,
        type: "error",
      });
      return;
    }
    // message: `获取列表成功！`
    articleList.value = res.data;
    nextTime.value = res.next_time;
    console.log("onBeforeMount", nextTime.value);
  });
});
</script>

<style scoped lang="less">
@normal_text_color: #86909c;

.common-layout {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 110px;
  left: 0;
  background: #f4f5f5;
  .main {
    width: 1200px;
    position: fixed;
    top: 120px;
    left: 15%;
    .list-wrap {
      background: rgb(255, 255, 255);
      height: calc(85vh);
      .item {
        height: 125px;
        border-bottom: solid 1px #dcdfe6;
        padding: 10px;
        box-sizing: border-box;
        .info-bar {
          height: 30px;
          display: inline-block;
          color: @normal_text_color;
          .divider {
            width: 1px;
            height: 14px;
            background: #e5e6eb;
            margin: 0 8px;
            display: inline-block;
          }
        }
        .info-main {
          .main-title {
            font-weight: bold;
            font-size: 20px;
            margin-bottom: 12px;
            /* 文本过长自动截断 */
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
          }
          .main-context {
            color: @normal_text_color;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
          }
        }
      }
    }
    .aside {
      background: rgb(255, 255, 255);
    }
    .grid-content {
      border-radius: 4px;
      min-height: 36px;
      padding: 21px;
    }
  }
}

el-row {
  margin-bottom: 20px;
}

el-col {
  border-radius: 4px;
}
</style>
