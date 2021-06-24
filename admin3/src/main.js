import { createApp } from 'vue'
import App from './App.vue'

import router from './router/index'

import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

import { install } from 'vs-tree'
// import vueTree from 'vue-tree-jf'
import 'vs-tree/dist/vs-tree.css'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.use(install)
// app.use(vueTree)
app.mount('#app')
