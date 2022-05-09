<template>
  <div class="w-[400px]">
    <div class="mt-9 text-[22px] font-semibold text-gray-700">
      {{ song.name }}
    </div>
    <div
      v-show="song.subName"
      class="mt-3 overflow-hidden overflow-ellipsis whitespace-nowrap text-sm"
    >
      {{ song.subName }}
    </div>
    <div class="mt-4 flex">
      <div class="mr-2 flex text-[13px] text-gray-600">
        专辑：<span class="w-24 cursor-pointer overflow-hidden overflow-ellipsis whitespace-nowrap text-blue-400">{{
          song.album
        }}</span>
      </div>
      <div class="mr-1 flex text-[13px] text-gray-600">
        歌手：<span
          class="w-24 cursor-pointer overflow-hidden overflow-ellipsis whitespace-nowrap text-blue-400"
          title="前往歌手详情"
          >{{ song.artists[0].name }}</span
        >
      </div>
    </div>
    <!-- 歌词滚动 -->

    <n-scrollbar
      ref="scrollBarRef"
      class="mt-6 max-h-80 w-[380px] border-r border-solid border-gray-200"
    >
      <div
        class="cursor-pointer"
        v-if="formatedLyrics && formatedLyrics.length > 0"
      >
        <div
          v-for="(lyric, index) in formatedLyrics"
          :key="index"
          class="lyric-item bg-opacity-5 pb-5 text-sm text-gray-600"
          :class="lyricClass(lyric.text, index)"
        >
          <div>{{ lyric.text }}</div>
          <div>{{ lyric.transText }}</div>
        </div>
      </div>
      <div
        v-else
        class="flex h-full items-center justify-center text-sm text-gray-800"
      >
        纯音乐，请您欣赏
      </div>
    </n-scrollbar>
  </div>
</template>

<script
  lang="ts"
  setup
>
import { ref, watch, computed } from "vue";
import { formatLyric } from "@//utils";
import { usePlayerStore } from "@/stores";
import { IPlaySong } from "@/interfaces";
import { NScrollbar } from "naive-ui";

const props = defineProps<{
  song: IPlaySong;
  lyric: string /** 歌曲歌词，为空字符串代表歌曲没有歌词 */;
  transLyric?: string /** 翻译歌曲歌词，为空字符串代表歌曲没有歌词 */;
}>();

const playerStore = usePlayerStore();

/** 滚动条ref */
const scrollBarRef = ref();
/** 当前播放的歌词的索引 */
const currentLyricIndex = ref(0);

/** 格式化后的歌词数组 */
const formatedLyrics = computed(() => {
  if (!props.transLyric) {
    return formatLyric(props.lyric);
  } else {
    const lyricList = formatLyric(props.lyric);
    const transLyricList = formatLyric(props.transLyric);
    return lyricList.map((item) => {
      const findResult = transLyricList.find((transItem) => item.time === transItem.time);
      if (findResult) {
        return {
          ...item,
          transText: findResult.text,
        };
      }
      return item;
    });
  }
});
/** 歌曲当前播放时间 */

const lyricClass = (text: string, index: number) => {
  return {
    "text-[18px] text-gray-900 font-bold":
      currentLyricIndex.value === index - 1 && !text.includes("作词") && !text.includes("作曲"),
  };
};

watch(
  () => playerStore.currentTime,
  (newTime, oldTime) => {
    if (newTime !== oldTime) {
      /** 获取比当前播放时间大的第一个元素 */
      for (let i = 0; i < formatedLyrics.value.length; i++) {
        if (Math.floor(formatedLyrics.value[i].time) === Math.floor(playerStore.currentTime)) {
          currentLyricIndex.value = i - 1;
          break;
        } else if (Math.floor(formatedLyrics.value[i].time) > Math.floor(playerStore.currentTime)) {
          currentLyricIndex.value = i - 2;
          break;
        }
      }
      let height = 0;
      if (currentLyricIndex.value !== 0) {
        const lyricEles = document.querySelectorAll(".lyric-item");
        lyricEles.forEach((ele, index) => {
          if (currentLyricIndex.value >= index) {
            height += ele.clientHeight;
          }
        });
        if (height >= 160) {
          scrollBarRef.value?.scrollTo({
            top: height - 160 + 20,
            behavior: "smooth",
          });
        } else {
          scrollBarRef.value?.scrollTo({
            top: 0,
            behavior: "smooth",
          });
        }
      }
    }
  }
);
</script>

<!-- <style lang="scss" scoped>
:deep(.el-scrollbar__view) {
  height: 100%;
}
</style> -->
