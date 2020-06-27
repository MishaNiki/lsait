<template>
  <v-content>

    <Controls toHome toRoot/>

    <v-container v-if="Status == 200">
      <v-row>
        <v-col class="text-center">
          <h1>{{Section.title}}</h1>
        </v-col>
      </v-row>
      <v-expansion-panels
        popout
        multiple
        accordion
      >
        <v-expansion-panel
          v-for="(item,i) in Section.themes"
          :key="i"
        >
          <v-expansion-panel-header style="font-size:1.3em">{{item.title}}</v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-list>
              <v-list-item
                v-for="(art,j) in item.articles"
                :key="j"
                @click="selectArt(art.id)"
              >
                  <v-list-item-content>
                    <v-list-item-title>{{art.title}}</v-list-item-title>
                  </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-expansion-panel-content>
        </v-expansion-panel>
        </v-expansion-panels>
    </v-container>
    <v-container v-else>
        <h1 class="text-center mt-12">Статус</h1>
        <h1 class="text-center" style="font-size:10rem">404</h1>
    </v-container>
  </v-content>
</template>

<script>
import Controls from '@/components/Controls.vue'
import { mapGetters } from 'vuex'


export default {
  name: 'Section',

  components: {
    Controls,
  },
  computed: mapGetters(['Section', 'Status']),
  methods: {
    selectArt(id) {
      this.$router.push({ name:'Article', params: {id:id}})
    }
  },
  created() {
    console.log(this.$route.params.uuid)
    this.$store.dispatch('getSection', this.$route.params.uuid)
  }
}
</script>
