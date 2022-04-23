<template>
  <div>
    <div class="page portrait">
      <v-row>
        <v-col>
          <h3>Patient</h3>
          <v-list dense>
            <v-list-item>
              <v-list-item-content>Fullname:</v-list-item-content>
              <v-list-item-content>
                {{ inject.roleplayer.fullName }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Age:</v-list-item-content>
              <v-list-item-content>
                {{ inject.roleplayer.age }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Gender:</v-list-item-content>
              <v-list-item-content>
                {{ inject.roleplayer.gender }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <h3>Patient's characteristics</h3>
          <v-list dense>
            <v-list-item>
              <v-list-item-content>Patient type:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.patient.type"
              >
                {{ inject.medicalCase.patient.type }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Gender:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.patient.gender"
              >
                {{ inject.medicalCase.patient.gender.join(', ') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Age:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.patient.age"
              >
                {{ inject.medicalCase.patient.age }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <h3>Medical assessment</h3>
          <v-list dense>
            <v-list-item>
              <v-list-item-content> S - Signs/Symptoms: </v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.signs"
              >
                {{ inject.medicalCase.medical.signs }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>A - Allergies:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.allergies"
              >
                {{ inject.medicalCase.medical.allergies }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>M - Medications:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.medication"
              >
                {{ inject.medicalCase.medical.medication }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                P - Past pertinent medical history:
              </v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.past"
              >
                {{ inject.medicalCase.medical.past }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content> L - Last oral intake: </v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.loi"
              >
                {{ inject.medicalCase.medical.loi }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                E - Events leading up to present illness/injury:
              </v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="inject.medicalCase.medical.events"
              >
                {{ inject.medicalCase.medical.events }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
      <v-row justify="center" class="mt-auto">
        <v-col cols="auto">
          <p class="caption">
            MC: {{ getMedicalCaseId(inject.medicalCase) }} | Case:
            {{ inject.id }}
          </p>
        </v-col>
      </v-row>
    </div>
    <div class="page portrait" v-if="tZero">
      <v-row>
        <v-col>
          <h3>Vital signs</h3>
          <v-list dense>
            <v-list-item>
              <v-list-item-content>Onset of symptoms:</v-list-item-content>
              <v-list-item-content class="align-end" v-if="tZero.data.oos">
                {{ tZero.data.oos }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>AVPU:</v-list-item-content>
              <v-list-item-content class="align-end" v-if="tZero.data.avpu">
                {{ tZero.data.avpu }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Mobility:</v-list-item-content>
              <v-list-item-content class="align-end" v-if="tZero.data.mobility">
                {{ tZero.data.mobility }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Pulse:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.pulse != null"
              >
                {{ valueToString(tZero.data.pulse, 'bpm') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Blood pressure:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="
                  tZero.data.bloodPressureSystolic != null ||
                  tZero.data.bloodPressureSystolic != null
                "
              >
                {{ valueToString(tZero.data.bloodPressureSystolic, '')
                }}{{ tZero.data.bloodPressureSystolic ? '/' : ''
                }}{{ valueToString(tZero.data.bloodPressureDiastolic, '') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Respiratory rate:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.respiratoryRate != null"
              >
                {{ valueToString(tZero.data.respiratoryRate, 'bpm') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Oxygen saturation:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.oxygenSaturation != null"
              >
                {{ valueToString(tZero.data.oxygenSaturation, '%') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Capillary refill:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.capillaryRefill != null"
              >
                {{ valueToString(tZero.data.capillaryRefill, 's') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Body temperatur:</v-list-item-content>
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.temperature != null"
              >
                {{ valueToString(tZero.data.temperature, '°C') }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content
                >Findings on examination:</v-list-item-content
              >
              <v-list-item-content
                class="align-end"
                v-if="tZero.data.expectations.foe"
              >
                {{ tZero.data.expectations.foe }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
      <v-row justify="center" class="mt-auto">
        <v-col cols="auto">
          <p class="caption">
            MC: {{ getMedicalCaseId(inject.medicalCase) }} | Case:
            {{ inject.id }}
          </p>
        </v-col>
      </v-row>
    </div>
    <div class="page portrait">
      <v-row class="is-dotted-bottom">
        <v-col>
          <Makeup :inject="inject"></Makeup>
        </v-col>
      </v-row>
      <v-row>
        <v-col
          v-if="
            inject.medicalCase.vitalSigns.length &&
            inject.medicalCase.vitalSigns[0].childs.length
          "
        >
          <p class="title text-center">Clinical status transitions</p>
          <v-treeview
            open-all
            :items="[vitalSignFlow(inject.medicalCase.vitalSigns[0])]"
          ></v-treeview>
          <p class="caption text-center mt-5">
            MC: {{ getMedicalCaseId(inject.medicalCase) }} | Case:
            {{ inject.id }}
          </p>
        </v-col>
        <v-col v-else-if="extractVitalSigns(inject).length">
          <VitalSign :vitalSign="extractVitalSigns(inject)[0]"></VitalSign>
        </v-col>
      </v-row>
    </div>
    <div
      v-if="
        inject.medicalCase.vitalSigns.length &&
        inject.medicalCase.vitalSigns[0].childs.length
      "
    >
      <div
        v-for="(chunk, idx) in _.chunk(extractVitalSigns(inject), 2)"
        :key="idx"
        class="page portrait"
      >
        <VitalSign
          class="is-half-page"
          v-for="(vs, idx) in chunk"
          :key="idx"
          :vitalSign="vs"
        ></VitalSign>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'nuxt-property-decorator'

import { Inject } from '~/store/inject/type'
import Makeup from '@/components/inject/print/Makeup.vue'
import VitalSign from '@/components/inject/print/VitalSign.vue'

import _ from 'lodash'
import { MedicalCase, nestedVitalSign } from '~/store/medicalCase/type'

interface vitalSign {
  id: string | undefined
  mcId: string
  fullname?: string
  data: vitalSignData
}

interface vitalSignData {
  title?: string
  respiratoryRate?: number
  pulse?: number
  temperature?: number
  capillaryRefill?: number
  bloodPressureSystolic?: number
  bloodPressureDiastolic?: number
  oxygenSaturation?: number
}

interface vitalSignFlow {
  name: string | undefined
  children: vitalSignFlow[] | undefined
}

@Component({
  layout: 'print',
  components: {
    Makeup,
    VitalSign,
  },
})
export default class PrintRoleplayer extends Vue {
  @Prop({ type: Object, required: true }) readonly inject!: Inject
  tZero: nestedVitalSign | null = null

  get _() {
    return _
  }

  getMedicalCaseId(medicalCase: MedicalCase) {
    const regex = /(P[0-9]+)/i
    const match = medicalCase.title.match(regex)?.[1]

    if (match) return match
    return 'ø'
  }

  valueToString(value: number, unit: string) {
    return value + (unit ? ' ' + unit : '')
  }

  vitalSignFlow(vitalSign: nestedVitalSign): vitalSignFlow {
    return {
      name: vitalSign.title,
      children: vitalSign.childs?.map((child) => this.vitalSignFlow(child)),
    }
  }

  extractVitalSigns(inject: Inject): vitalSign[] | null {
    const result: vitalSign[] = []
    if (inject.medicalCase.vitalSigns) {
      inject.medicalCase.vitalSigns.forEach((vt) => {
        const res = this.vitalSignData(vt)
        if (res) {
          res.forEach((r) => {
            result.push({
              id: inject.id,
              mcId: this.getMedicalCaseId(inject.medicalCase),
              fullname: inject.roleplayer.fullName,
              data: r,
            })
          })
        }
      })
    }
    return result
  }

  vitalSignData(vitalSign: nestedVitalSign): vitalSignData[] {
    const results: vitalSignData[] = []
    if (
      vitalSign.data?.respiratoryRate ||
      vitalSign.data?.pulse ||
      vitalSign.data?.oxygenSaturation ||
      vitalSign.data?.bloodPressureSystolic ||
      vitalSign.data?.bloodPressureDiastolic ||
      vitalSign.data?.temperature
    ) {
      results.push({
        title: vitalSign.title,
        ...vitalSign.data,
      })
    }
    if (vitalSign.childs) {
      vitalSign.childs.forEach((c) => {
        const res = this.vitalSignData(c)
        res.forEach((r) => results.push(r))
      })
    }
    return results
  }
}
</script>

<style scoped>
html {
  -webkit-print-color-adjust: exact;
}

@page {
  size: 8.27in 11.69in;
  margin: 0;
}

html,
body {
  display: block;
  min-height: 0;
  background-color: #ffffff;
}

.page {
  position: relative;
  overflow: hidden;
  padding: 0.8in 0.8in;
  page-break-before: always;
}

.page.landscape {
  width: 11.69in;
  height: 8.27in;
  transform: translateY(11.69in) rotate(270deg);
  transform-origin: 0% 0%;
}

.page.portrait {
  width: 8.27in;
  height: 11.69in;
}

.is-third-page {
  height: 5.045in;
}

.is-third-page:not(:last-child) {
  border-bottom: 1px dotted black;
  padding-bottom: 0.5in;
}

.is-third-page:not(:first-child) {
  padding-top: 0.5in;
}

.is-half-page {
  height: 5.045in;
}

.is-half-page:first-child {
  border-bottom: 1px dotted black;
  padding-bottom: 0.5in;
}

.is-half-page:not(:first-child):last-child {
  padding-top: 0.5in;
}

.is-dotted-bottom {
  border-bottom: 1px dotted black;
}
</style>