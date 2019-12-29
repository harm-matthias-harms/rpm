<template>
  <!-- TODO add files to submit-->
  <v-form @submit.prevent="atSubmit({medicalCase: medicalCase, files: files})">
    <v-text-field
      v-model="medicalCase.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <v-expansion-panels v-model="panel" multiple class="mb-4">
      <v-expansion-panel>
        <v-expansion-panel-header>General Information</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-row>
            <v-col>
              <v-checkbox v-model="medicalCase.generalInformation.surgical" label="Surgical"></v-checkbox>
            </v-col>
            <v-col>
              <v-checkbox
                v-model="medicalCase.generalInformation.hospilisation"
                label="Need for Hospilisation"
              ></v-checkbox>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-checkbox v-model="medicalCase.generalInformation.usar" label="USAR"></v-checkbox>
            </v-col>
            <v-col>
              <v-checkbox v-model="medicalCase.generalInformation.medivac" label="MEDIVAC"></v-checkbox>
            </v-col>
          </v-row>
          <v-select
            v-model="medicalCase.generalInformation.triage"
            :items="['Less Urgent', 'Urgent', 'Emergent']"
            label="Triage"
          />
          <v-textarea v-model="medicalCase.generalInformation.shortSummary" label="Short Summary"></v-textarea>
          <v-select
            v-model="medicalCase.generalInformation.age"
            :items="['0-10', '11-17', '18-30', '31-60', '60+']"
            label="Age"
          />
          <v-select
            v-model="medicalCase.generalInformation.gender"
            :items="['Male', 'Female']"
            label="Gender"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Medical History</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-text-field v-model="medicalCase.medicalHistory.problems" label="Problems/Conditions"></v-text-field>
          <v-text-field v-model="medicalCase.medicalHistory.vaccinations" label="Vaccinations"></v-text-field>
          <v-text-field v-model="medicalCase.medicalHistory.allergies" label="Allergies"></v-text-field>
          <v-text-field v-model="medicalCase.medicalHistory.medication" label="Medication"></v-text-field>
          <v-text-field
            v-model="medicalCase.medicalHistory.implantedDevices"
            label="Implantable Devices"
          ></v-text-field>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          Vital Signs
          <template v-slot:actions v-if="medicalCase.vitalSigns.length == 0">
            <v-icon @click="addVitalSign">fas fa-plus</v-icon>
          </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-expansion-panels
            :value="[...Array(this.medicalCase.vitalSigns).keys()].map((k, i) => i)"
            multiple
            class="mb-4"
          >
            <VitalSign v-for="(vs, i) in medicalCase.vitalSigns" :key="i" :vitalSign.sync="vs" />
          </v-expansion-panels>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Expectations</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-textarea v-model="medicalCase.expectations.generalStatus" label="General Status"></v-textarea>
          <v-textarea v-model="medicalCase.expectations.onExamination" label="On Examination"></v-textarea>
          <v-textarea v-model="medicalCase.expectations.expectations" label="Expectations"></v-textarea>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-textarea v-model="medicalCase.otherInformation" label="Other Information"></v-textarea>
    <v-textarea v-model="medicalCase.makeup" label="Needed Make-Up and Attributes"></v-textarea>
    <v-file-input v-model="files" chips multiple show-size label="Files"></v-file-input>
    <v-btn
      :disabled="medicalCase.title && errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >create</v-btn>
    <v-btn @click="$router.back()">cancel</v-btn>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import VitalSign from './vitalSign.vue'
@Component({
  components: {
    VitalSign
  }
})
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly medicalCase!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: void

  panel: Array<Number> = [0, 3]
  files: Array<any> = []
  basisVitalSignOpen: boolean = false
  emptyVitalSign: object = {
    title: undefined,
    reason: undefined,
    data: {
      oos: undefined,
      avpu: undefined,
      mobility: undefined,
      respiratoryRate: undefined,
      pulse: undefined,
      temperature: undefined,
      capillaryRefill: undefined,
      bloodPressureSystolic: undefined,
      bloodPressureDiastolic: undefined,
      oxygenSaturation: undefined,
      weight: undefined,
      height: undefined
    },
    childs: []
  }

  addVitalSign() {
    if (!this.basisVitalSignOpen && this.medicalCase.vitalSigns.length == 0) {
      this.medicalCase.vitalSigns.push(this.emptyVitalSign)
    }
    this.basisVitalSignOpen = !this.basisVitalSignOpen
  }
}
</script>
