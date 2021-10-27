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
    <v-expansion-panels v-model="expansionPanel" multiple class="mb-4">
      <v-expansion-panel>
        <v-expansion-panel-header>General</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-select
            v-model="medicalCase.general.discipline"
            :items="[
              'Internal med',
              'Surgery',
              'Gyn/Obs',
              'Infectious diseases',
              'Trauma',
              'Public health',
            ]"
            label="Area/Discipline"
            multiple
          />
          <v-select
            v-model="medicalCase.general.context"
            :items="['Europe', 'LAMIC']"
            :multiple="true"
            label="Context"
          />
          <v-select
            v-model="medicalCase.general.scenario"
            :multiple="true"
            :items="[
              'Conflict',
              'Natural disaster',
              'Man-made disaster',
              'Mass-Casualty-Incident',
              'Outbreak',
              'People displacement/Refugees',
            ]"
            label="Scenario"
          />
          <v-checkbox
            v-model="medicalCase.general.preHospital"
            label="Prehospital playable"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header
          >Patient's characteristics</v-expansion-panel-header
        >
        <v-expansion-panel-content>
          <v-select
            v-model="medicalCase.patient.type"
            :items="['Chronic', 'Acute']"
            label="Patient type"
          />
          <v-select
            v-model="medicalCase.patient.triage"
            :items="['Deceased/Unsalvageable', 'Red', 'Yellow', 'Green']"
            label="Triage"
          />
          <v-select
            v-model="medicalCase.patient.gender"
            :multiple="true"
            :items="['Male', 'Female', 'Undefined']"
            label="Gender"
          />
          <v-select
            v-model="medicalCase.patient.age"
            :items="['Neonate', 'Infant', 'Adolescent', 'Adult', 'Elderly']"
            label="Age"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Medical assessment</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-text-field
            v-model="medicalCase.medical.signs"
            label="S – Signs/Symptoms"
          />
          <v-text-field
            v-model="medicalCase.medical.allergies"
            label="A – Allergies"
          />
          <v-text-field
            v-model="medicalCase.medical.medication"
            label="M – Medications"
          />
          <v-text-field
            v-model="medicalCase.medical.past"
            label="P – Past pertinent medical history"
          />
          <v-text-field
            v-model="medicalCase.medical.loi"
            label="L – Last oral intake"
          />
          <v-text-field
            v-model="medicalCase.medical.events"
            label="E – Events leading up to present illness/injury"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          Vital signs
          <template
            v-if="
              !medicalCase.vitalSigns || medicalCase.vitalSigns.length === 0
            "
            v-slot:actions
          >
            <v-icon @click="addVitalSign"> fas fa-plus </v-icon>
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
              :level="0"
            />
          </v-expansion-panels>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header
          >Make-up and attributes</v-expansion-panel-header
        >
        <v-expansion-panel-content>
          <v-textarea
            v-model="medicalCase.makeup.makeup"
            label="Make-up instructions"
          />
          <v-textarea
            v-model="medicalCase.makeup.acting"
            label="Acting instructions"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-file-input v-model="files" chips multiple show-size label="Files" />
    <v-btn
      :disabled="!medicalCase.title || errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >
      {{ isNew ? 'create' : 'edit' }}
    </v-btn>
    <v-btn @click="$router.back()"> cancel </v-btn>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import VitalSign from './vital_signs/form.vue'
@Component({
  components: {
    VitalSign,
  },
})
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly medicalCase!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: (
    payload
  ) => void
  @Prop({ type: Boolean, required: true }) readonly isNew!: boolean

  files: Array<any> = []
  expansionPanel: Array<number> = [0, 1, 2, 4]
  emptyVitalSign: object = {
    title: undefined,
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
      expectations: {
        foe: undefined,
        treatmentExpected: undefined,
      },
    },
    childs: [],
  }

  mounted() {
    this.setExpansionPanel()
  }

  setExpansionPanel() {
    if (this.medicalCase) {
      if (this.anyFieldPresent(this.medicalCase.general)) {
        if (!this.expansionPanel.includes(0)) {
          this.expansionPanel.push(0)
        }
      }
      if (this.anyFieldPresent(this.medicalCase.patient)) {
        if (!this.expansionPanel.includes(1)) {
          this.expansionPanel.push(1)
        }
      }
      if (this.anyFieldPresent(this.medicalCase.medical)) {
        if (!this.expansionPanel.includes(2)) {
          this.expansionPanel.push(2)
        }
      }
      if (
        this.medicalCase.vitalSigns &&
        this.medicalCase.vitalSigns.length > 0 &&
        this.anyFieldPresent(this.medicalCase.vitalSigns[0])
      ) {
        if (!this.expansionPanel.includes(3)) {
          this.expansionPanel.push(3)
        }
      }
      if (this.anyFieldPresent(this.medicalCase.makeup)) {
        if (!this.expansionPanel.includes(4)) {
          this.expansionPanel.push(4)
        }
      }
    }
  }

  anyFieldPresent(obj: Object) {
    const exist = (elem) => {
      if (elem) {
        return true
      }
    }
    return Object.values(obj).some(exist)
  }

  addVitalSign() {
    if (
      !this.medicalCase.vitalSigns ||
      this.medicalCase.vitalSigns.length === 0
    ) {
      this.medicalCase.vitalSigns = [Object.assign({}, this.emptyVitalSign)]
    }
  }

  vitalSignsEmptyStringFix(vs) {
    const keys = Object.keys(vs.data)
    keys.forEach((k) => {
      if (vs.data[k] === '') {
        delete vs.data[k]
      }
    })
    Object.keys(vs.data.expectations).forEach((k) => {
      if (vs.data.expectations[k] === '') {
        delete vs.data.expectations[k]
      }
    })
    vs.childs.forEach((c) => {
      this.vitalSignsEmptyStringFix(c)
    })
  }

  submit() {
    this.medicalCase.vitalSigns.forEach((vs) => {
      this.vitalSignsEmptyStringFix(vs)
    })
    this.atSubmit({ medicalCase: this.medicalCase, files: this.files })
  }
}
</script>
