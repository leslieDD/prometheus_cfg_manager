import { createApp } from 'vue'
import App from './App.vue'

import router from './router/index'

import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import './assets/icon/iconfont.css'
import IconSvg from '@/components/IconSvg.vue'
import './assets/svg/iconfont.js'
 
const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.component('icon-svg', IconSvg)
app.mount('#app')
