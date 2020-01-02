<template>
  <v-expansion-panels
    multiple
    :value="[...Array(vitalSigns).keys()].map((k, i) => i)"
  >
    <v-expansion-panel
      v-for="(vitalSign, i) in vitalSigns"
      :key="i"
    >
      <v-expansion-panel-header>{{ vitalSign.title }}</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-list dense>
          <v-list-item v-if="vitalSign.reason">
            <v-list-item-content>Reason:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ vitalSign.reason }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.oos">
            <v-list-item-content>Onset of Symptoms:</v-list-item-content>
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
          <v-list-item v-if="vitalSign.data.pulse">
            <v-list-item-content>Pulse:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ valueToString(vitalSign.data.pulse, 'bpm') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            v-if="vitalSign.data.bloodPressureSystolic || vitalSign.data.bloodPressureSystolic "
          >
            <v-list-item-content>Blood Pressure:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.bloodPressureSystolic, '') }}{{ vitalSign.data.bloodPressureSystolic ? "/" : "" }}{{ valueToString(vitalSign.data.bloodPressureDiastolic, '') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.respiratoryRate">
            <v-list-item-content>Respiratory Rate:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.respiratoryRate, 'bpm') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.oxygenSaturation">
            <v-list-item-content>Oxygen Saturation:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.oxygenSaturation, '%') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.capillaryRefill">
            <v-list-item-content>Capillary Refill:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.capillaryRefill, 's') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.temperature">
            <v-list-item-content>Temperatur:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.temperature, '°C') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.weight">
            <v-list-item-content>Weight:</v-list-item-content>
            <v-list-item-content class="align-end">
              {{ valueToString(vitalSign.data.weight, 'kg') }}
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-if="vitalSign.data.height">
            <v-list-item-content>Height:</v-list-item-content>
            <v-list-item-content
              class="align-end"
            >
              {{ valueToString(vitalSign.data.height / 100, 'm') }}
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

@Component
  export default class VitalSigns extends Vue {
  @Prop({ type: Array, required: true }) readonly vitalSigns!: any

  valueToString (value: number, unit: string) {
    if (value !== 0) {
      return value + (unit ? ' ' + unit : '')
    }
  }
  }
</script>
