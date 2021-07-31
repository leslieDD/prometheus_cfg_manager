import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
    {
      path: '/',
      name: 'login',
      title: '登录',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/person',
      name: 'person',
      title: '个人中心',
      component: () => import('@/views/Person.vue')
    },
    {
      path: '/register',
      name: 'register',
      title: '注册',
      component: () => import('@/views/Register.vue')
    },
    {
      path: '/menu',
      name: 'menu',
      // director: ''
      component: () => import('@/views/Menu.vue'),
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
          path: '/ip-batch-opt',
          name: 'BatchOpt',
          title: 'IP批量导入',
          components: {
            ipManager: () => import('@/views/BatchOpt.vue')
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
          path: '/labels-jobs',
          name: 'labelsJobs',
          title: '分组标签', // 为分组再分子组，再为子组挂上标签
          components:  {
            jobs: () => import('@/views/LabelsJobs.vue')
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
            },
            {
              path: '/options',
              name: 'options',
              title: '选项编辑',
              component: () => import('@/views/Options.vue')
            },
            {
              path: '/user',
              name: 'user',
              title: '用户管理',
              component: () => import('@/views/User.vue')
            },
            {
              path: '/group',
              name: 'group',
              title: '组管理',
              component: () => import('@/views/Group.vue')
            },
            {
              path: '/privileges',
              name: 'privileges',
              title: '权限管理',
              component: () => import('@/views/Privileges.vue')
            }
          ]
        }
      ]
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('@/views/Admin.vue'),
      children: [
        {
          path: '/user',
          name: 'user',
          title: '用户管理',
          component: () => import('@/views/User.vue')
        },
        {
          path: '/group',
          name: 'group',
          title: '组管理',
          component: () => import('@/views/Group.vue')
        },
        {
          path: '/privileges',
          name: 'privileges',
          title: '权限管理',
          component: () => import('@/views/Privileges.vue')
        }
      ]
    },
    {
      path: "/:catchAll(.*)",
      name: '404',
      title: '404',
      component: () => import('@/views/404.vue')
      // redirect: {name: 'login'}
    }
]

export default  createRouter({
  // 指定路由的模式,此处使用的是hash模式
  history: createWebHashHistory(),
  routes // short for `routes: routes`  
})

