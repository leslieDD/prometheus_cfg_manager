import { createRouter, createWebHashHistory } from "vue-router";

// import Manager from '@/views/Manager.vue'
import Menu from '@/views/Menu.vue'

const routes = [
    {
      path: '/',
      name: 'index',
      // director: ''
      component: Menu,
      children: [
        {
          path: '/ip-manager',
          name: 'ipManager',
          title: 'IP管理',
          components: {
            ipManager: () => import('@/views/Manager.vue')
          }
        },
        {
          path: '/jobs',
          name: 'jobs',
          title: '分组管理',
          components:  {
            jobs: () => import('@/views/Jobs.vue')
          }
        },
        {
          path: '/preview',
          name: 'preview',
          title: '配置预览',
          components: {
            "preview": () => import('@/views/Preview.vue')
          }
        },
        {
          path: '/ftree',
          name: 'ftree',
          title: '分组预览',
          components: {
            'ftree': () => import('@/views/FTree.vue')
          }
        },
        {
          path: '/notice-manager',
          name: 'noticeManager',
          title: '告警管理',
          components: {
            'noticeManager': () => import('@/views/Notice.vue')
          }
        },
        {
          path: '/rule-view',
          name: 'ruleView',
          title: '告警规则预览',
          components: {
            'ruleView': () => import('@/views/RuleView.vue')
          }
        },
        {
          path: '/base-config',
          name: 'baseConfig',
          title: '基本配置',
          components: {
            'baseConfig': () => import('@/views/Config.vue')
          },
          children: [
            {
              path: '/empty',
              name: 'empty',
              title: '空状态',
              component: () => import('@/views/Empty.vue')
            },
            {
              path: '/base-labels',
              name: 'baseLabels',
              title: '公共标签',
              component: () => import('@/views/BaseLabels.vue')
            },
            {
              path: '/re-labels',
              name: 'reLabels',
              title: '标签重写',
              component: () => import('@/views/RelabelConfig.vue')
            },
            {
              path: '/default-jobs',
              name: 'defaultJobs',
              title: '默认分组',
              component: () => import('@/views/ContailAll.vue')
            },
            {
              path: '/check-yml',
              name: 'checkYml',
              title: '测试配置',
              component: () => import('@/views/CheckYml.vue')
            },
            {
              path: '/edit-prometheus-yml',
              name: 'editPrometheusYml',
              title: '模板编辑',
              component: () => import('@/views/Tmpl.vue')
            }
          ]
        }
      ]
    }
]

export default  createRouter({
  // 指定路由的模式,此处使用的是hash模式
  history: createWebHashHistory(),
  routes // short for `routes: routes`  
})

