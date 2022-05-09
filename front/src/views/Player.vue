<script
  lang="ts"
  setup
>
import { usePlayerStore } from "@/stores";
import RotateCover from "./components/RotateCover.vue";
import SongLyric from "./components/SongLyric.vue";
import { onMounted, ref, watch, getCurrentInstance } from "vue";
import _ from "lodash";
import { useRouter } from "vue-router";
import { NCard } from "naive-ui";

const instance = getCurrentInstance();

const playerStore = usePlayerStore();
const router = useRouter();

/**
 * 初始化音频
 */
function initAudio() {
  if (playerStore.playerRef && instance) {
    const progress = instance.appContext.config.globalProperties.$Progress;

    // 监听播放进度
    playerStore.playerRef.ontimeupdate = (el) => {
      _.debounce(() => {
        playerStore.currentTime = playerStore.playerRef?.currentTime || 0;

        const timePercent = (playerStore.currentTime / (playerStore.currentSong!.duration! / 1000)) * 100;
        console.log(timePercent);
        progress.set(timePercent < 1 ? 1 : timePercent);
      }, 1000)();
    };

    // 监听音乐播放结束
    playerStore.playerRef.onended = () => {
      // 结束时如果没有新歌曲了 则重新播放当前音乐
      if (playerStore.musicList.length > 1) {
        playerStore.musicList.shift();
      }

      playerStore.playerRef?.play();
    };
  }
}

onMounted(() => {
  initAudio();
});
</script>

<template>
  <div
    class="flex min-h-screen flex-col justify-center bg-slate-100"
    v-if="playerStore.currentSong"
  >
    <div class="flex justify-center">
      <div
        class="mr-32 flex items-center"
        @click="() => playerStore.playerRef?.play()"
        @dblclick="() => router.push('/')"
      >
        <rotate-cover :img="playerStore.currentSong.coverImg" />
      </div>
      <song-lyric
        :song="playerStore.currentSong"
        :lyric="playerStore.currentSong.lyric || ''"
      />
    </div>
    <div class="mt-10 flex h-56 justify-center">
      <n-card
        title="播放列表"
        class="!w-96 !overflow-hidden"
      >
        <div
          v-for="(song, index) in playerStore.musicList"
          :key="index"
        >
          {{ index + 1 }}. {{ `${song.name} - ${song.artists[0].name}` }}
        </div>
      </n-card>
      <n-card
        title="点歌说明"
        class="ml-10 !w-96 !overflow-hidden"
      >
        <p>发送弹幕：</p>
        <p class="mt-2"><strong>点歌歌曲名</strong> 或 <strong>点歌歌曲名-歌手</strong></p>
      </n-card>
    </div>
  </div>
</template>
