<script lang="ts" setup>
const Routes = [
  {
    title: "首页",
    path: "/",
  },
  {
    title: "网易云配置",
    path: "/netease_config",
  },
  {
    title: "→ 播放器",
    path: "/player",
  },
];
</script>

<template>
  <nav
    class="absolute bottom-0 flex w-screen flex-grow justify-end pr-0 sm:static sm:w-left-bar-m sm:pr-4 xl:w-left-bar"
  >
    <div
      class="fixed bottom-0 z-20 flex w-screen justify-around border-t bg-white sm:fixed sm:top-0 sm:w-left-bar-m sm:flex-col sm:justify-start sm:border-t-0 xl:w-left-bar"
    >
      <h1 class="mt-4 mb-8 hidden text-3xl font-bold dark:text-white sm:block">
        应用配置
      </h1>
      <router-link
        v-for="item in Routes"
        :key="item.path"
        class="navLink"
        :to="item.path"
        >{{ item.title }}</router-link
      >
    </div>
  </nav>
</template>

<style scoped>
.navLink {
  @apply mb-1.5 
  rounded-xl 
  py-2.5 
  pl-4 
  pr-4 
  text-lg 
  font-bold
  text-gray-400 
  transition-all 
  duration-300;
}
.navLink:hover,
.router-link-active {
  @apply bg-transparent 
    text-blue-nav-text 
    dark:bg-blue-800 
    dark:text-blue-200 
    sm:bg-blue-nav-sky;
}
</style>
