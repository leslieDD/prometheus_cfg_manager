import { createRouter, createWebHashHistory } from "vue-router";

import Manager from '@/views/Manager.vue'

const routes = [
    {
        path: '/',
        name: 'Manager',
        component: Manager
    }
]

export default  createRouter({
    // 指定路由的模式,此处使用的是hash模式
    history: createWebHashHistory(),
    routes // short for `routes: routes`  
})

