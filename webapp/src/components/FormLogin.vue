<template>
  <v-content>
    <v-card>
      <v-tabs
        grow
        background-color="blue darken-2"
        dark
        color="blue lighten-5"
      >
        <v-tab>
          Вход
        </v-tab>
        <v-tab>
          Регистрация
        </v-tab>

        <v-tab-item class="pa-4">
          
          <v-card flat>
            <v-card-text>
              <form>
              <v-text-field
                v-model=loginRequest.email
                label= "E-mail"
                required
                :error-messages="loginErrors"
                @input="$v.loginRequest.email.$touch()"
                @blur="$v.loginRequest.email.$touch()"
                />
              <v-text-field
                v-model=loginRequest.password
                label="Пароль"
                required
                :error-messages="loginPasswordErrors"
                @input="$v.loginRequest.password.$touch()"
                @blur="$v.loginRequest.password.$touch()"
                type='password'
                />
              </form>
              <div>
                <h4><a href="#" @click="openForgot()">Забыли пароль?</a></h4>
              </div>
            </v-card-text>
            <v-card-actions>
              <v-btn color="blue" outlined @click="closeDialog()">Вернуться</v-btn>
                <v-spacer></v-spacer>
              <v-btn color="blue" outlined @click="login()">Войти</v-btn>
            </v-card-actions>
          </v-card>
        </v-tab-item>

        <v-tab-item class="pa-4">
          <v-card flat>
            <v-card-text>
              <v-text-field
                v-model=signupForm.email
                label= "E-mail"
                required
                :error-messages="signupEmailErrors"
                @input="$v.signupForm.email.$touch()"
                @blur="$v.signupForm.email.$touch()"
                />
              <v-text-field
                v-model=signupForm.password
                label="Пароль"
                required
                :error-messages="signupPasswordErrors"
                @input="$v.signupForm.password.$touch()"
                @blur="$v.signupForm.password.$touch()"
                :append-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showPass ? 'text' : 'password'"
                @click:append="showPass = !showPass"
                />
              <v-text-field
                v-model=signupForm.confirm
                label="Подтвердите пароль"
                required
                :error-messages="signupConfirmErrors"
                @input="$v.signupForm.confirm.$touch()"
                @blur="$v.signupForm.confirm.$touch()"
                type="password"
                />
            </v-card-text>
            <v-card-actions>
              <v-btn color="blue" outlined @click="closeDialog()">Вернуться</v-btn>
                <v-spacer></v-spacer>
              <v-btn color="blue" outlined @click="signUp()">Регистрация</v-btn>
            </v-card-actions>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-card>
    <v-dialog v-model="forgotDialog" max-width="600px">
      <v-card class="pa-2">
        <v-card-title>
          Востановление пароля
        </v-card-title>
        <v-card-text>
          <v-text-field
            v-model="forgotEmail"
            label="E-mail"
            required
            :error-messages="emailForgotErrors"
            @input="$v.forgotEmail.$touch()"
            @blur="$v.forgotEmail.$touch()"
          />
        </v-card-text>
        <v-card-actions>
            <v-btn color="blue darken-1" outlined @click="forgotDialog=false">Вернуться</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" outlined @click="forgot()">Продолжить</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
   
    <v-dialog v-model="showBaner" max-width="600px">
      <v-card class="pa-2">
        <v-card-title>
          {{baner.title}}
        </v-card-title>
        <v-card-subtitle>
          {{baner.message}}
        </v-card-subtitle>
        <v-card-actions>
            <v-btn color="blue darken-1" outlined @click="showBaner = false">Вернуться</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-content>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, maxLength, email, minLength } from 'vuelidate/lib/validators'
import axios from 'axios'

export default {
  name: 'FormLogin',

  data: () => ({
    showBaner: false,
    baner: {},
    showPass: false,
    forgotDialog: false,
    forgotEmail: '',
    loginRequest : {
        email: '',
        password: ''
    },
    signupForm: {
        email: '',
        password: '',
        confirm: ''
    },
  }),

  mixins: [validationMixin],
    validations: {
      loginRequest : {
        email: { required},
        password: { required}
      },
      signupForm: {
          email: { required, email },
          password: { required, maxLength: maxLength(100), minLength: minLength(8)},
          confirm: {
            required,
            compare(value) {
              return value == this.signupForm.password
            }
          }
        },
        forgotEmail: {required, email}
    },


  computed: {
    loginErrors () {
      const errors = []
      if (!this.$v.loginRequest.email.$dirty) return errors
      !this.$v.loginRequest.email.required && errors.push('Пустой e-mail')
      return errors
    },
    loginPasswordErrors() {
      const errors = []
      if (!this.$v.loginRequest.password.$dirty) return errors
      !this.$v.loginRequest.password.required && errors.push('Пустой пароль')
      return errors
    },

    signupEmailErrors () {
      const errors = []
      if (!this.$v.signupForm.email.$dirty) return errors
      !this.$v.signupForm.email.email && errors.push('Must be valid e-mail')
      !this.$v.signupForm.email.required && errors.push('E-mail is required')
      return errors
    },

    signupPasswordErrors () {
      const errors = []
      if (!this.$v.signupForm.password.$dirty) return errors
      !this.$v.signupForm.password.required && errors.push('Пустой пароль')
      !this.$v.signupForm.password.minLength && errors.push('Пароль слишком короткий, минимум 8 символов')
      !this.$v.signupForm.password.maxLength && errors.push('Пароль слишком длинный, максимум 100 символов')
      return errors
    },

    signupConfirmErrors () {
        const errors = []
        if (!this.$v.signupForm.confirm.$dirty) return errors
        !this.$v.signupForm.confirm.compare && errors.push('Пароли не совпадают')
        return errors
    },

    emailForgotErrors () {
        const errors = []
        if (!this.$v.forgotEmail.$dirty) return errors
        !this.$v.forgotEmail.email && errors.push('E-mail не валиден')
        !this.$v.forgotEmail.required && errors.push('Пустой e-mail')
        return errors
    },
  },

  methods: {
    openForgot() {
      this.forgotDialog = true
      this.$emit('close')
    },
    closeDialog() {
      this.$emit('close')
    },

    async login() {
      this.$v.loginRequest.$touch()
      if(this.$v.loginRequest.$invalid) {
          return
      }
      
      await axios.post(this.$store.getters.backurl + '/auth/login', {
        email: this.loginRequest.email,
        password: this.loginRequest.password         
      })
      .then(response => {
          this.$store.commit('refrash', response.data)
          this.$router.push({name: 'Home'})
      })
      .catch(function(e){
          console.log(e)
      })

    },
    async signUp() {
      this.$v.signupForm.$touch()
      if(this.$v.signupForm.$invalid) {
          return
      }

      await axios.post(this.$store.getters.backurl + '/auth/signup', {
          email: this.signupForm.email,
          password: this.signupForm.password         
      })
      .then(() => {
          this.forgotDialog = false
          this.$emit('close')
          this.baner = {
            title: 'Регистрация почти завершена!',
            message: 'Чтобы завершить регистрацию, подтвердите свой e-mail',
          }
          this.showBaner = true
      })
      .catch(function(e){
          console.log(e)
      })
    },
    async forgot() {
      this.$v.forgotEmail.$touch()
      if(this.$v.forgotEmail.$invalid) {
          return
      }
      await axios.post(this.$store.getters.backurl + '/auth/forgot', {
          email: this.forgotEmail,
      })
      .then(() => {
        this.forgotDialog = true
        this.$emit('close')
        this.baner = {
          title: 'Смена пароля',
          message: 'На ваш e-mail придёт письмо со ссылкой для продолжения процедуры смены пароля',
        }
        this.showBaner = true
      })
      .catch(function(e){
          console.log(e)
      })
    }
  },
}
</script>
