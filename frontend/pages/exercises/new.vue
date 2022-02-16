<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
      <v-card>
        <v-card-title primary-title> New Exercise </v-card-title>
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
import { cloneDeep } from 'lodash'

import Form from '@/components/exercise/Form.vue'
import { defaultExercise } from '~/store/exercise/state'
import { Exercise } from '~/store/exercise/type'

@Component({
  components: {
    Form,
  },
  methods: {
    ...mapActions('exercise', {
      create: 'create',
    }),
  },
})
export default class NewExercise extends Vue {
  exercise: Exercise = cloneDeep(defaultExercise)
  create!: (exercise: Exercise) => void
}
</script>
