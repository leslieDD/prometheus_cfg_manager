import { createStore } from 'vuex'
import { getUserInfo } from '@/utils/auth.js'
import { getStorageUserInfo } from '@/utils/localStorage.js'

const store = createStore({
  state() {
    return {
      userinfo: null,
    };
  },
  mutations: {
    setUserInfo(state, payload) {
      state.userinfo = payload;
    },
    resetToken(state) {
      state.userinfo.session.token = ''
    }
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
    }
  },
  actions: {
    setUserInfo(context, payload) {
      context.commit("setUserInfo", payload);
    },
    resetToken(context){
      context.commit("resetToken");
    }
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