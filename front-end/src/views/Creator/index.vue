<template>
  <CommHeader />
  <div class="md"><md-editor v-model="text" /></div>
</template>

<script setup lang="ts">
import CommHeader from "@/components/common/CommHeader.vue";
import { ref } from "vue";
// md-editor-v3
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
const text = ref("Hello Editor!");
console.log(text.value);
</script>

<style scoped>
.md {
  height: 100%;
  width: 100%;
  position: fixed;
  top: 70px;
  left: 0;
}
</style>
