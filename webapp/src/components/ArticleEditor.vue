<template>
  <v-content class="pa-md-8">
    <v-row>
      <v-col cols="12" md="6" sm="12" xs="12">
      <v-bottom-navigation background-color="blue darken-2" dark>
        
        <v-btn @click="openImg = true">
          <span>Изображение</span>
          <v-icon>mdi-image</v-icon>
        </v-btn>

        <v-btn @click="saveArticle()">
          <span>Сохранить</span>
          <v-icon>mdi-content-save</v-icon>
        </v-btn>

        <v-btn @click="OpenUploadArt()">
          <span>Выложить</span>
          <v-icon>mdi-upload</v-icon>
        </v-btn>

        <v-btn>
          <span>Помощь</span>
          <v-icon>mdi-help-circle</v-icon>
        </v-btn>
      </v-bottom-navigation>

        <v-textarea
          :value="TextArticle"
          @input="preview"
          auto-grow
          outlined
          rows='20'
          style="font-family: 'Hack';"
          v-on:keydown.tab.prevent="tabber($event)"
        />
      </v-col>
      <v-col cols="12" md="6" sm="12" xs="12">
        <PreviewArticle :Article="Article"/>
      </v-col>
    </v-row>

    <v-dialog v-model=openImg max-width="600px">
      <v-card>
        <v-card-title>
          Выбрать изображение
        </v-card-title>
        <v-container fluid>
        <v-row justify="center" class="pa-4">
          <v-file-input 
            accept="image/png, image/jpeg, image/bmp"
            prepend-icon="mdi-image"
            label="Загрузить изобpажение"
            :hide-input="true"
          ></v-file-input>
        </v-row>
        <v-row justify="center">
           <v-col
              v-for="(img, ind) in Images"
              :key="ind"
              class="d-flex child-flex"
              cols="4"
            >
             <v-img
              :src="img"
              aspect-ratio="1"
              class="grey lighten-2"
            >
             </v-img>
           </v-col>
        </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    
    <v-dialog v-model=uploadArt max-width="600px">
      <v-card>
        <v-card-title>
          Выложить статью
        </v-card-title>
        <v-container fluid class="pa-8">
          <v-text-field
            v-model="uploadArticle.title"
            outlined
            label="Название статьи"
          />
          <v-textarea
            v-model="uploadArticle.description"
            outlined
            name="input-7-4"
            label="Описание"
          />
          <v-select
            v-model="uploadArticle.Section"
            :items="ListSectionsAndThemes"
            item-text="title"
            label="Выберите раздел"
            outlined
            return-object
          >
          
          </v-select>

          <v-select
            v-model="uploadArticle.idTheme"
            :items="uploadArticle.Section.themes "
            item-text="title"
            item-value="id"
            label="Выберите тему"
            outlined
          ></v-select>
        </v-container>
        <v-card-actions>
          <v-btn text @click="uploadArt = false">
            Отменить
          </v-btn>
          <v-spacer/>
          <v-btn text @click="upArt()">
            Выложить
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-content>
</template>

<script>
import PreviewArticle from '@/components/PreviewArticle.vue'
import { mapGetters} from 'vuex'

export default {
  name: 'ArticleEditor',

  components: {
    PreviewArticle,
  },

  methods : {
    preview(e) {
      console.log(e)
      this.$store.commit('updateArticle', e)
    },

    tabber(event){
        let text = event.target.value,
          originalSelectionStart = event.target.selectionStart,
          textStart = text.slice(0, originalSelectionStart),
          textEnd =  text.slice(originalSelectionStart);

      this.text = `${textStart}\t${textEnd}`
      event.target.value = this.text // required to make the cursor stay in place.
      event.target.selectionEnd = event.target.selectionStart = originalSelectionStart + 1
    },

    saveArticle() {
      this.$store.dispatch('putArticle')
    },

    OpenUploadArt(){
      this.uploadArt = true
      this.uploadArticle.title = this.$store.getters.ArticleMD.title
      this.uploadArticle.description = this.$store.getters.ArticleMD.description
    },
    
    upArt() {
      this.$store.dispatch('uploadArticle', this.uploadArticle)
    }
  },

  computed: mapGetters(['Images', 'TextArticle', 'ListSectionsAndThemes', 'Article']),

  data: () => ({
    text: '',
    openImg: false,
    uploadArt: false,
    uploadArticle: {
      Section: {},
      idTheme: 0,
      title: '',
      description: '',
    },
  }),

//   watch: {
//     updateTextArticle(text) {
//       this.$nextTick(() => {
//         this.$store.commit('updateArticle', text)
//       })
//     }
//   }


  created() {
    this.$store.dispatch('getListSectionsAndThemes')
  }
}
</script>