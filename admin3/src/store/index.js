import { createStore } from 'vuex'

export default createStore({
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
