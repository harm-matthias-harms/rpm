<template>
  <v-form @submit.prevent="submit()">
    <v-text-field
      v-model="medicalCase.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <v-expansion-panels
      v-model="expansionPanel"
      multiple
      class="mb-4"
    >
      <v-expansion-panel>
        <v-expansion-panel-header>General Information</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-row>
            <v-col>
              <v-checkbox
                v-model="medicalCase.generalInformation.surgical"
                label="Surgical"
              />
            </v-col>
            <v-col>
              <v-checkbox
                v-model="medicalCase.generalInformation.hospilisation"
                label="Need for hospilisation"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-checkbox
                v-model="medicalCase.generalInformation.usar"
                label="USAR"
              />
            </v-col>
            <v-col>
              <v-checkbox
                v-model="medicalCase.generalInformation.medivac"
                label="MEDIVAC"
              />
            </v-col>
          </v-row>
          <v-select
            v-model="medicalCase.generalInformation.triage"
            :items="['Less Urgent', 'Urgent', 'Emergent']"
            label="Triage"
          />
          <v-textarea
            v-model="medicalCase.generalInformation.shortSummary"
            label="Short summary"
          />
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
        <v-expansion-panel-header>Medical history</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-text-field
            v-model="medicalCase.medicalHistory.problems"
            label="Problems/conditions"
          />
          <v-text-field
            v-model="medicalCase.medicalHistory.vaccinations"
            label="Vaccinations"
          />
          <v-text-field
            v-model="medicalCase.medicalHistory.allergies"
            label="Allergies"
          />
          <v-text-field
            v-model="medicalCase.medicalHistory.medication"
            label="Medication"
          />
          <v-text-field
            v-model="medicalCase.medicalHistory.implantedDevices"
            label="Implantable devices"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          Vital Signs
          <template
            v-if="!medicalCase.vitalSigns || medicalCase.vitalSigns.length === 0"
            v-slot:actions
          >
            <v-icon @click="addVitalSign">
              fas fa-plus
            </v-icon>
          </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-expansion-panels
            :value="medicalCase.vitalSigns.map((k, i) => i)"
            multiple
            class="mb-4"
          >
            <VitalSign
              v-for="(vs, i) in medicalCase.vitalSigns"
              :key="i"
              :vital-sign.sync="medicalCase.vitalSigns[i]"
            />
          </v-expansion-panels>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Expectations</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-textarea
            v-model="medicalCase.expectations.generalStatus"
            label="General status"
          />
          <v-textarea
            v-model="medicalCase.expectations.onExamination"
            label="On examination"
          />
          <v-textarea
            v-model="medicalCase.expectations.expectations"
            label="Expectations"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-textarea
      v-model="medicalCase.otherInformation"
      label="Other information"
    />
    <v-textarea
      v-model="medicalCase.makeup"
      label="Needed make-up and attributes"
    />
    <v-file-input
      v-model="files"
      chips
      multiple
      show-size
      label="Files"
    />
    <v-btn
      :disabled="!medicalCase.title || errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >
      {{ isNew ? "create" : "edit" }}
    </v-btn>
    <v-btn @click="$router.back()">
      cancel
    </v-btn>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import VitalSign from './vital_signs/form.vue'
@Component({
  components: {
    VitalSign
  }
})
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly medicalCase!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: (payload) => void
  @Prop({ type: Boolean, required: true }) readonly isNew!: boolean

  files: Array<any> = []
  expansionPanel: Array<number> = [0, 3]
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

  mounted () {
    this.setExpansionPanel()
  }

  setExpansionPanel () {
    if (this.medicalCase) {
      if (this.anyFieldPresent(this.medicalCase.generalInformation)) {
        if (!this.expansionPanel.includes(0)) {
          this.expansionPanel.push(0)
        }
      }
      if (this.anyFieldPresent(this.medicalCase.medicalHistory)) {
        if (!this.expansionPanel.includes(1)) {
          this.expansionPanel.push(1)
        }
      }
      if (
        this.medicalCase.vitalSigns &&
        this.medicalCase.vitalSigns.length > 0 &&
        this.anyFieldPresent(this.medicalCase.vitalSigns[0])
      ) {
        if (!this.expansionPanel.includes(2)) {
          this.expansionPanel.push(2)
        }
      }
      if (this.anyFieldPresent(this.medicalCase.expectations)) {
        if (!this.expansionPanel.includes(3)) {
          this.expansionPanel.push(3)
        }
      }
    }
  }

  anyFieldPresent (obj: Object) {
    const exist = (elem) => {
      if (elem) {
        return true
      }
    }
    return Object.values(obj).some(exist)
  }

  addVitalSign () {
    if (
      !this.medicalCase.vitalSigns ||
      this.medicalCase.vitalSigns.length === 0
    ) {
      this.medicalCase.vitalSigns = [Object.assign({}, this.emptyVitalSign)]
    }
  }

  vitalSignsEmptyStringFix (vs) {
    const keys = Object.keys(vs.data)
    keys.forEach((k) => {
      if (vs.data[k] === '') {
        delete vs.data[k]
      }
    })
    vs.childs.forEach((c) => {
      this.vitalSignsEmptyStringFix(c)
    })
  }

  submit () {
    this.medicalCase.vitalSigns.forEach((vs) => {
      this.vitalSignsEmptyStringFix(vs)
    })
    this.atSubmit({ medicalCase: this.medicalCase, files: this.files })
  }
}
</script>
