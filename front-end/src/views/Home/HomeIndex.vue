<template>
  <CommHeader />
  <SubHeader />
  <div class="common-layout">
    <el-container class="container">
      <el-main class="main">
        <ul
          infinite-scroll-immediate="false"
          infinite-scroll-distance="400"
          v-infinite-scroll="load"
          class="list-wrap"
          style="overflow: auto"
        >
          <!-- class="item" -->
          <ListItem
            v-for="item in articleList"
            :key="item.post_id"
            :articleListItem="item"
          />
        </ul>
      </el-main>
      <el-aside width="300px" class="aside">
        <div class="aside-wrap">Aside</div>
      </el-aside>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import CommHeader from "@/components/common/CommHeader.vue";
import SubHeader from "@/components/common/SubHeader.vue";
import { GetPostListWithTime } from "@/utils/api/article";
import { ref, onBeforeMount } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import ListItem from "./ListItem.vue";
const router = useRouter();

function gotoDetails(id: any) {
  router.push({ path: "/details", query: { id: id } });
}

const articleList = ref();
const nextTime = ref();
// 监听页面滚动到底，是时候加载数据了
const load = () => {
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
};

onBeforeMount(() => {
  // 获取当前时间戳
  nextTime.value = new Date().getTime();
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
    articleList.value = res.data;
    nextTime.value = res.next_time;
  });
});
</script>

<style scoped lang="less">
@normal_text_color: #86909c;

.common-layout {
  margin-top: 120px !important;
  .container {
    width: 1100px;
    .main {
      padding: 10px;
      .list-wrap {
        background: rgb(255, 255, 255);
        height: 600px;
        .item {
          height: 125px;
          border-bottom: solid 1px #dcdfe6;
          padding: 10px;
          box-sizing: border-box;
          cursor: pointer;
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

        .item:hover {
          background-color: #fafafa;
        }
      }
    }
    .aside {
      padding: 10px;

      .aside-wrap {
        padding: 10px;
        background: rgb(255, 255, 255);
      }
    }
  }
}
</style>
