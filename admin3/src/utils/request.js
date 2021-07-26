import axios from 'axios'
// import { ElMessageBox } from 'element-plus'
import {ElNotification} from 'element-plus'
import store from '@/store/index.js'
// import route from '@/router/index.js'
import { getToken, removeToken } from '@/utils/auth.js'

// create an axios instance
const service = axios.create({
  // baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 1000 * 8// request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers.Authorization = store.getters.token;
    }
    if (getToken()) {
      config.headers.Authorization = getToken()
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== 200000) {
      if (res.code === 401000) {
        ElNotification({
          title: '错误',
          message: res.msg || 'Error',
          type: 'error'
        })
        store.dispatch('resetToken')
        removeToken()
        location.reload()
      } else {
        ElNotification({
          title: '错误',
          message: res.msg || 'Error',
          type: 'error'
        })
      }
      return Promise.reject(new Error(res.msg || 'Error'))
    } else {
      return res
    }
  },
  error => {
    ElNotification({
      title: '错误',
      message: error.message,
      type: 'error'
    })
    if (error.response.status === 401) {
      store.dispatch('resetToken')
      removeToken()
      location.reload()
    }
    if (error.response.data && error.response.data.code === "401000") {
      ElNotification({
        title: '错误',
        message: res.msg || 'Error',
        type: 'error'
      })
      store.dispatch('resetToken')
      removeToken()
      location.reload()
    }
    return Promise.reject(error)
  }
)

export default service
