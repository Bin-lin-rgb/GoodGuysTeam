<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
  >
    <el-menu-item index="1">文章</el-menu-item>
    <el-menu-item index="2">草稿箱</el-menu-item>
  </el-menu>

  <div class="count-bar">
    <el-tabs
      type="border-card"
      v-model="activeName"
      class="tabs"
      @tab-click="handleClick"
    >
      <el-tab-pane label="全部" class="tab-all">
        <a-empty v-if="isEmpty" />
        <ListItem
          v-else
          v-for="item in articleList"
          :key="item.id"
          :articleListItem="item"
          @success-delete="SuccessDelete"
        />
      </el-tab-pane>
      <el-tab-pane label="发布">发布</el-tab-pane>
      <el-tab-pane label="审核中（0）">审核中（0）</el-tab-pane>
      <el-tab-pane label="未通过（0）">未通过（0）</el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount } from "vue";
import ListItem from "@/views/Creator/ListItem.vue";
import { ElMessage, type TabsPaneContext } from "element-plus";
import { GetPostListByUserId } from "@/utils/api/article";
// import { useRouter } from "vue-router";
// const router = useRouter();

const isEmpty = ref(true);

const activeName = ref("0");
const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

const activeIndex = ref("1");
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

const articleList = ref();
const loadBlog = async () => {
  const res = await GetPostListByUserId();
  if (res.data.length == 0) {
    isEmpty.value = true;
    return;
  }
  isEmpty.value = false;
  articleList.value = res.data;
  // console.log("object", articleList.value);
  for (let i = 0; i < articleList.value.length; i++) {
    articleList.value[i].id = i;
  }
};

const SuccessDelete = async () => {
  ElMessage({
    message: `删除成功！`,
    type: "success",
  });
  const res = await GetPostListByUserId();
  if (res.data.length == 0) {
    isEmpty.value = true;
    return;
  }
  isEmpty.value = false;
  articleList.value = res.data;
};

onBeforeMount(() => {
  loadBlog();
});
</script>

<style lang="less" scoped>
.count-bar {
  display: flex;
  .tabs {
    width: 100%;
    height: 90%;
    overflow: hidden;
    .tab-all {
      overflow: auto;
    }
  }
}
</style>
