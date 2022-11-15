import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home/HomeIndex.vue";

// class meta {
//   isAuth: boolean;
//   title: string;
//   // 构造方法
//   constructor(isAuth: boolean, title: string) {
//     this.isAuth = isAuth;
//     this.title = title;
//   }
// }

// const creator = new meta(true, "创造者中心");

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home,
      meta: {
        isAuth: false,
        title: "Blog",
      },
    },
    {
      path: "/login",
      name: "Login",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("@/views/UserInfo/UserLogin.vue"),
      meta: {
        isAuth: false,
        title: "登录",
      },
    },
    {
      path: "/creator",
      name: "Creator",
      component: () => import("../views/Creator/CreatorIndex.vue"),
      meta: {
        isAuth: true,
        title: "创作者中心",
      },
      children: [
        {
          path: "list",
          component: () => import("../views/Creator/ShowArticle.vue"),
        },
        {
          path: "index",
          component: () => import("../views/Creator/RightIndex.vue"),
        },
      ],
    },
    {
      path: "/write",
      name: "Write",
      component: () => import("../views/Creator/WriteArticle.vue"),
      meta: {
        isAuth: true,
        title: "写文章",
      },
    },
    {
      path: "/success",
      name: "SuccessPage",
      component: () => import("../views/Creator/SuccessPage.vue"),
      meta: {
        isAuth: false,
        title: "操作成功",
      },
    },
    {
      path: "/userinfo",
      name: "UserInfo",
      component: () => import("../views/UserInfo/UserIndex.vue"),
      meta: {
        isAuth: true,
        title: "用户中心",
      },
      children: [
        {
          path: "list",
          component: () => import("../views/UserInfo/UserArticleList.vue"),
        },
      ],
    },
    {
      path: "/userinfo/setting",
      name: "Setting",
      component: () => import("../views/UserInfo/EditIndex.vue"),
      meta: {
        isAuth: true,
        title: "设置",
      },
      children: [
        {
          path: "profile",
          component: () => import("../views/UserInfo/EditProfile.vue"),
        },
        {
          path: "account",
          component: () => import("../views/UserInfo/EditAccount.vue"),
        },
      ],
    },
    {
      path: "/details",
      name: "ArticleDetails",
      component: () => import("../views/ArticleDetails/ArticleIndex.vue"),
      meta: {
        isAuth: false,
        title: "文章详情",
      },
    },
    {
      path: "/test",
      name: "Test",
      component: () => import("@/components/testData.vue"),
    },
    {
      path: "/NoFunction",
      name: "NoFunction",
      component: () => import("@/components/common/ErrPage.vue"),
      meta: {
        isAuth: false,
        title: "额哦,403啦",
      },
    },
  ],
});

// pinia
import { useUserStore } from "@/stores/user";
router.beforeEach((to, from, next) => {
  const userStore = useUserStore();
  if (to.meta.isAuth === "ture" && userStore.token !== "")
    next({ name: "Login" });
  else next();
});

//全局后置路由守卫————初始化的时候被调用、每次路由切换之后被调用
router.afterEach((to, from) => {
  let str = "Blog";
  if (String(to.meta.title) !== "") {
    str = String(to.meta.title);
  }
  document.title = str;
});

export default router;
