<template>
  <div>
    <v-text-field
      v-model="vitalSigns.oos"
      label="Onset of symptoms"
      name="oos"
    />
    <v-select
      v-model="vitalSigns.avpu"
      :items="['A - Awake', 'V - Verbal', 'P - Pain', 'U - Unresponsive']"
      label="AVPU"
    />
    <v-select
      v-model="vitalSigns.mobility"
      :items="['Walking', 'Sitting', 'Laying down']"
      label="Mobility"
    />
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
          label="Blood pressure systolic"
          data-vv-name="systolic"
          :error-messages="errors.collect('systolic')"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model.number="vitalSigns.bloodPressureDiastolic"
          v-validate="'integer'"
          type="number"
          label="Blood pressure diastolic"
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
      label="Respiratory rate"
      data-vv-name="respiratory rate"
      :error-messages="errors.collect('respiratory rate')"
    />
    <v-text-field
      v-model.number="vitalSigns.oxygenSaturation"
      v-validate="'integer'"
      type="number"
      suffix="%"
      label="Oxygen saturation"
      data-vv-name="oxygen saturation"
      :error-messages="errors.collect('oxygen saturation')"
    />
    <v-text-field
      v-model.number="vitalSigns.capillaryRefill"
      type="number"
      suffix="s"
      label="Capillary refill"
    />
    <v-text-field
      v-model.number="vitalSigns.temperature"
      type="number"
      suffix="°C"
      label="Body temperature"
    />
    <v-textarea
      v-model.number="vitalSigns.expectations.foe"
      type="text"
      label="Findings on examination"
    />
    <v-textarea
      v-model.number="vitalSigns.expectations.treatmentExpected"
      type="text"
      label="Expected treatment"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
@Component
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSigns!: object
  @Watch('vitalSigns', { immediate: true, deep: true })
  updateVitalSignsChanged (val: any) {
    this.$emit('update:vitalSigns', val)
  }
}
</script>
