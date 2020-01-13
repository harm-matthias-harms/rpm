<template>
  <v-row justify="center">
    <v-col>
      <p class="headline mb-2 pl-4">
        Vital Signs
      </p>
      <v-list
        dense
        class="body-1"
      >
        <v-list-item v-if="vitalSigns.oos">
          <v-list-item-content>Onset of Symptoms:</v-list-item-content>
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
        <v-list-item v-if="vitalSigns.bloodPressureSystolic != null || vitalSigns.bloodPressureDiastolic != null">
          <v-list-item-content>Blood Pressure:</v-list-item-content>
          <v-list-item-content
            class="align-end"
          >
            {{ valueToString(vitalSigns.bloodPressureSystolic, '') }}{{ vitalSigns.bloodPressureSystolic ? "/" : "" }}{{ valueToString(vitalSigns.bloodPressureDiastolic, '') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.respiratoryRate != null">
          <v-list-item-content>Respiratory Rate:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.respiratoryRate, 'bpm') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.oxygenSaturation != null">
          <v-list-item-content>Oxygen Saturation:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.oxygenSaturation, '%') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.capillaryRefill != null">
          <v-list-item-content>Capillary Refill:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.capillaryRefill, 's') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.temperature != null">
          <v-list-item-content>Temperatur:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.temperature, '°C') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.weight != null">
          <v-list-item-content>Weight:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.weight, 'kg') }}
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-if="vitalSigns.height != null">
          <v-list-item-content>Height:</v-list-item-content>
          <v-list-item-content class="align-end">
            {{ valueToString(vitalSigns.height / 100, 'm') }}
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
