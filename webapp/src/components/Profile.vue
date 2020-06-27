<template>
  <v-col
    cols="12"
    md="7"
    sm="12"
    xs="12"
    >
     <v-card
      class="overflow-hidden"
      >
      <v-toolbar
        flat
        color="blue darken-1"
        dark
        >
        <v-icon>mdi-account</v-icon>
        <v-toolbar-title class="font-weight-bold">{{surname}} {{name}}</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn
          color="blue darken-4"
          fab
          small
          @click="editProf()"
        >
          <v-icon v-if="editProfile">mdi-close</v-icon>
          <v-icon v-else>mdi-pencil</v-icon>
        </v-btn>
      </v-toolbar>

      <v-card-text v-if="!editProfile">
          <v-col cols="12">
            <h3>Должность:</h3> {{ position }}
          </v-col>
          <v-col cols="12">
            <h3>Дополнительное описание:</h3> {{ description }}
          </v-col>
      </v-card-text>
      <v-card-text v-else>
        <v-row justify="center">
          <v-col cols='12' style="margin-bottom: -25px">
          <v-text-field
            v-model="profile.name"
            dense
            outlined
            persistent-hint
            label="Имя"
            />
          </v-col>

          <v-col cols='12' style="margin-bottom: -25px">
          <v-text-field
            v-model="profile.surname"
            dense
            outlined
            persistent-hint
            label="Фамилия"
            />
          </v-col>

          <v-col cols='12' style="margin-bottom: -25px">
          <v-text-field
            v-model="profile.position"
            dense
            outlined
            persistent-hint
            label="Должность"
            />
          </v-col>
        
          <v-col cols='12' style="margin-bottom: -25px">
            <v-textarea
              v-model="profile.description"
              outlined
              name="input-7-4"
              label="Дополнительное описание"
            />
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions v-if="editProfile">
        <v-spacer/>
        <v-btn class="mx-2" color="blue darken-1" dark @click="updateProfile()">
          Сохранить
        </v-btn>
      </v-card-actions>
     </v-card>
  </v-col>
</template>

<script>
export default {
  name: 'Profile',

  props: {
    name: String,
    surname: String,
    position: String,
    description: String,
  },

  data: () =>({
    editProfile: false,
    profile: {
      name: '',
      surname: '',
      position: '',
      description: '',
    }
  }),

  methods: {
    updateProfile() {
      this.$store.dispatch('putProfile', this.profile)
      this.editProfile = false
    },
    editProf() {
      this.editProfile = !this.editProfile 
      if (this.editProfile) {
        this.profile.name = this.name
        this.profile.surname = this.surname
        this.profile.description = this.description
        this.profile.position = this.position
      }
    }
  },
}
</script>