import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home/HomeIndex.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home,
    },
    {
      path: "/login",
      name: "Login",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("@/views/UserInfo/UserLogin.vue"),
    },
    {
      path: "/creator",
      name: "Creator",
      component: () => import("../views/Creator/CreatorIndex.vue"),
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
    },
    {
      path: "/success",
      name: "SuccessPage",
      component: () => import("../views/Creator/SuccessPage.vue"),
    },
    {
      path: "/userinfo",
      name: "UserInfo",
      component: () => import("../views/UserInfo/UserIndex.vue"),
    },
    {
      path: "/details",
      name: "ArticleDetails",
      component: () => import("../views/ArticleDetails/ArticleIndex.vue"),
    },
    {
      path: "/test",
      name: "Test",
      component: () => import("@/components/testData.vue"),
    },
  ],
});

export default router;
