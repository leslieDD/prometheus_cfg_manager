import { createStore } from 'vuex'
import { getUserInfo } from '@/utils/auth.js'

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

loadUserInfoFromCookie()

export default store