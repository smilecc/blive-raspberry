<script
  setup
  lang="ts"
>
import { NConfigProvider, NMessageProvider, zhCN } from "naive-ui";
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
      <vue-progress-bar></vue-progress-bar>
      <router-view />
      <audio
        ref="audioPlayerRef"
        class="hidden"
        :src="playerStore.currentSong?.songUrl"
        controls
      ></audio>
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped></style>
