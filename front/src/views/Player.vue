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
import { NCard, useNotification } from "naive-ui";

const instance = getCurrentInstance();

const playerStore = usePlayerStore();
const router = useRouter();
const notification = useNotification();

/**
 * 初始化音频
 */
function initAudio() {
  if (playerStore.playerRef && instance) {
    const progress = instance.appContext.config.globalProperties.$Progress;
    let isFinish = true;

    playerStore.playerRef.oncanplay = () => {
      console.log("music oncanplay");
      if (isFinish) {
        playerStore.playerRef?.play();
        isFinish = false;
      }
    };

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
      } else {
        playerStore.playerRef?.play();
      }

      isFinish = true;
    };
  }
}

onMounted(() => {
  initAudio();

  playerStore.onDanmuCommand = (event) => {
    if (event.commandName == "点歌") {
      notification.success({
        closable: false,
        duration: 5000,
        title: "收到点歌",
        content: `由 [${event.senderName}] 所点的 [${event.arg1}]，请等待下载`,
      });
    }
  };
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
        @click="() => router.push('/')"
      >
        <rotate-cover :img="playerStore.currentSong.coverImg" />
      </div>
      <song-lyric
        :song="playerStore.currentSong"
        :lyric="playerStore.currentSong.lyric || ''"
      />
    </div>
    <div class="mt-5 flex h-56 justify-center">
      <n-card
        title="播放列表"
        class="!w-96 !overflow-hidden"
        @dblclick="() => playerStore.playerRef?.play()"
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
  <div v-else>
    <button @click="() => router.push('/')">去配置中心</button>
  </div>
</template>
