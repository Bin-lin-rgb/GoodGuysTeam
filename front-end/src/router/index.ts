import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/HomeIndex.vue";

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
      component: () => import("../views/UserLogin.vue"),
    },
    {
      path: "/creator",
      name: "Creator",
      component: () => import("../views/Creator/index.vue"),
    },
  ],
});

export default router;
