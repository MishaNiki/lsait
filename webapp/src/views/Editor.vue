<template>
  <v-content>
    <Controls @open="dialog=true" toHome toRoot />
    <v-container fluid class="pt-10">
      <ArticleEditor/>
    </v-container>
  </v-content>
</template>

<script>
import ArticleEditor from '@/components/ArticleEditor.vue'
import Controls from '@/components/Controls.vue'

export default {
  name: 'Editor',

  components: {
    ArticleEditor,
    Controls,
  },

  methods: {
  },

  created() {
    this.$store.commit('clearArticle')
    if(!this.$store.getters.auth) {
      this.$router.push({name: 'Root'})
      return
    }
    console.log(this.$route.params.id)
     this.$store.dispatch('getEditArticle', this.$route.params.id)
  }
}
</script>