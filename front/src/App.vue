<script setup lang="ts">
import { NConfigProvider, NMessageProvider, zhCN } from "naive-ui";
import Navbar from "@/components/Navbar.vue";
</script>

<template>
  <NConfigProvider :locale="zhCN">
    <NMessageProvider>
      <div
        class="dark:bg-dark-background relative flex h-full flex-row bg-white text-xl text-black dark:text-white"
      >
        <Navbar />
        <main class="min-h-screen flex-grow bg-gray-50">
          <div className="flex flex-shrink">
            <div className="w-full sm:w-main pb-20 sm:pb-0">
              <router-view />
            </div>
            <div className="hidden xl:block w-right-bar"></div>
          </div>
        </main>
      </div>
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped>
.content {
  width: 860px;
  margin: 0 auto;
}
</style>

<style>
.n-thing-header__title {
  font-weight: bold !important;
}
</style>
