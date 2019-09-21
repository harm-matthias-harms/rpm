<template>
  <v-form @submit.prevent="atSubmit(preset)">
    <v-text-field
      v-model="preset.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <v-text-field v-model="preset.vitalSigns.oos" label="Onset of Symptoms" />
    <v-select
      v-model="preset.vitalSigns.avpu"
      :items="['Altert', 'Voice', 'Pain', 'Unresponsive'] "
      label="AVPU"
    />
    <v-text-field v-model="preset.vitalSigns.mobility" label="Mobility" />
    <v-text-field
      v-model="preset.vitalSigns.pulse"
      v-validate="'integer'"
      type="number"
      suffix="bpm"
      label="Pulse"
      data-vv-name="pulse"
      :error-messages="errors.collect('pulse')"
    />
    <v-row>
      <v-col>
        <v-text-field
          v-model="preset.vitalSigns.bloodPressureSystolic"
          v-validate="'integer'"
          type="number"
          label="Blood pressure systolic"
          data-vv-name="systolic"
          :error-messages="errors.collect('systolic')"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model="preset.vitalSigns.bloodPressureDiastolic"
          v-validate="'integer'"
          type="number"
          label="Blood pressure diastolic"
          data-vv-name="diastolic"
          :error-messages="errors.collect('diastolic')"
        />
      </v-col>
    </v-row>
    <v-text-field
      v-model="preset.vitalSigns.respiratoryRate"
      v-validate="'integer'"
      type="number"
      suffix="bpm"
      label="Respiratory Rate"
      data-vv-name="respiratory rate"
      :error-messages="errors.collect('respiratory rate')"
    />
    <v-text-field
      v-model="preset.vitalSigns.oxygenSaturation"
      v-validate="'integer'"
      type="number"
      suffix="%"
      label="Oxygen Saturation"
      data-vv-name="oxygen saturation"
      :error-messages="errors.collect('oxygen saturation')"
    />
    <v-text-field
      v-model="preset.vitalSigns.capillaryRefill"
      type="number"
      suffix="s"
      label="Capillary Refill"
    />
    <v-text-field
      v-model="preset.vitalSigns.temperature"
      type="number"
      suffix="°C"
      label="Temperatur"
    />
    <v-text-field v-model="preset.vitalSigns.weight" type="number" suffix="kg" label="Weight" />
    <v-text-field
      v-model="preset.vitalSigns.height"
      v-validate="'integer'"
      type="number"
      suffix="cm"
      label="Height"
      data-vv-name="height"
      :error-messages="errors.collect('height')"
    />

    <v-btn
      :disabled="preset.title && errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >
      create
    </v-btn>
    <v-btn @click="$router.back()">
      cancel
    </v-btn>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class PresetForm extends Vue {
  @Prop({ type: Object, required: true }) readonly preset!: object
  @Prop({ type: Function, required: true }) readonly atSubmit!: void
}
</script>
