<template>
  <v-row justify="center">
    <v-col
      lg="6"
      md="10"
      sm="12"
    >
      <v-card>
        <v-card-title primary-title>
          Edit Exercise
        </v-card-title>
        <v-card-text>
          <Form
            :exercise="exercise"
            :at-submit="update"
            :is-new="false"
          />
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
      find: 'find',
      update: 'update'
    })
  }
})
export default class Edit extends Vue {
  find!: (id) => Promise<any>
  update!: (exercise) => void
  exercise: any = JSON.parse(JSON.stringify(this.$store.state.exercise.exercise))

  mounted () {
    const id = this.$route.params.id
    if (this.exercise.id !== id) {
      this.find({ id }).then(() => {
        this.exercise = JSON.parse(JSON.stringify(this.$store.state.exercise.exercise))
      })
    }
  }
}
</script>
