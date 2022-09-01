<template>
  <div class="item" @click="gotoDetails(articleListItem.post_id)">
    <div class="info-bar">
      {{ articleListItem.author_name }}
      <div class="divider"></div>
      {{ articleListItem.created_at }}
      <div class="divider"></div>
      {{ articleListItem.community_name }}
    </div>

    <div class="info-main">
      <div class="main-title">{{ articleListItem.title }}</div>
      <div class="main-context">{{ articleListItem.content }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
const router = useRouter();
const props = defineProps<{
  articleListItem: any;
}>();

function gotoDetails(id: any) {
  router.push({
    path: "/details",
    query: { id: props.articleListItem.post_id },
  });
}
</script>

<style scoped lang="less">
@normal_text_color: #86909c;

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
</style>
