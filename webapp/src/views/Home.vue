<template>

  <v-content>

    <Controls @open="newDraftDial=true" newArticle toRoot/>

    <v-container class="pt-10">
      <v-row
        justify="center"
        class="pa-2"
      >
      
      <Profile 
        :name=Home.name
        :surname=Home.surname
        :position=Home.position
        :description=Home.description
      />

      <ListArticles/>
     
      </v-row>

      <v-dialog v-model=banerProfile persistent max-width="600px">
        <v-card>
          <v-card-title>
            Заполните свой профиль
          </v-card-title>
          <v-card-text>
            <v-text-field
              v-model=profile.name
              label="Имя"
              required
              :error-messages="putNameError"
              @input="$v.profile.name.$touch()"
              @blur="$v.profile.name.$touch()"
            />
            <v-text-field
              v-model=profile.surname
              label="Фамилия"
              required
              :error-messages="putSurnameError"
              @input="$v.profile.surname.$touch()"
              @blur="$v.profile.surname.$touch()"
            />
            <v-text-field
              v-model=profile.position
              label="Должность"
              required
            />
            <v-textarea
              v-model=profile.description            
              outlined
              name="input-7-4"
              label="Дополнительное описание"
            />
          </v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn color="blue darken-1" text @click="putProfile()">Продолжить</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog v-model="newDraftDial" max-width="600px">
        <v-card>
          <v-card-title>Новая статья</v-card-title>
          <v-card-subtitle>Будет созданно как черновик</v-card-subtitle>
          <v-card-text>
            <v-text-field
              v-model=newArticle.title
              label="Название"
              required
            />
            <v-textarea
              v-model=newArticle.description
              outlined
              name="input-7-4"
              label="Описание"
            />
          </v-card-text>
          <v-card-actions>
            <v-btn text @click="newDraftDial = false">Закрыть</v-btn>
            <v-spacer/>
            <v-btn text @click="CreateArticle()">Создать</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-container>
  </v-content>
</template>

<script>
import Profile from '@/components/Profile.vue'
import ListArticles from '@/components/ListArticles.vue'
import Controls from '@/components/Controls.vue'
import { mapGetters } from 'vuex'

import { validationMixin } from 'vuelidate'
import { required } from 'vuelidate/lib/validators'

import axios from 'axios'

export default {
  name: 'Home',


  mixins: [validationMixin],
  validations: {
    profile : {
      name: { required},
      surname: { required}
    }
  },

  data: () => ({
    banerProfile: false,
    newDraftDial: false,
    profile: {
      name: '',
      surname: '',
      position: '',
      description: '',
    },
    newArticle :{
      title: '',
      description: '',
    }
  }),

  components: {
    Profile,
    ListArticles,
    Controls
  },

  computed:{
    putNameError() {
      const errors = []
      if (!this.$v.profile.name.$dirty) return errors
      !this.$v.profile.name.required && errors.push('Путое имя')
      return errors
    },
    putSurnameError() {
      const errors = []
      if (!this.$v.profile.surname.$dirty) return errors
      !this.$v.profile.surname.required && errors.push('Пустая фамилия')
      return errors
    },
    ...mapGetters(['Home']),
  },

  methods: {
    putProfile() {
      this.$v.profile.$touch()
      if(this.$v.profile.$invalid) {
          return
      }
      this.$store.dispatch('putProfile', this.profile)
      this.$store.commit('updateEmptyProfile', false)
      this.banerProfile = false
    },

    async CreateArticle() {
       await axios({
        url: this.$store.getters.backurl + '/article/draft',
        method: 'post',
        headers: {
          'Token': this.$store.getters.accessToken,
        },
        data: {
          title: this.newArticle.title,
          description: this.newArticle.description
        }
        })
        .then(response => {
          console.log(response.data.id)
          this.$router.push({ name:'Editor', params: {id:response.data.id}})
        }) 
        .catch(err => {
          if(err.response.status == 426) {
            this.$store.dispatch('refrash')
            this.CreateArticle()
          } else {
            this.$store.commit('logout')
          }
        })
    },
  },

  created() {
    if(!this.$store.getters.auth) {
      this.$router.push({name: 'Root'})
      return
    }
    this.$store.dispatch('getHome')
    

    if(this.$store.getters.EmptyProfile) {
      this.banerProfile = true
    } else {
      this.banerProfile = false
    }
  }
}
</script>


<style scoped>
.buttons {
  position: absolute;
  right: 1rem;
}
</style>