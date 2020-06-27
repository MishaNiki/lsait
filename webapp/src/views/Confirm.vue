<template>
    <v-content>
    </v-content>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Confirm',

  async created() {
    var vueth = this
    const key = this.$route.query.key
    if(key != null) {
      console.log(key)
      await axios.post(this.$store.getters.backurl + '/auth/signup/confirm', {key: key})
      .then(response => {
        vueth.$store.commit('refrash', response.data)
        vueth.$store.commit('updateEmptyProfile', true)
        vueth.$router.push({name: 'Home'})
      })
      .catch(function(e){
        vueth.$router.push('/')
        console.log(e)
      })
    } else {
      this.$router.push({name: 'Root'})
    }
  }
}
</script>