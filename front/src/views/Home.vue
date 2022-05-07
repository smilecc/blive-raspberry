<script lang="ts" setup>
import { reactive } from "vue";
import { NButton, NForm, NFormItem, NInput, NInputNumber } from "naive-ui";

const state = reactive({
  live: {
    roomId: 0,
    liveUrl: "",
    livePassword: "",
  },
});
</script>

<template>
  <div class="p-5">
    <n-form>
      <n-form-item label="直播间ID">
        <n-input-number v-model:value="state.live.roomId" />
      </n-form-item>
      <n-form-item label="直播服务器地址（RTMP）">
        <n-input v-model:value="state.live.liveUrl" />
      </n-form-item>
      <n-form-item label="直播串流密钥">
        <n-input v-model:value="state.live.livePassword" />
      </n-form-item>
    </n-form>
  </div>
</template>
