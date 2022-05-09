<script
  lang="ts"
  setup
>
import { reactive, onBeforeMount } from "vue";
import { NButton, NForm, NFormItem, NInput, NInputNumber, NThing, useMessage } from "naive-ui";
import { ConfigService, LiveService } from "@/services";
import { usePlayerStore } from "@/stores";
import ConfigLayout from "@/components/ConfigLayout.vue";
import VConsole from "vconsole";

const playerStore = usePlayerStore();
const message = useMessage();
const state = reactive({
  liveState: false,
  live: {
    roomId: 0,
    liveUrl: "",
    livePassword: "",
  },
});

onBeforeMount(() => {
  getLiveState();
  ConfigService.getConfig("live_room").then(({ data }) => {
    if (data.data) {
      state.live = JSON.parse(data.data.value);
    }
  });
});

function getLiveState() {
  LiveService.getLiveState().then(({ data }) => {
    state.liveState = data.data;
  });
}

function saveLiveConfig() {
  return ConfigService.setConfig("live_room", JSON.stringify(state.live)).then(() => {
    message.info("配置已保存");
  });
}

async function startLive() {
  await saveLiveConfig();
  await LiveService.startLive(state.live.roomId);
  state.liveState = true;
}

async function stopLive() {
  await LiveService.stopLive();
  state.liveState = false;
}

function onDebug() {
  new VConsole();
}
</script>

<template>
  <config-layout>
    <div
      class="p-5"
      id="index-page"
    >
      <n-form>
        <n-thing title="直播间设置">
          <n-form-item label="直播间ID">
            <n-input-number v-model:value="state.live.roomId" />
          </n-form-item>
          <div>
            <n-button
              type="error"
              @click="stopLive"
              v-if="state.liveState"
              >停止直播</n-button
            >
            <n-button
              type="primary"
              @click="startLive"
              v-else
              >开始直播</n-button
            >

            <n-button
              type="primary"
              secondary
              @click="saveLiveConfig"
              class="!ml-2"
              >保存配置</n-button
            >
            <n-button
              type="info"
              secondary
              @click="onDebug"
              class="!ml-2"
              >Debug</n-button
            >
          </div>
        </n-thing>
      </n-form>
    </div>
  </config-layout>
</template>
