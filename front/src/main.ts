import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import { Router } from "./routes";
import "./main.css";
import VueProgressBar from "@aacassandra/vue3-progressbar";

const meta = document.createElement("meta");
meta.name = "naive-ui-style";
document.head.appendChild(meta);

createApp(App)
  .use(VueProgressBar as any, {
    color: '#845EC2',
    thickness: "6px",
    autoRevert: true,
    autoFinish: false,
  })
  .use(createPinia())
  .use(Router)
  .mount("#app");
