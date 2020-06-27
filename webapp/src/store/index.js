import Vue from 'vue'
import Vuex from 'vuex'

import root from './modules/root'
import section from './modules/section'
import home from './modules/home'
import editor from './modules/editor'
import user from './modules/user'
import article from './modules/article'


Vue.use(Vuex)



export default new Vuex.Store({
  state: {
    BACK_URL : "http://192.168.1.40:9090",
    THEME_DARK: false,
  },
  mutations: {
    initialiseStore(state) {
      if(localStorage.getItem('themeDark')) {
					state.THEME_DARK = JSON.parse(localStorage.getItem('themeDark'))
      
      }
      state.user.auth = window.$cookies.isKey("refrashToken")
    },
    themeDark  (state, event)  {
      state.THEME_DARK = event
      localStorage.setItem('themeDark', event)
      console.log(state.THEME_DARK)
    },
  },
  actions: {
  },
  getters: {
    themeDark(state) {
      return state.THEME_DARK
    },
    backurl(state) {
  
      return state.BACK_URL
    }
  },
  modules: {
    root,
    section,
    home,
    editor,
    user,
    article,
  },
})
