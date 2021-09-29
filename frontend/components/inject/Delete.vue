<template>
  <div style="display: inline">
    <v-icon color="red" @click="dialog = true">
      delete
    </v-icon>
    <Confirm
      text="Are you sure you want to delete this case?"
      :dialog.sync="dialog"
      :item="item"
      :at-submit="deleteInject"
      :at-cancel="onCancel"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import Confirm from '@/components/utils/Confirm.vue'
import { Inject } from '~/store/inject/type'
@Component({
  components: { Confirm },
  methods: {
    ...mapActions('inject', {
      delete: 'delete'
    })
  }
})
export default class DeleteButton extends Vue {
  @Prop({ type: Object, required: true }) readonly item!: object
  @Prop({ type: String, required: true }) readonly exerciseId!: string
  @Prop({ type: Boolean, required: true }) readonly goBack!: boolean

  dialog: boolean = false
  delete!: ({ id, exerciseID, goBack }) => void

  deleteInject (inject: Inject) {
    this.delete({
      id: inject.id,
      exerciseID: this.exerciseId,
      goBack: this.goBack
    })
    this.dialog = false
  }

  onCancel () {
    this.dialog = false
  }
}
</script>
