import {createApp,} from 'vue'
import App from './App.vue'
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import "@icon-park/vue-next/styles/index.css";
import router from './router'
import store from './store'
const app = createApp(App)

app.use(ElementPlus).use(router).use(store)

app.mount('#app')
