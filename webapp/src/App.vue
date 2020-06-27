<template>
  <v-app>
    <router-view/>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  created() {
    this.$store.commit('initialiseStore')
    this.$vuetify.theme.dark = this.$store.getters.themeDark
  }
}
</script>