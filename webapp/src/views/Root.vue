<template>
  <v-content>
    <Controls @open="dialog=true" :toHome="auth"/>

    <v-container fluid>
      <v-row
        justify="center"
        class="mt-8"
      >
        <v-col
          cols="12"
          sm="9"
          md="5"
        >
          <v-text-field
            rounded
            solo
            label="Поиск"
            prepend-inner-icon='mdi-magnify'
          >
          </v-text-field>
        </v-col>
      </v-row>

      <!-- РАЗДЕЛЫ -->
      <v-row
        justify="center"
        class="pa-1"
      >
      <v-col cols="12" md="9" sm="12" xs="12">
        <v-row>
          <SectionBtn
            v-for="section in Sections"
            :key="section.id"
            :uuid="section.uuid"
            :title="section.title"
            @selectsec="selectSection(section.uuid)"
         />
        </v-row>
      </v-col>
      </v-row>

      <v-row>
        <v-dialog v-model="dialog" max-width="460px">
          <FormLogin @close='dialog=false'/>
        </v-dialog>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>

// @ is an alias to /src
import SectionBtn from '@/components/SectionBtn.vue'
import FormLogin from '@/components/FormLogin.vue'
import Controls from '@/components/Controls.vue'
import { mapGetters } from 'vuex'

export default {
  name: 'Root',

  components: {
    SectionBtn,
    FormLogin,
    Controls,
  },
  computed: mapGetters(['Sections', 'auth']),
  data: () => ({dialog: false}),
  methods: {
    search() {
    },
    selectSection(uuid) {
      this.$router.push({name: 'Section', params: {uuid:uuid}})
    }
  },
  created() {
    this.$store.dispatch('getSections')
  }
}
</script>