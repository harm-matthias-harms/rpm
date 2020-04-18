<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
      <v-card>
        <v-card-title primary-title>
          New Exercise
        </v-card-title>
        <v-card-text>
          <Form :exercise="exercise" :at-submit="create" :is-new="true" />
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import Form from '@/components/exercise/Form.vue'

@Component({
  components: {
    Form
  },
  methods: {
    ...mapActions('exercise', {
      create: 'create'
    })
  }
})
export default class NewExercise extends Vue {
  exercise = { teams: [], roleplayManager: [], makeupCenter: [] }
  create!: (exercise) => void
}
</script>
