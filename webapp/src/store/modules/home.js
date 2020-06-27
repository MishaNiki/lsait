import axios from 'axios'

export default {
  state: {
    home: {},
    EmptyProfile: false,
  },
  actions: {
    async getHome(ctx) {
      await axios({
        url: ctx.getters.backurl + '/article/profile',
        method: 'get',
        headers: {
          'Token': ctx.getters.accessToken,
        }
      })
     .then(response => {
        console.log(response.data.profile)
        ctx.commit('upadateHome', response.data.profile)
      }) 
     .catch(err => {
        console.log(err.response.status)
        if(err.response.status == 426) {
          ctx.dispatch('refrash')
          ctx.dispatch('getHome')
        } else {
          ctx.commit('logout')
        }
     })
    },
    
    async putProfile(ctx, profile) {
      await axios({
        url: ctx.getters.backurl + '/article/profile',
        method: 'put',
        headers: {
          'Token': ctx.getters.accessToken,
        },
        data: {
          name: profile.name,
          surname: profile.surname,
          position: profile.position,
          description: profile.description
        }
      })
     .then(() => {
        ctx.dispatch('getHome')
      }) 
     .catch(err => {
        if(err.response.status == 426) {
          ctx.dispatch('refrash')
          ctx.dispatch('putProfile')
        } else {
          ctx.commit('logout')
        }
     })
    },
    async deleteArticle(ctx, id) {
      await axios({
        url: ctx.getters.backurl + '/article/article',
        method: 'delete',
        headers: {
          'Token': ctx.getters.accessToken,
        },
        data: {
          id: id
        }
      })
     .then(() => {
        ctx.dispatch('getHome')
      }) 
     .catch(err => {
        if(err.response.status == 426) {
          ctx.dispatch('refrash')
          ctx.dispatch('deleteArticle')
        } else {
          ctx.commit('logout')
        }
     })
    }
  },
  mutations: {
    upadateHome(state, profile) {
      state.home = profile
    },
    updateEmptyProfile(state, empr) {
      state.EmptyProfile = empr
    }
  },
  getters: {
    Home(state) {
      return state.home
    },
    ArticleHome(state) {
      if(state.home.articles) {
        return state.home.articles
      } else {
        return []
      }
    },
    DraftHome(state) {
      if(state.home.drafts != null) {
        return state.home.drafts
      } else {
        return []
      }
    },
    EmptyProfile(state) {
      return state.EmptyProfile
    }
  }
}