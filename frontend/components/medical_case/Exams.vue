<template>
  <v-container>
    <p class="title mx-4">
      Normal exams
    </p>
    <v-row class="ma-4">
      <v-col
        v-for="(file, i) in files"
        :key="i"
        cols="12"
        sm="6"
        md="4"
        lg="3"
      >
        <v-card>
          <v-icon>attach_file</v-icon>
          <v-card-title class="body-2">
            <a :href="`/exams/default/${file.path}`">
              {{ file.name }}
            </a>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
@Component
export default class Exams extends Vue {
  files = [
    {
      name: 'ABG',
      path: 'abg.pdf'
    },
    {
      name: 'Blood tests + Troponin',
      path: 'blood_tests+troponin.pdf'
    },
    {
      name: 'Chest X-Ray',
      path: 'chest_x-ray.pdf'
    },
    {
      name: 'Chest CT',
      path: 'ct_chest.pdf'
    },
    {
      name: 'ECG',
      path: 'ecg.pdf'
    },
    {
      name: 'Ultrasound fast',
      path: 'ultrasound_fast.pdf'
    },
    {
      name: 'Urinanalysis',
      path: 'urinalysis.pdf'
    }
  ]
}
</script>
