<template>
    <v-col
      cols="12"
      md="7"
      sm="12"
      xs="12"
      >
        <v-tabs
          grow
        >
        <v-tab>
          Ваши статьи
        </v-tab>
        <v-tab>
          Черновики
        </v-tab>
        <v-tab>
          Закладки
        </v-tab>
        
        <v-tab-item>

          <v-row v-if="ArticleHome.length == 0" justify="center">
            <p class="mt-8">Пока что у вас нет статей</p>
          </v-row>
          <v-list
            three-line
          >
            <v-list-item
              v-for="art in ArticleHome"
              :key="art.id"
              @click="click()"
            >
              <v-list-item-content @click="selectArt(art.id)" style="cursor: pointer;">
                <v-list-item-title v-text="art.title"></v-list-item-title>
                <v-list-item-subtitle v-text="art.description"></v-list-item-subtitle>

                <v-list-item-subtitle class="text--secondary text-end">
                  Последнее изменение: {{ art.date }}
                </v-list-item-subtitle>
                <v-divider></v-divider>
              </v-list-item-content>

                <!-- <v-list-item-action>
                <v-btn icon>
                  <v-icon>mdi-comment</v-icon>
                  0
                </v-btn>
              </v-list-item-action> -->

            </v-list-item>
          </v-list>
        </v-tab-item>

        <v-tab-item>
          <v-row v-if="DraftHome == 0" justify="center">
            <p class="ma-8">Пока что у вас нет черновиков</p>
          </v-row>
          <v-list
            three-line
          >
            <v-list-item
              v-for="art in DraftHome"
              :key="art.id"
              @click="click()"
            >
              <v-list-item-content @click="selectArt(art.id)" style="cursor: pointer;">
                <v-list-item-title v-text="art.title"></v-list-item-title>
                <v-list-item-subtitle v-text="art.description"></v-list-item-subtitle>
                <v-list-item-subtitle class="text--secondary text-end">
                  Последнее изменение: {{ art.date }}
                </v-list-item-subtitle>
                <v-divider></v-divider>
              </v-list-item-content>
                <v-list-item-action>
                <v-btn icon @click="deleteDiaArt(art.id, art.title)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>
          </v-list>
        </v-tab-item>
        <v-tab-item>
          <v-row justify="center">
            <p class="text-center ma-8">Работа с закладками пока что не организованна</p>
          </v-row>
        </v-tab-item>
      </v-tabs>
    <v-dialog v-model=dialogConfirmDelete max-width="460px">
      <v-card>
        <v-card-text class="pa-4">
          <h3>Подтвердите удаление статьи: {{ titleArt}}</h3>
        </v-card-text>
        <v-card-actions>
          <v-btn @click="dialogConfirmDelete = false" color="blue darken-1" text>
            Отменить
          </v-btn>
          <v-spacer/>
          <v-btn @click="deleteArt(idArt)" color="pink darken-1" text>
            Удалить
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'ListArticles',

  data: () => ({
    idArt: 0,
    titleArt:'',
    dialogConfirmDelete: false,
  }),

  computed:{
    ...mapGetters(['ArticleHome', 'DraftHome']),
  },
  methods: {
    click() {},
    selectArt(id) {
      this.$router.push({ name:'Editor', params: {id:id}})
    },
    deleteArt(id) {
      this.$store.dispatch('deleteArticle', id)
      this.dialogConfirmDelete = false
    },
    deleteDiaArt(id, title) {
      this.idArt = id
      this.titleArt = title
      this.dialogConfirmDelete = true
    }
  },
}
</script>