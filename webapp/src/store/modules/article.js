import ParseMD from '@/editor/ParseMD.js'
import axios from 'axios'

export default {
  state: {
    viewArticle: {
      title: '',
      description: '',
      body: [],
    }
  },
  actions: {
    getViewArticle(ctx, id) {
        axios.get(ctx.getters.backurl + '/article/article/' + id)
        .then(response => {
          ctx.commit('updateViewArticle', response.data.article)
        })
    }
  },
  mutations: {
    updateViewArticle(state, article) {
      state.viewArticle.title = article.title
      state.viewArticle.description = article.description == null ? '' : article.description 
      state.viewArticle.body = ParseMD.ParseMD(article.text)
    }
  },
  getters: {
    viewArticle(state) {
        return state.viewArticle
    }
  }
}