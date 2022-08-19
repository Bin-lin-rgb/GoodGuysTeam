<template>
  <div class="header">
    <el-input
      v-model="title"
      placeholder="输入文章标题..."
      clearable
      maxlength="30"
      show-word-limit
      type="text"
      class="w-50 m-2 input"
      size="large"
    />
    <span class="plainText">文章将自动保存至草稿箱</span>
    <el-button type="primary" plain :disabled="!isLogin">草稿箱</el-button>

    <el-button type="primary" @click="showClick" :disabled="!isLogin"
      >发布</el-button
    >

    <router-link to="/login" class="content-login" v-if="!isLogin"
      >登录 / 注册</router-link
    >

    <!-- 划过头像显示  -->
    <el-dropdown v-else>
      <span class="el-dropdown-link">
        <el-avatar
          class="avatar"
          @mouseenter="isShow = true"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      /></span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item @click="gotoWrite">写文章</el-dropdown-item>
          <el-dropdown-item @click="gotoUserinfo">我的主页</el-dropdown-item>
          <el-dropdown-item>设置</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
  <div class="md" @click="activePubilsh = false">
    <md-editor v-model="text" />
  </div>
  <div class="pop-up" v-show="activePubilsh">
    <h2 id="publish-article">发布文章</h2>
    <h2 id="choose">请选择分类</h2>
    <div class="mb-2 flex items-center text-sm">
      <el-radio-group v-model="radio" class="ml-4">
        <el-radio
          :label="item.community_id"
          size="large"
          v-for="item in ClassificationList"
          :key="item.community_id"
          >{{ item.community_name }}</el-radio
        >
      </el-radio-group>
    </div>
    <div class="submit-area">
      <el-button type="primary" @click="onSubmit">确定</el-button>
      <el-button @click="activePubilsh = false">取消</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
// pinia
// import { useUserStore } from "@/stores/user";
// Router
import { useRouter } from "vue-router";
// api
import { getUserInfo } from "@/utils/api/login";
import { PostArticle } from "@/utils/api/article";
// md-editor-v3
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { ElMessage } from "element-plus";

const title = ref("");
const text = ref("Hello Editor!");
console.log(text.value);

const router = useRouter();
//用户是否是登录状态
const isLogin = ref(false);
//显示用户更多数据
const isShow = ref(false);
const userInfo = ref();

function gotoUserinfo() {
  router.push("/userinfo");
}
function gotoWrite() {
  const routeData = router.resolve({
    // query: {id:deviceId},
    path: "/write",
  });
  window.open(routeData.href, "_blank");
}

onBeforeMount(() => {
  getUserInfo().then((res: any) => {
    if (res.code == 1000) {
      isLogin.value = true;
      userInfo.value = res.data;
    }
  });
});

// 发布
const activePubilsh = ref(false);
const radio = ref("1");
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

function showClick() {
  activePubilsh.value = true;
}

function onSubmit() {
  PostArticle({
    title: title.value,
    content: text.value,
    community_id: radio.value,
  }).then(() => {
    ElMessage({
      message: "发布成功！",
      type: "success",
    });
    router.push("/success");
  });
}
</script>

<style scoped lang="less">
.header {
  height: 70px;
  width: 100%;
  position: fixed;
  top: 0px;
  left: 0;
  .input {
    margin: 10px 10px 10px;
    width: 60%;
  }
  .plainText {
    margin: 27px 30px 27px 70px;
    color: #afafaf;
    font-size: 16px;
  }
  .content-login {
    margin-left: 20px;
  }

  .avatar {
    margin: 15px 45px 15px;
  }
}
.md {
  height: 100%;
  width: 100%;
  position: fixed;
  top: 70px;
  left: 0;
}
.pop-up {
  position: fixed;
  top: 65px;
  right: 50px;
  width: 430px;
  height: 310px;
  padding: 20px;
  border: solid 1px #dcdfe6;
  background: rgb(255, 255, 255);
  #publish-article {
    font-size: 23px;
    border-bottom: solid 1px #dcdfe6;
    padding-bottom: 10px;
  }
  #choose {
    font-size: medium;
    margin-top: 10px;
    margin-bottom: 12px;
  }
  .submit-area {
    margin-top: 53px;
    float: right;
  }
}
</style>
