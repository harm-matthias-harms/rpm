<template>
  <v-dialog
    v-model="dialog"
    width="500"
  >
    <v-card>
      <v-card-title>{{ text }}</v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn
          text
          @click="atCancel()"
        >
          cancel
        </v-btn>
        <v-btn
          color="primary"
          text
          @click="atSubmit(item)"
        >
          ok
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
  import { Watch, Prop, Component, Vue } from 'vue-property-decorator'
@Component
  export default class Confirm extends Vue {
  @Prop({ type: Boolean, required: true }) readonly dialog!: boolean
  @Prop({ type: String, required: true }) readonly text!: string
  @Prop({ type: Object, required: true }) readonly item!: object
  @Prop({ type: Function, required: true }) readonly atSubmit!: (
    payload
  ) => void

  @Prop({ type: Function, required: true }) readonly atCancel!: () => void

  @Watch('dialog', { immediate: true })
  dialogChanged (val: any) {
    this.$emit('update:dialog', val)
  }
  }
</script>
