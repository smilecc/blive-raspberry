import { createApp } from "vue";
import App from "./App.vue";
import { Router } from "./routes";
import "./main.css";

const meta = document.createElement("meta");
meta.name = "naive-ui-style";
document.head.appendChild(meta);

createApp(App).use(Router).mount("#app");
