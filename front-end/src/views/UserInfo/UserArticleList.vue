<template>
  <a-empty v-if="isEmpty" />
  <ListItem
    v-else
    v-for="item in articleList"
    :key="item.id"
    :articleListItem="item"
  />
</template>

<script lang="ts" setup>
import { ref, onBeforeMount } from "vue";
import ListItem from "@/views/UserInfo/ListItem.vue";
import { GetPostListByUserId } from "@/utils/api/article";
const isEmpty = ref(true);

const articleList = ref();
const loadBlog = async () => {
  const res = await GetPostListByUserId();
  if (res.data.length == 0) {
    isEmpty.value = true;
    return;
  }
  isEmpty.value = false;
  articleList.value = res.data;
  for (let i = 0; i < articleList.value.length; i++) {
    articleList.value[i].id = i;
  }
};

onBeforeMount(() => {
  loadBlog();
});
</script>
