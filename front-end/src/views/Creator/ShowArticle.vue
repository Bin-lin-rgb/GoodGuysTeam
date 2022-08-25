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
  <div class="common-layout">
    <el-container>
      <el-main>
        <div class="count-bar">
          <el-tabs
            type="border-card"
            v-model="activeName"
            class="tabs"
            @tab-click="handleClick"
          >
            <el-tab-pane label="全部" class="tab-all">
              <ListItem
                v-for="item in articleList"
                :key="item.id"
                :articleListItem="item"
              />
            </el-tab-pane>
            <el-tab-pane label="发布">发布</el-tab-pane>
            <el-tab-pane label="审核中（0）">审核中（0）</el-tab-pane>
            <el-tab-pane label="未通过（0）">未通过（0）</el-tab-pane>
          </el-tabs>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount } from "vue";
import ListItem from "@/views/Creator/ListItem.vue";
import type { TabsPaneContext } from "element-plus";
import { GetPostListByUserId } from "@/utils/api/article";

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
  articleList.value = res.data;
  for (let i = 0; i < articleList.value.length; i++) {
    articleList.value[i].id = i;
  }
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
    // height: 500px;
    .tab-all {
      overflow: auto;
    }
  }
}
</style>
