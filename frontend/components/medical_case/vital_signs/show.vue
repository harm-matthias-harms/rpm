<template>
  <v-expansion-panels multiple :value="vitalSigns.map((k, i) => i)">
    <v-expansion-panel v-for="(vitalSign, i) in vitalSigns" :key="i">
      <v-expansion-panel-header>{{ vitalSign.title ? vitalSign.title : "No title" }}</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-list dense>
          <p class="headline mb-2 pl-4">
            Vital signs
          </p>
          <v-list-item v-if="vitalSign.data.oos">
            <v-list-item-content>Onset of symptoms:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.data.oos }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.avpu">
            <v-list-item-content>AVPU:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.data.avpu }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.mobility">
            <v-list-item-content>Mobility:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.data.mobility }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.pulse != null">
            <v-list-item-content>Pulse:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ valueToString(vitalSign.data.pulse, 'bpm') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            v-if="vitalSign.data.bloodPressureSystolic != null || vitalSign.data.bloodPressureSystolic != null"
          >
            <v-list-item-content>Blood pressure:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.bloodPressureSystolic, '') }}{{ vitalSign.data.bloodPressureSystolic ? "/" : "" }}{{ valueToString(vitalSign.data.bloodPressureDiastolic, '') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.respiratoryRate != null">
            <v-list-item-content>Respiratory rate:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.respiratoryRate, 'bpm') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.oxygenSaturation != null">
            <v-list-item-content>Oxygen saturation:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.oxygenSaturation, '%') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.capillaryRefill != null">
            <v-list-item-content>Capillary refill:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.capillaryRefill, 's') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.temperature != null">
            <v-list-item-content>Body temperatur:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.temperature, '°C') }}
            </v-list-item-content>
          </v-list-item>
          <p class="headline mb-2 pl-4">
            Expectations
          </p>
          <v-list-item v-if="vitalSign.data.expectations.foe">
            <v-list-item-content>Findings on examination:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.data.expectations.foe }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.expectations.treatmentExpected">
            <v-list-item-content>Expected treatment:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.data.expectations.treatmentExpected }}
            </v-list-item-content>
          </v-list-item>
        </v-list>
        <VitalSigns :vital-signs="vitalSign.childs" />
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'

@Component({
  name: 'VitalSigns'
})
export default class VitalSigns extends Vue {
  @Prop({ type: Array, required: true }) readonly vitalSigns!: any

  valueToString (value: number, unit: string) {
    return value + (unit ? ' ' + unit : '')
  }
}
</script>

<style scoped>
a {
  text-decoration: none;
}
p,
.v-list-item__content {
  white-space: pre-line;
}
</style>
