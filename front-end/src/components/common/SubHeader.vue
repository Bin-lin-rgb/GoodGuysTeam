<template>
  <header class="subHeader">
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      @select="handleSelect"
    >
      <el-menu-item
        index="item.community_id"
        class="content-item"
        v-for="item in ClassificationList"
        :key="item.community_id"
        >{{ item.community_name }}</el-menu-item
      >
    </el-menu>
  </header>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
// import { GetClassification } from "@/utils/api/article";
// 获取的分类列表
const ClassificationList = ref();
ClassificationList.value = [
  { community_id: 1, community_name: "综合" },
  { community_id: 2, community_name: "后端" },
  { community_id: 3, community_name: "前端" },
  { community_id: 4, community_name: "Android" },
  { community_id: 5, community_name: "iOS" },
  { community_id: 6, community_name: "人工智能" },
  { community_id: 7, community_name: "开发工具" },
  { community_id: 8, community_name: "阅读" },
];

const activeIndex = ref("1");
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

onBeforeMount(() => {
  //   GetClassification().then((res: any) => {
  //     ClassificationList.value = res.data;
  //   });
});
</script>

<style scoped lang="less">
.subHeader {
  height: 40px;
  width: 100%;
  position: fixed;
  top: 70px;
  left: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 6;
  .el-menu-demo {
    border-top: solid 1px #dcdfe6;
    border-bottom: solid 1px #dcdfe6;
    display: flex;
    justify-content: center;
    width: 1700px;
    .content-item {
      height: 38px;
    }
  }
}
</style>
