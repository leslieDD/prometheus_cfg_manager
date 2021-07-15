import { createApp } from 'vue'
import App from './App.vue'

import router from './router/index'

import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import './assets/icon/iconfont.css'
import IconSvg from '@/components/IconSvg.vue'
import './assets/svg/iconfont.js'
import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale/lang/zh-cn'
 
const app = createApp(App)
app.use(ElementPlus, { locale })
app.use(router)
app.component('icon-svg', IconSvg)
app.mount('#app')
