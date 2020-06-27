import axios from 'axios'

export default {
  state: {
    Sections: [],
  },
  actions: {
    async getSections(ctx) {
      axios.get(ctx.getters.backurl + '/article/section')
        .then(response => {
          console.log(response.data.sections)
          ctx.commit('updateSection', response.data.sections)
        })
        .catch(error => {
          console.log(error)
        });
    }
  },
  mutations: {
    updateSection(state, sections) {
      state.Sections = sections
    },
  },
  getters: {
    Sections(state) {
      return state.Sections
    }
  }
}
