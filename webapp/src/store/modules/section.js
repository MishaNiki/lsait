import axios from 'axios'

export default {
  state: {
    Section: {},
    Status: 200,
    ListSectionsAndThemes: [],
  },
  actions: {
    async getSection(ctx, uuid) {
      axios.get(ctx.getters.backurl + "/article/section/" + uuid)
        .then(response => {
          console.log(response.data.section)
          ctx.commit('updateSection', response.data.section)
          ctx.commit('updateStatus', 200)
        })
        .catch(() => {
          ctx.commit('updateStatus', 404)
        });
    },
    async getListSectionsAndThemes(ctx) {
      axios.get(ctx.getters.backurl + "/article/section/theme")
        .then(response => {
          console.log(response.data.sections)
          ctx.commit('updateListSectionsAndThemes', response.data.sections)
        })
    }
  },
  mutations: {
    updateSection(state, section) {
      state.Section = section
    },
    updateStatus(state, status) {
      state.Status = status
    },
    updateListSectionsAndThemes(state, sections) {
      state.ListSectionsAndThemes = sections
    }
  },
  getters: {
    Section(state) {
      return state.Section
    },
    Status(state) {
      return state.Status
    },
    ListSectionsAndThemes(state) {
      return state.ListSectionsAndThemes
    }
  }
}