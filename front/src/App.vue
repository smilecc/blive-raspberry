<script
  setup
  lang="ts"
>
import { NConfigProvider, NMessageProvider, NNotificationProvider, zhCN } from "naive-ui";
import { onMounted, ref } from "vue";
import { usePlayerStore } from "./stores";

const playerStore = usePlayerStore();
const audioPlayerRef = ref<HTMLAudioElement>();

onMounted(() => {
  playerStore.playerRef = audioPlayerRef.value;
  playerStore.connectWebsocket();
});
</script>

<template>
  <NConfigProvider :locale="zhCN">
    <NMessageProvider>
      <NNotificationProvider
        :max="3"
        placement="bottom-left"
      >
        <vue-progress-bar></vue-progress-bar>
        <router-view />
        <audio
          ref="audioPlayerRef"
          class="hidden"
          :src="playerStore.currentSong?.songUrl"
          controls
        ></audio>
      </NNotificationProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped></style>
