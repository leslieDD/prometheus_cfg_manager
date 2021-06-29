import { createRouter, createWebHashHistory } from "vue-router";

import Manager from '@/views/Manager.vue'

const routes = [
    {
      path: '/',
      name: 'Manager',
      component: Manager,
      children: [
        {
          path: '/empty',
          name: 'empty',
          component: () => import('@/views/Empty.vue')
        },
        {
          path: '/base-labels',
          name: 'baseLabels',
          component: () => import('@/views/BaseLabels.vue')
        },
        {
          path: '/re-labels',
          name: 'reLabels',
          component: () => import('@/views/RelabelConfig.vue')
        }
      ]
    }
]

export default  createRouter({
  // 指定路由的模式,此处使用的是hash模式
  history: createWebHashHistory(),
  routes // short for `routes: routes`  
})

