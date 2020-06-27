<template>
  <div class="buttons">
    <v-toolbar dense dark color="blue darken-1" class="tool">
    
      <v-btn v-if=toRoot color="blue darken-4" fab small class="ma-1" @click="ToRoot()">
        <v-icon>mdi-home</v-icon>
      </v-btn>

      <v-btn v-if=toHome color="blue darken-4" fab small class="ma-1" @click="ToHome()">
        <v-icon>mdi-account</v-icon>
      </v-btn>

      <v-btn v-if=newArticle color="blue darken-4" fab small class="ma-1" @click="openDialog()">
        <v-icon>mdi-file-document</v-icon>
      </v-btn>

      <v-btn color="blue darken-4" fab small class="ma-1" @click="changeTheme()">
        <v-icon>mdi-theme-light-dark</v-icon>
      </v-btn>

      <v-btn v-if="!auth" color="blue darken-4" fab small class="ma-1" @click="openDialog()">
        <v-icon>mdi-login</v-icon>
      </v-btn>
      
      <v-btn v-else color="blue darken-4" fab small class="ma-1" @click="logout()">
        <v-icon>mdi-logout</v-icon>
      </v-btn>
    </v-toolbar>
  </div>
</template>

<script>

import { mapGetters } from 'vuex'
import axios from 'axios'

export default {
  name: 'Controls',

  props: {
    newArticle: Boolean,
    toHome : Boolean,
    toRoot: Boolean,
  },

  computed: mapGetters(['auth']),

  methods: {
    openDialog() {
      this.$emit('open')
    },

    changeTheme() {
      this.$store.commit('themeDark', !this.$store.getters.themeDark)
      this.$vuetify.theme.dark = this.$store.getters.themeDark
    },
    async logout() {
      var vueth = this
      await axios.post(this.$store.getters.backurl + '/auth/logout', {refrashToken: vueth.$store.getters.refrashToken})
      .then(() => {
         vueth.$router.push({name: 'Root'})
      })
      this.$store.commit('logout')

      this.$router.push({name: 'Root'})
    },

    ToRoot() {

      this.$router.push({name: 'Root'})
    },
    ToHome() {
      this.$router.push({name: 'Home'})
    },
    ToEditor() {
      this.$router.push({name: 'Editor'})
    },
  }
}
</script>

<style scoped>
.buttons {
  position: absolute;
  right: 0rem;
  z-index: 1000;
}
.tool {
  border-radius: 0px 0px 25px 25px;
}
</style>