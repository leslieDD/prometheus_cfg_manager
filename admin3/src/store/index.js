import { createStore } from 'vuex'
import { getUserInfo } from '@/utils/auth.js'
import { getStorageUserInfo } from '@/utils/localStorage.js'

const store = createStore({
  state() {
    return {
      userinfo: null,
      pageNum: 1,
      pageSize: 15,
    };
  },
  mutations: {
    setUserInfo(state, payload) {
      state.userinfo = payload;
    },
    resetToken(state) {
      state.userinfo.session.token = ''
    },
    setPageInfo(state, pageSize, pageNum) {
      state.pageSize = pageSize;
      state.pageNum = pageNum;
    },
  },
  getters: {
    token(state) {
      if (!state.userinfo) {
        return ''
      }
      return state.userinfo.session.token
    },
    userInfo(state) {
      return state.userinfo
    },
    userID(state) {
      return state.userinfo.id
    },
    niceName(state) {
      return state.userinfo.nice_name
    },
    getPageInfo(state) {
      let info = {}
      if(state.pageNum === 0) {
        info.pageNum = 1
      }
      info.pageNum = state.pageNum
      if(state.pageSize === 0) {
        info.pageSize = 1
      }
      info.pageSize = state.pageSize
      return info
    },
  },
  actions: {
    setUserInfo(context, payload) {
      context.commit("setUserInfo", payload);
    },
    resetToken(context){
      context.commit("resetToken");
    },
    setPageInfo(context, pageSize, pageNum) {
      context.commit("setPageInfo", pageSize, pageNum);
    },
  },
});

function loadUserInfoFromCookie() {
  const userInfo = getUserInfo()
  if (userInfo) {
    store.dispatch('setUserInfo', userInfo)
  }
}

function loadUserInfoFromLocalStorage() {
  const userInfo = getStorageUserInfo()
  if (userInfo) {
    store.dispatch('setUserInfo', userInfo)
  }
}

// loadUserInfoFromCookie()
loadUserInfoFromLocalStorage()

export default store