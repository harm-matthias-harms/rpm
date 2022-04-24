<template>
  <div>
    <p class="title text-center">Makeup Instruction Card</p>

    <v-row justify="center">
      <v-col cols="6">
        <b>Casualty</b>
        <br />
        {{ inject.roleplayer.fullName }}
      </v-col>
      <v-col cols="6">
        <p>
          Gender: {{ inject.roleplayer.gender }}
          <br />
          Age: {{ inject.roleplayer.age ? inject.roleplayer.age : null }}
        </p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p><b>Symptoms description</b></p>
        <p>{{ inject.medicalCase.medical.signs }}</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p><b>Makeup description</b></p>
        <p>{{ inject.medicalCase.makeup.makeup }}</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p><b>Acting description</b></p>
        <p>{{ inject.medicalCase.makeup.acting }}</p>
      </v-col>
    </v-row>
    <v-row justify="center" class="mt-4">
      <v-col cols="auto">
        <p class="caption">
          MC: {{ getMedicalCaseId(inject.medicalCase) }} | Case:
          {{ inject.id }}
        </p>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { MedicalCase } from '~/store/medicalCase/type'

@Component
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly inject!: Object

  getMedicalCaseId(medicalCase: MedicalCase) {
    const regex = /(PT[0-9]+)/i
    const match = medicalCase.title.match(regex)?.[1]

    if (match) return match
    return 'ø'
  }
}
</script>

<style scoped>
p {
  margin-bottom: 0px;
}
</style>