<template>
  <div>
    <p class="title text-center">Dynamic Casualties Card</p>
    <v-row>
      <v-col class="text-center">
        <p>
          Casualty
          <br />
          <span class="text-h5 grey lighten-2 px-6 py-1">
            {{ vitalSign.fullname }}
          </span>
        </p>
      </v-col>
      <v-col class="text-center">
        <p>
          Status
          <br />
          <span class="text-h5 grey lighten-2 px-6 py-1">
            {{ vitalSign.data.title }}
          </span>
        </p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-icon size="64">fas fa-lungs</v-icon>
          </v-col>
          <v-col cols="auto" class="text-center">
            Respiratory Rate
            <br />
            <span class="text-h3 has-border-grey px-4">
              {{ vitalSign.data.respiratoryRate }}
            </span>
          </v-col>
        </v-row>
      </v-col>
      <v-col>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-icon size="64">fas fa-heartbeat</v-icon>
          </v-col>
          <v-col cols="auto" class="text-center">
            Heart Rate
            <br />
            <span class="text-h3 has-border-grey px-4">
              {{ vitalSign.data.pulse }}
            </span>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-icon size="64">fas fa-wind</v-icon>
          </v-col>
          <v-col cols="auto" class="text-center">
            SpO2
            <br />
            <span class="text-h3 has-border-grey px-4">
              {{ vitalSign.data.oxygenSaturation }}
            </span>
          </v-col>
        </v-row>
      </v-col>
      <v-col>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-icon size="64">fas fa-stethoscope</v-icon>
          </v-col>
          <v-col cols="auto" class="text-center">
            Blood Pressure
            <br />
            <span class="text-h3 has-border-grey px-4">
              {{ vitalSign.data.bloodPressureSystolic }}/{{
                vitalSign.data.bloodPressureDiastolic
              }}
            </span>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-icon size="64">fas fa-thermometer-three-quarters</v-icon>
          </v-col>
          <v-col cols="auto" class="text-center">
            <span class="mb-1"> Temperature </span>
            <br />
            <span class="text-h3 has-border-grey px-4">
              {{ vitalSign.data.temperature }}
            </span>
          </v-col>
        </v-row>
      </v-col>
      <v-col class="text-center"> </v-col>
    </v-row>
    <v-row justify="center" class="mt-4">
      <v-col cols="auto">
        <p class="caption">
          Prior status: {{ vitalSign.data.priorStatus || '-' }}
        </p>
      </v-col>
      <v-col cols="auto">
        <p class="caption">
          MC: {{ vitalSign.mcId }} | Case:
          {{ vitalSign.id }}
        </p>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'

@Component
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSign!: Object
}
</script>

<style scoped>
.has-border-grey {
  border: 1px solid grey;
}

.text-h3 {
  font-size: 3rem !important;
  font-weight: 500;
}

.row + .row {
  margin-top: 0;
}
</style>