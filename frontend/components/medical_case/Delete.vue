<template>
  <div style="display: inline;">
    <v-icon
      color="red"
      @click="dialog = true"
    >
      delete
    </v-icon>
    <Confirm
      text="Are you sure you want to delete this medical case?"
      :dialog.sync="dialog"
      :item="item"
      :at-submit="deleteMedicalCase"
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
    ...mapActions('medicalCase', {
      delete: 'delete'
    })
  }
})
export default class DeleteButton extends Vue {
  @Prop({ type: Object, required: true }) readonly item!: object
  @Prop({ type: Boolean, required: true }) readonly goBack!: boolean

  dialog: boolean = false
  delete!: ({ id: string, goBack: boolean }) => void

  deleteMedicalCase (mc) {
    this.delete({ id: mc.id, goBack: this.goBack })
    this.dialog = false
  }

  onCancel () {
    this.dialog = false
  }
}
</script>
