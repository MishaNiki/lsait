import axios from 'axios'

export default {
  state: {
    auth: false,
    accessToken: '',
    refrashToken: '',
  },
  actions: {
    async refrash(ctx) {
      await axios.post(ctx.getters.backurl + '/auth/refrash', {
          refrashToken: ctx.getters.refrashToken,
      })
      .then(response => {
        ctx.commit('refrash', response.data)
      })
      .catch(function(e){
          console.log(e)
      })
    }
  },
  mutations: {
    refrash(state, couple) {
      console.log(couple)
      state.auth = true
      state.accessToken = couple.accessToken
      state.refrashToken = couple.refrashToken
      window.$cookies.set("accessToken", couple.accessToken,"30MIN")
      window.$cookies.set("refrashToken", couple.refrashToken,"60D")
    },
    auth(state, au) {
      state.auth = au
    },
    logout(state) {
      state.auth = false
      window.$cookies.remove("accessToken")
      window.$cookies.remove("refrashToken")
    }
  },
  getters: {
    accessToken(state) {
      state.accessToken = window.$cookies.get("accessToken")
      return state.accessToken
    },
    refrashToken(state) {
      state.refrashToken = window.$cookies.get("refrashToken")
      return state.refrashToken
    },
    auth(state) {
      state.auth = window.$cookies.isKey("refrashToken")
      return state.auth
    }
  }
}
