<template>
    <v-col
        cols="12"
        md="4"
        sm="6"
        xs="12"
        >
            <v-card
                v-ripple
                min-height="160"
                max-height="160"
                @click="selectSection()"
            >
            <v-card-text class="pa-4">
                <h2>{{title}}</h2>
            </v-card-text>
            </v-card>
        </v-col>
</template>

<script>
export default {
    name: 'SectionBtn',

    props: {
      id: Number,
      uuid: String,
      title: String,
      description: String,
    },

    methods: {
        selectSection() {
            this.$emit('selectsec')
        }
    }
}
</script>