<template>
  <div class="main-wrap">
    <div class="font-2">个人资料</div>
    <a-divider />
    <div class="input-item">
      <div class="item-text-area">用户名</div>
      <el-input
        v-model="data.username"
        maxlength="20"
        placeholder="请填写你的用户名"
        show-word-limit
        type="text"
      />
    </div>
    <a-divider />
    <div class="input-item">
      <div class="item-text-area">职位</div>
      <el-input
        v-model="data.position"
        maxlength="50"
        placeholder="请填写你的职位"
        show-word-limit
        type="text"
      />
    </div>
    <a-divider />
    <div class="input-item">
      <div class="item-text-area">公司</div>
      <el-input
        v-model="data.company"
        maxlength="50"
        placeholder="请填写你的公司"
        show-word-limit
        type="text"
      />
    </div>
    <a-divider />
    <div class="input-item">
      <div class="item-text-area">个人主页</div>
      <el-input
        v-model="data.user_page"
        maxlength="100"
        placeholder="请填写你的个人主页"
        show-word-limit
        type="text"
      />
    </div>
    <a-divider />
    <div class="input-item">
      <div class="item-text-area">个人介绍</div>
      <el-input
        v-model="data.user_introduce"
        maxlength="100"
        placeholder="请填写你的个人介绍"
        show-word-limit
        type="textarea"
      />
    </div>

    <div class="bottom-save">
      <a-button type="primary" size="large" @click="HanderCommit"
        >保存修改</a-button
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { ElMessage } from "element-plus";
import { reactive } from "vue";
import { updateProfile } from "@/utils/api/login";

const data = reactive({
  username: "",
  position: "",
  company: "",
  user_page: "",
  user_introduce: "",
});

const HanderCommit = async () => {
  const res: any = await updateProfile(data);
  if (res.code === 1000) {
    ElMessage({
      message: "修改成功",
      type: "success",
    });
  }
};
</script>

<style scoped lang="less">
.main-wrap {
  width: 700px;
  .font-2 {
    font-size: large;
  }
  .input-item {
    display: flex;
    .item-text-area {
      width: 100px;
      height: 35px;
      text-align: center;
      line-height: 35px;
    }
  }
  .bottom-save {
    margin: 50px;
  }
}
</style>
