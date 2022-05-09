<template>
  <div class="relative h-80 w-80 rounded-full border-8 border-gray-200">
    <img class="absolute top-0 left-0 w-full" src="/image/singlecover.svg" />
    <img
      class="absolute top-[51px] left-[52px] w-[200px] rounded-full"
      :class="{ 'animate-spin-30': !isPause }"
      :src="img"
    />
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { usePlayerStore } from "@/stores";

defineProps<{
  img: string /** 歌曲封面 */;
}>();

const playerStore = usePlayerStore();

/** 歌曲是否已暂停 */
const isPause = computed(() => playerStore.isPause);
</script>

<style lang="scss" scoped></style>
