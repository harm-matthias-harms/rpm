<template>
  <div>
    <div
      v-for="(chunk, idx) in _.chunk(vitalSigns, 2)"
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
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { mapActions } from 'vuex'

import VitalSign from '@/components/inject/print/VitalSign.vue'
import { Inject } from '~/store/inject/type'
import { MedicalCase, nestedVitalSign } from '~/store/medicalCase/type'
import _ from 'lodash'

interface vitalSign {
  id: string | undefined
  mcId: string
  fullname?: string
  data: vitalSignData
}

interface vitalSignData {
  title?: string
  priorStatus: string | null
  respiratoryRate?: number
  pulse?: number
  temperature?: number
  capillaryRefill?: number
  bloodPressureSystolic?: number
  bloodPressureDiastolic?: number
  oxygenSaturation?: number
}

@Component({
  layout: 'print',
  components: {
    VitalSign,
  },
  methods: {
    ...mapActions('inject', {
      findInject: 'find',
    }),
  },
})
export default class PrintVitalSigns extends Vue {
  findInject!: ({ id, exerciseID, disableLoader }) => Promise<Inject>

  injects: Inject[] = []

  async created() {
    const ids: string | (string | null)[] =
      typeof this.$route.query.injects === 'string'
        ? [this.$route.query.injects]
        : this.$route.query.injects

    for (const id of ids) {
      const inject: Inject = await this.$axios.$get(
        `/api/exercises/${this.$route.params.id}/injects/${id}`
      )
      this.injects.push(inject)
    }
    setTimeout(window.print, 500)
  }

  get vitalSigns(): vitalSign[] {
    const results: vitalSign[] = []
    this.injects.forEach((inject: Inject) => {
      const vt: vitalSign[] | null = this.extractVitalSigns(inject)
      if (vt) vt.forEach((v) => results.push(v))
    })
    return results
  }

  get _() {
    return _
  }

  getMedicalCaseId(medicalCase: MedicalCase) {
    const regex = /(P[0-9]+)/i
    const match = medicalCase.title.match(regex)?.[1]

    if (match) return match
    return 'ø'
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

  vitalSignData(
    vitalSign: nestedVitalSign,
    priorStatus: string | null = null
  ): vitalSignData[] {
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
        priorStatus: priorStatus,
        ...vitalSign.data,
      })
    }
    if (vitalSign.childs) {
      vitalSign.childs.forEach((c) => {
        const res = this.vitalSignData(c, vitalSign.title)
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
</style>