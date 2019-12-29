<template>
  <div>
    <v-text-field v-model="vitalSigns.oos" label="Onset of Symptoms" name="oos" />
    <v-select
      v-model="vitalSigns.avpu"
      :items="['Altert', 'Voice', 'Pain', 'Unresponsive']"
      label="AVPU"
    />
    <v-text-field v-model="vitalSigns.mobility" label="Mobility" />
    <v-text-field
      v-model.number="vitalSigns.pulse"
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
          v-model.number="vitalSigns.bloodPressureSystolic"
          v-validate="'integer'"
          type="number"
          label="Blood Pressure Systolic"
          data-vv-name="systolic"
          :error-messages="errors.collect('systolic')"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model.number="vitalSigns.bloodPressureDiastolic"
          v-validate="'integer'"
          type="number"
          label="Blood Pressure Diastolic"
          data-vv-name="diastolic"
          :error-messages="errors.collect('diastolic')"
        />
      </v-col>
    </v-row>
    <v-text-field
      v-model.number="vitalSigns.respiratoryRate"
      v-validate="'integer'"
      type="number"
      suffix="bpm"
      label="Respiratory Rate"
      data-vv-name="respiratory rate"
      :error-messages="errors.collect('respiratory rate')"
    />
    <v-text-field
      v-model.number="vitalSigns.oxygenSaturation"
      v-validate="'integer'"
      type="number"
      suffix="%"
      label="Oxygen Saturation"
      data-vv-name="oxygen saturation"
      :error-messages="errors.collect('oxygen saturation')"
    />
    <v-text-field
      v-model.number="vitalSigns.capillaryRefill"
      type="number"
      suffix="s"
      label="Capillary Refill"
    />
    <v-text-field
      v-model.number="vitalSigns.temperature"
      type="number"
      suffix="°C"
      label="Temperature"
    />
    <v-text-field v-model.number="vitalSigns.weight" type="number" suffix="kg" label="Weight" />
    <v-text-field
      v-model.number="vitalSigns.height"
      v-validate="'integer'"
      type="number"
      suffix="cm"
      label="Height"
      data-vv-name="height"
      :error-messages="errors.collect('height')"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
@Component
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSigns!: object
  @Watch('vitalSigns', { immediate: true, deep: true })
  updateVitalSignsChanged(val: any, oldVal: any) {
    this.$emit('update:vitalSigns', val)
  }
}
</script>
