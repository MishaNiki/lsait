<template>
  <v-content>
    <ComponentArticle
      v-for="(ca, index) in Article"
      :key="index"
      :name=ca.name
      :value=ca.value
      :size=ca.size
      :lang=ca.lang
      :list=ca.list
      :alt=ca.alt
      :link=ca.link
      :table=ca.table
    />
  </v-content>
</template>

<script>
import ComponentArticle from '@/components/ComponentArticle.vue'

export default {
  name: 'PreviewArticle',
  
  components: {
      ComponentArticle,
  },
  props: {
    Article: Array,
  },
}
</script>