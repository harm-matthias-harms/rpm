<template>
  <v-container>
    <v-row justify="center">
      <v-col md="4" sm="6">
        <h3 class="headline">
          <v-chip small label color="primary">{{ medicalCasesList.count }}</v-chip>Medical Cases
        </h3>
      </v-col>
      <v-col md="4" sm="6" class="justify-end d-flex">
        <v-btn small color="primary" class="mr-2" to="/medical_case/new">
          <v-icon small>fas fa-plus</v-icon>
        </v-btn>
        <v-btn small color="primary" @click="getMedicalCases()">
          <v-icon small>fas fa-redo</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="8" sm="12">
        <MedicalCaseTable :items="medicalCasesList.medicalCases" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import MedicalCaseTable from '@/components/medical_case/table.vue'
@Component({
  components: {
    MedicalCaseTable
  },
  computed: {
    ...mapState('medicalCase', {
      medicalCasesList: 'medicalCasesList',
      medicalCasesLoaded: 'medicalCasesLoaded'
    })
  },
  methods: {
    ...mapActions('medicalCase', {
      getMedicalCases: 'get_all'
    })
  }
})
export default class MedicalCases extends Vue {
  getMedicalCases!: () => void
  medicalCasesLoaded!: boolean
  medicalCasesList!: object

  loading = false

  loadMedicalCases() {
    this.loading = true
    this.getMedicalCases()
    this.loading = false
  }

  mounted() {
    if (!this.medicalCasesLoaded) {
      this.loadMedicalCases()
    }
  }
}
</script>
