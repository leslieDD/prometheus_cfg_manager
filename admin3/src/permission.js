import router from './router/index.js'
import store from './store/index.js'

const whiteList = ['/', '/register'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
  if (store.getters.token) {
    next()
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next({name: 'login'})
    }
  }
})
