<template>
  <v-form @submit.prevent="submit()">
    <v-text-field
      v-model="team.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <v-select
      :items="['AMP', 'AMPS', 'EMT 1', 'EMT 1 mobile', 'EMT 2', 'EMT 3', 'other']"
      v-model="team.type"
      label="Type"
      v-validate="'required'"
    />
    <v-checkbox v-model="team.medivac" label="Medivac" />
    <v-btn
      :disabled="!team.title || !team.type || errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >{{ isNew ? "create" : "edit" }}</v-btn>
    <v-btn @click="$router.back()">cancel</v-btn>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class TeamForm extends Vue {
  @Prop({ type: Object, required: true }) readonly team!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: (
    payload
  ) => void
  @Prop({ type: Boolean, required: true }) readonly isNew!: boolean

  submit() {
    this.atSubmit(this.team)
  }
}
</script>
