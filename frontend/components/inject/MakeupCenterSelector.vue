<template>
  <div>
    <v-autocomplete
      v-model="makeupCenter"
      :items="options"
      item-text="title"
      item-value="title"
    >
    </v-autocomplete>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import { MakeupCenter } from '~/store/exercise/type'
import { Team } from '~/store/team/type'

@Component({
  methods: {
    ...mapActions('exercise', {
      findExercise: 'find',
    }),
  },
})
export default class MakeupCenterSelector extends Vue {
  @Prop({ type: String }) readonly value!: string | undefined

  findExercise!: ({ id }) => Promise<void>

  get makeupCenter(): string | undefined {
    return this.value
  }

  set makeupCenter(makeupCenter: string | undefined) {
    this.$emit('input', makeupCenter)
  }

  options: MakeupCenter[] = []

  mounted() {
    const id = this.$route.params.id
    if (this.$store.state.exercise.exercise.id == id) {
      this.options = this.$store.state.exercise.exercise.makeupCenter
    } else {
      this.findExercise({ id }).then(() => {
        this.options = this.$store.state.exercise.exercise.makeupCenter
      })
    }
  }
}
</script>