<template>
  <CommHeader />
  <SubHeader />
  <div class="article-list">
    <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
      <router-link
        v-for="i in count"
        :key="i"
        class="infinite-list-item"
        to="/details"
      >
        这是第{{ i }}篇文章,点击跳转详情页
      </router-link>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import CommHeader from "@/components/common/CommHeader.vue";
import SubHeader from "@/components/common/SubHeader.vue";
import { ref } from "vue";
// import { useRouter } from "vue-router";
// const router = useRouter();

const count = ref(0);
const load = () => {
  count.value += 2;
};
</script>

<style scoped>
.article-list {
  width: 1000px;
  height: 900px;
  position: fixed;
  top: 130px;
  left: 12%;
  align-items: center;
  justify-content: center;
}
.infinite-list {
  height: 100%;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>
