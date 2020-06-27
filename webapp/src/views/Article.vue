<template>
  <v-content>
    <Controls toHome toRoot />
    <v-container fluid>
      <v-row justify="center">
        <v-col cols="12" md="8" sm="12" xs="12">
          <PreviewArticle :Article="viewArticle.body"/>
        </v-col>
      </v-row>
    </v-container>
    <!-- <v-container>
        <h1 class="text-center mt-12">Статус</h1>
        <h1 class="text-center" style="font-size:10rem">404</h1>
    </v-container> -->
  </v-content>
</template>

<script>
import PreviewArticle from '@/components/PreviewArticle.vue'
import Controls from '@/components/Controls.vue'
import { mapGetters } from 'vuex'

export default {
  name: 'Article',
  components: {
    PreviewArticle,
    Controls,
  },

  computed: mapGetters(['viewArticle', 'Status']),

  created() {
   this.$store.dispatch('getViewArticle', this.$route.params.id)
  }
}
</script>