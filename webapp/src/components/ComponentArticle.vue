<template>
  <v-col>
    <div v-if="name==='paragraph'" v-html="value">
    </div>

    <div v-else-if="name==='header'">
      <h1 v-if="size===1" v-html="value"></h1>
      <hr v-if="size===1">
      <h2 v-else-if="size===2" v-html="value"></h2>
      <h3 v-else-if="size===3" v-html="value"></h3>
      <h4 v-else-if="size===4" v-html="value"></h4>
      <h5 v-else-if="size===5" v-html="value"></h5>
      <h6 v-else-if="size===6" v-html="value"></h6>
    </div>

    <div v-else-if="name=='marklist'">
      <ul>
        <li v-for="(item, index) in list"  :key=index>
          <span v-html="item.value"></span>
          <ul>
            <li v-for="(item1, index) in item.list"  :key=index>
              <span v-html="item1.value"></span>
              <ul>
                <li v-for="(item2, index) in item1.list"  :key=index>
                  <span v-html="item2.value"></span>
                </li>
              </ul>
            </li>
          </ul>    
        </li>
      </ul>
    </div>

    <div v-else-if="name=='numlist'">
      <ol>
        <li v-for="(item, index) in list"  :key=index>
          <span v-html="item.value"></span>
          <ol>
            <li v-for="(item1, index) in item.list" :key=index>
              <span v-html="item1.value"></span>
              <ol>
                <li v-for="(item2, index) in item1.list" :key=index>
                  <span v-html="item2.value"></span>
                </li>
              </ol>
            </li>
          </ol>    
        </li>
      </ol>
    </div>

    <div v-else-if="name=='quote'">
      <v-alert
        border="left"
        colored-border
        color="deep-purple accent-4"
        elevation="2"
        class="font-italic font-weight-medium" style="font-size: 1em"
      >
        <div v-for="(item, index) in list"  :key=index v-html="item.value">
        </div>
      </v-alert>
    </div>

    <div v-else-if="name==='line'">
      <hr>
    </div>

    <div v-else-if="name==='code'">
      <v-card>
        <v-card-text
          class="code"
          :style='this.$store.getters.themeDark?"color: #FFFFFF":"color: #000000"'
        >
          <div
            v-for="(str, index) in list"
            :key=index
          >
            {{str}}
          </div>
        </v-card-text>
      </v-card>
    </div>

    <div v-else-if="name=='image'" class="pa-1">
        <v-img :alt='alt' :src="value"/>
    </div>

    <div v-else-if="name=='imagelink'" class="pa-1">
      <a :href="link">
        <v-img :alt='alt' :src="value"/>
      </a>
    </div>

    <div v-else-if="name=='images'">
    </div>

     <div v-else-if="name=='table'">
       <v-simple-table>
        <template v-slot:default>
          <thead>
            <tr>
              <th
                v-for="(item, index) in table.header"
                :key=index
                :class="item.align"
                style="font-size:1em"
              >{{item.text}}</th>
            </tr>
          </thead>
          <tbody> 
            <tr
                v-for="(item, index) in table.items"
                :key=index
              >
                <td
                  v-for="(it, ind) in item"
                  :key=ind
                  :class="table.header[ind].align"
                  style="font-size:1em"
                  v-html="it"
                ></td>
            </tr>
          </tbody>
        </template>
       </v-simple-table>
    </div>

  </v-col>
</template>

<script>
import 'katex/dist/katex.min.css'

export default {
  name: 'ComponentArticle',

  props: {
    name: String,
    value: String,
    size: Number,
    lang: String,
    list: Array,
    alt: String,
    link: String,
    table: Object
  },

  methods: {
  }
  
}
</script>

<style>

@font-face {
  font-family: 'Hack';
  src: url('~@/assets/fonts/Hack/Hack-Regular.ttf') format('truetype');
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('~@/assets/fonts/Hack/Hack-Bold.ttf') format('truetype');
  font-weight: bold;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('~@/assets/fonts/Hack/Hack-Italic.ttf') format('truetype');
  font-weight: normal;
  font-style: italic;
}

@font-face {
  font-family: 'Hack';
  src: url('~@/assets/fonts/Hack/Hack-BoldItalic.ttf') format('truetype');
  font-weight: bold;
  font-style: italic;
}

</style>

<style scoped>

.code {
  font-family: 'Hack';
  font-weight: 800;
  white-space: pre-wrap;
}

</style>