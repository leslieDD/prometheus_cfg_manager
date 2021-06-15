// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
// import Element from 'element-ui'
// import 'element-ui/lib/theme-chalk/index.css'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

import * as filters from './filters'

Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false
// Vue.use(Element) // global filters
Vue.use(ElementPlus) // global filters

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
