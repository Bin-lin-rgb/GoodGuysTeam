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
  <div class="list-wrap">
    <div class="count-bar">
      <el-tabs
        type="border-card"
        v-model="activeName"
        class="tabs"
        @tab-click="handleClick"
      >
        <el-tab-pane label="全部">
          <ListItem
            v-for="item in articleList"
            :key="item.id"
            :articleListItem="item"
          />
        </el-tab-pane>
        <el-tab-pane label="发布">Config</el-tab-pane>
        <el-tab-pane label="审核中（0）">Role</el-tab-pane>
        <el-tab-pane label="未通过（0）">Task</el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import ListItem from "@/views/Creator/ListItem.vue";
import type { TabsPaneContext } from "element-plus";

const activeName = ref("0");

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

const activeIndex = ref("1");
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

const articleList = ref([
  {
    id: 0,
    title: "title",
    created_at: "2022-10-22",
  },
  {
    id: 1,
    title: "title",
    created_at: "2022-10-22",
  },
  {
    id: 2,
    title: "title",
    created_at: "2022-10-22",
  },
]);
</script>

<style lang="less" scoped>
.count-bar {
  display: flex;
  .tabs {
    width: 100%;
  }
}
</style>
