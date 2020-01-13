<template>
  <v-form @submit.prevent="submit()">
    <v-text-field
      v-model="preset.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <Form :vital-signs.sync="preset.vitalSigns" />

    <v-btn
      :disabled="preset.title && errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >
      {{ isNew ? "create" : "edit" }}
    </v-btn>
    <v-btn @click="$router.back()">
      cancel
    </v-btn>
  </v-form>
</template>

<script lang="ts">
  import { Prop, Component, Vue } from 'vue-property-decorator'
  import Form from '@/components/vital_signs/form.vue'
@Component({
  components: {
    Form,
  },
})
  export default class PresetForm extends Vue {
  @Prop({ type: Object, required: true }) readonly preset!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: (payload) => void
  @Prop({ type: Boolean, required: true }) readonly isNew!: boolean

  submit () {
    const keys = Object.keys(this.preset.vitalSigns)
    keys.forEach((k) => {
      if (this.preset.vitalSigns[k] === '') {
        delete this.preset.vitalSigns[k]
      }
    })
    this.atSubmit(this.preset)
  }
  }
</script>
