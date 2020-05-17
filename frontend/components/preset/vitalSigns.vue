<template>
  <v-row justify="center">
    <v-col>
      <p class="headline mb-2 pl-4">
        Vital signs
      </p>
      <v-list dense class="body-1">
        <v-list-item v-if="vitalSigns.oos">
          <v-list-item-content>Onset of symptoms:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ vitalSigns.oos }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.avpu">
          <v-list-item-content>AVPU:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ vitalSigns.avpu }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.mobility">
          <v-list-item-content>Mobility:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ vitalSigns.mobility }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.pulse != null">
          <v-list-item-content>Pulse:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.pulse, 'bpm') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item
          v-if="vitalSigns.bloodPressureSystolic != null || vitalSigns.bloodPressureDiastolic != null"
        >
          <v-list-item-content>Blood pressure:</v-list-item-content>
          <v-list-item-content
            class="align-end"
          >
            {{ valueToString(vitalSigns.bloodPressureSystolic, '') }}{{ vitalSigns.bloodPressureSystolic ? "/" : "" }}{{ valueToString(vitalSigns.bloodPressureDiastolic, '') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.respiratoryRate != null">
          <v-list-item-content>Respiratory rate:</v-list-item-content>
          <v-list-item-content
            class="align-end"
          >
            {{ valueToString(vitalSigns.respiratoryRate, 'bpm') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.oxygenSaturation != null">
          <v-list-item-content>Oxygen saturation:</v-list-item-content>
          <v-list-item-content
            class="align-end"
          >
            {{ valueToString(vitalSigns.oxygenSaturation, '%') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.capillaryRefill != null">
          <v-list-item-content>Capillary refill:</v-list-item-content>
          <v-list-item-content
            class="align-end"
          >
            {{ valueToString(vitalSigns.capillaryRefill, 's') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.temperature != null">
          <v-list-item-content>Body temperatur:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.temperature, '°C') }}
          </v-list-item-content>
        </v-list-item>
        <p class="headline mb-2 pl-4">
          Expectations
        </p>
        <v-list-item v-if="vitalSigns.expectations.foe">
          <v-list-item-content>Findings on examination:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ vitalSigns.expectations.foe }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.expectations.treatmentExpected">
          <v-list-item-content>Expected treatment:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ vitalSigns.expectations.treatmentExpected }}
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class VitalSigns extends Vue {
  @Prop({ type: Object }) readonly vitalSigns?: object

  valueToString (value: number, unit: string) {
    return value + (unit ? ' ' + unit : '')
  }
}
</script>

<style scoped>
p,
.v-list-item__content {
  white-space: pre-line;
}
</style>
