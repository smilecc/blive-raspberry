import { createRouter, createWebHashHistory } from "vue-router";

export const Router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      component: () => import("@/views/Home.vue"),
    },
    {
      path: "/netease_config",
      component: () => import("@/views/NeteaseConfig.vue"),
    },
  ],
});
