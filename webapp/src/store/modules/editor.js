import ParseMD from '@/editor/ParseMD.js'
import axios from 'axios'

export default {
  state: {
    Article: [],
    ArticleMD: {},
    downArt: {},
    Image: [],
  },
  actions: {
    parseMD(ctx, md) {
      ctx.commit('updateArticle', ParseMD.ParseMD(md))
    },
    getImages() {

    },
    uploadImg() {

    },
    async getEditArticle(ctx, id) {
      await axios({
        url: ctx.getters.backurl + '/article/edit/' + id,
        method: 'get',
        headers: {
          'Token': ctx.getters.accessToken,
        },
        })
        .then(response => {
          console.log(response.data.article)
          ctx.commit('updateDownArticle', response.data.article)
        }) 
        .catch(err => {
          console.log(err.response.status)
          if(err.response.status == 426) {
            ctx.dispatch('refrash')
            ctx.dispatch('getEditArticle', id)
          } else if(err.response.status == 404) {
            console.log(404)
          } else {
            this.$store.commit('logout')
          }
        })
    },
    async putArticle(ctx) {
      var art = ctx.getters.ArticleMD
      await axios({
        url: ctx.getters.backurl + '/article/article',
        method: 'put',
        headers: {
          'Token': ctx.getters.accessToken,
        },
        data: {
          id: art.id,
          title: art.title != null ? art.title : "",
          description: art.description != null ? art.description : "",
          text: art.text != null ? art.text : "",
          idtheme: art.idtheme != null ? art.idtheme : 0,
        }
        })
        .then(() => {
          console.log('OK :)')
        }) 
        .catch(err => {
          console.log(err.response.status)
          if(err.response.status == 426) {
            ctx.dispatch('refrash')
            ctx.dispatch('putArticle')
          } else if(err.response.status == 404) {
            console.log(404)
          } else {
            this.$store.commit('logout')
          }
        })
    },

    async uploadArticle(ctx, info) {
      var art = ctx.getters.ArticleMD
      await axios({
        url: ctx.getters.backurl + '/article/article',
        method: 'post',
        headers: {
          'Token': ctx.getters.accessToken,
        },
        data: {
          id: art.id,
          title: info.title,
          description: info.description,
          text: art.text,
          idtheme: info.idTheme,
        }
        })
        .then(() => {
          console.log('OK :)')
        }) 
        .catch(err => {
          console.log(err.response.status)
          if(err.response.status == 426) {
            ctx.dispatch('refrash')
            ctx.dispatch('uploadArticle', info)
          } else if(err.response.status == 404) {
            console.log(404)
          } else {
            this.$store.commit('logout')
          }
        })
    }
  },
  mutations: {
    updateArticle(state, md) {
      state.ArticleMD.text = md
      if(state.ArticleMD.text != null) {
        state.Article = ParseMD.ParseMD(md)
      }
    },
    updateImages(state, images) {
      state.Images = images
    },
    updateDownArticle(state, art) {
      state.ArticleMD = art
      state.Article = ParseMD.ParseMD(art.text)
    },
    clearArticle(state) {
      state.Article = []
    }
  },
  getters: {
    Article(state) {
      return state.Article
    },
    Images(state) {
      return state.Images
    },
    TextArticle(state) {
      if(state.ArticleMD.text == null) {
        state.ArticleMD.text = ''
      }
      return state.ArticleMD.text
    },
    ArticleMD(state) {
      return state.ArticleMD
    }
  }
}