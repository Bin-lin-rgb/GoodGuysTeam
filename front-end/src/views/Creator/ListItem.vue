<template>
  <div class="info-main">
    <div class="title-time" @click="gotoDetails">
      <div class="main-title">{{ articleListItem.title }}</div>
      <div class="main-info">{{ articleListItem.created_at2 }}</div>
    </div>
    <div class="buttons">
      <a-button type="outline" class="edit">编辑</a-button>

      <a-button status="danger" type="outline" class="delete" @click="Delete">
        <template #icon>
          <icon-delete />
        </template>
        <template #default>删除</template>
      </a-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconDelete } from "@arco-design/web-vue/es/icon";
import { DeleteBlog } from "@/utils/api/article";
import { useRouter } from "vue-router";
const router = useRouter();

const props = defineProps<{
  articleListItem: any;
}>();

const emits = defineEmits(["SuccessDelete"]);

function gotoDetails() {
  router.push({
    path: "/details",
    query: { id: props.articleListItem.post_id },
  });
}

const Delete = async () => {
  const res: any = await DeleteBlog({ id: props.articleListItem.post_id });
  console.log(res);
  if (res.code == 1011) {
    emits("SuccessDelete");
  }
};
</script>

<style lang="less" scoped>
.info-main {
  height: 80px;
  padding: 10px;
  border-bottom: solid #e5e6eb 1px;
  cursor: pointer;
  .title-time {
    display: inline-block;
    .main-title {
      font-size: 22px;
      width: 750px;
      height: 30px;
      overflow: hidden;
    }
  }
  .buttons {
    float: right;
    padding: 15px;
    .edit {
      margin-right: 8px;
    }
  }
}
.info-main:hover {
  background: #fafafa;
}
</style>
