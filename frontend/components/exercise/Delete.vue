<template>
  <div style="display: inline;">
    <v-icon
      color="red"
      @click="dialog = true"
    >
      delete
    </v-icon>
    <Confirm
      text="Are you sure you want to delete this exercise?"
      :dialog.sync="dialog"
      :item="item"
      :at-submit="deleteExercise"
      :at-cancel="onCancel"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import Confirm from '@/components/utils/Confirm.vue'
@Component({
  components: { Confirm },
  methods: {
    ...mapActions('exercise', {
      delete: 'delete'
    })
  }
})
export default class DeleteButton extends Vue {
  @Prop({ type: Object, required: true }) readonly item!: object

  dialog: boolean = false
  delete!: ({ id: string }) => void

  deleteExercise (exercise) {
    this.delete({ id: exercise.id })
    this.dialog = false
  }

  onCancel () {
    this.dialog = false
  }
}
</script>
