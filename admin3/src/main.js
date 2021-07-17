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

import 'vite-plugin-svg-icons/register'
// 需要全局引入再添加
import singleSvg from './components/singleSvg.vue' // 全局svg图标组件
import Highlight from './directive/highlight.js' // 这里是你项目highlight.js所在路径

const app = createApp(App)
app.use(ElementPlus, { locale })
app.use(router)
app.use(Highlight)
app.component('icon-svg', IconSvg)
app.component('single-svg', singleSvg)
app.mount('#app')
