<template>
  <v-container
    class="fill-height"
    fluid
  >
  <v-row
    align="center"
    justify="center"
  >
    <v-col
      cols="12"
      sm="8"
      md="4"
    >
    <v-card max-width="600px">
      <v-card-title>
        Смена пароля
      </v-card-title>
      <v-card-text>
        <v-text-field
          v-model=forgotPass.password
          label="Новый пароль"
          required
          :error-messages="forgotPasswordErrors"
          @input="$v.forgotPass.password.$touch()"
          @blur="$v.forgotPass.password.$touch()"
          :append-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
          :type="showPass ? 'text' : 'password'"
          @click:append="showPass = !showPass"
        />
        <v-text-field
          v-model=forgotPass.confirm
          required
          label="Подтвердите пароль"
          :error-messages="forgotConfirmErrors"
          @input="$v.forgotPass.confirm.$touch()"
          @blur="$v.forgotPass.confirm.$touch()"
          type="password"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer/>
        <v-btn color="blue" outlined @click="forgot()">Подтвердить</v-btn>
      </v-card-actions>
    </v-card>
    </v-col>
  </v-row>
  </v-container>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'
import axios from 'axios'

export default {
  name: 'Forgot',
  data:() => ({
    showPass: false,
    forgotPass: {
      password: '',
      confirm: '',
    },
    key :'',
  }),

  mixins: [validationMixin],
  validations: {
    forgotPass: {
      password: { required, maxLength: maxLength(100), minLength: minLength(8)},
      confirm: {
        required,
        compare(value) {
          return value == this.signupForm.password
        }
      }
    },
  },

  methods: {
    async forgot() {
      var vueth = this
      await axios.post(this.$store.getters.backurl + '/auth/forgot/confirm', {key: vueth.key, password: vueth.forgotPass.password})
      .then(() => {
        vueth.$router.push({name: 'Root'})
      })
      .catch(function(e){
        vueth.$router.push({name: 'Root'})
        console.log(e)
      })
    }
  },

  computed: {
    forgotPasswordErrors() {
      const errors = []
      if (!this.$v.forgotPass.password.$dirty) return errors
      !this.$v.forgotPass.password.required && errors.push('Пустой пароль')
      !this.$v.forgotPass.password.minLength && errors.push('Пароль слишком короткий, минимум 8 символов')
      !this.$v.forgotPass.password.maxLength && errors.push('Пароль слишком длинный, максимум 100 символов')
      return errors
    },
    forgotConfirmErrors() {
      const errors = []
      if (!this.$v.forgotPass.confirm.$dirty) return errors
      !this.$v.forgotPass.confirm.compare && errors.push('Пароли не совпадают')
      return errors
    },
  },

  created() {
    const key = this.$route.query.key
    if(key != null) {
      this.key = this.$route.query.key
    } else {
      this.$router.push({name: 'Root'})
    }
  }

}
</script>