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
        <!-- <p class="caption">Case title: {{ inject.medicalCase.title }}</p> -->
        <p class="caption">Case id: {{ inject.id }}</p>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'

@Component
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly inject!: Object
}
</script>

<style scoped>
p {
  margin-bottom: 0px;
}
</style>