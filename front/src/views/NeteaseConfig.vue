<script lang="ts" setup>
import { onBeforeMount, reactive } from "vue";
import {
  NButton,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  useMessage,
  NThing,
  NAlert,
} from "naive-ui";
import { ConfigService, NeteaseService } from "@/services";

const message = useMessage();
const state = reactive({
  apiHost: "",
  cookie: "",
  profile: null as any,
  mobile: null,
  captcha: null,
});

onBeforeMount(() => {
  ConfigService.getConfig("netease_api_host").then(({ data }) => {
    state.apiHost = data.data.value;
  });
  ConfigService.getConfig("netease_cookie").then(({ data }) => {
    state.cookie = data.data?.value || "";
    getNeteaseAccountInfo();
  });
});

function sendCaptcha() {
  NeteaseService.sendCaptcha(state.apiHost, `${state.mobile}`).then(
    ({ data }) => {
      if (data.code != 200) {
        message.error(data.message);
      } else {
        message.success("验证码发送成功");
      }
    }
  );
}

function neteaseLogin() {
  NeteaseService.loginByCaptcha(
    state.apiHost,
    `${state.mobile}`,
    `${state.captcha}`
  )
    .then(({ data }) => {
      ConfigService.setConfig("netease_cookie", data.cookie);
      state.cookie = data.cookie;
      message.success("登录成功");
      getNeteaseAccountInfo();
    })
    .catch((error) => {
      if (error.isAxiosError) {
        message.error(error.response?.data.message);
      }
    });
}

function getNeteaseAccountInfo() {
  NeteaseService.getAccountInfo(state.apiHost, state.cookie).then(
    ({ data }) => {
      if (data.profile) {
        state.profile = data.profile;
      } else {
        state.cookie = "";
      }
    }
  );
}
</script>

<template>
  <div class="p-5">
    <n-form>
      <n-thing title="服务配置">
        <n-form-item label="网易云服务地址">
          <n-input v-model:value="state.apiHost" />
        </n-form-item>
      </n-thing>

      <n-thing title="用户登录">
        <div v-if="state.cookie && state.profile" class="mb-5">
          <n-alert title="网易云音乐账户已登录" type="success">
            <div>已经登录账户 {{ state.profile.nickname }}</div>
          </n-alert>
        </div>
        <n-form-item label="手机号码">
          <n-input-number v-model:value="state.mobile" :show-button="false" />
        </n-form-item>
        <n-form-item label="验证码">
          <n-input-number v-model:value="state.captcha" :show-button="false" />
          <n-button class="!ml-2" @click="sendCaptcha" type="info" secondary
            >发送验证码</n-button
          >
        </n-form-item>
        <n-button type="primary" @click="neteaseLogin">登录</n-button>
      </n-thing>
    </n-form>
  </div>
</template>
