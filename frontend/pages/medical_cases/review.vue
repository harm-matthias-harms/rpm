<template>
  <v-container>
    <v-row justify="center">
      <v-col
        md="5"
        sm="6"
      >
        <h3 class="headline">
          <v-chip
            small
            label
            color="primary"
          >
            {{ items().length }}
          </v-chip>Medical Cases
        </h3>
      </v-col>
      <v-col
        md="5"
        sm="6"
        class="justify-end d-flex"
      >
        <v-btn
          small
          color="primary"
          class="mr-2"
          to="/medical_cases/new"
        >
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
        <v-btn
          small
          color="primary"
        >
          <v-icon
            small
            @click="loadMedicalCases"
          >
            fas fa-redo
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col
        md="10"
        sm="12"
      >
        <Table
          :items="items()"
          :loading="loading"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapGetters, mapState, mapActions } from 'vuex'
import Table from '@/components/medical_case/table.vue'
@Component({
  components: {
    Table
  },
  computed: {
    ...mapGetters('medicalCase', {
      items: 'requireReview'
    }),
    ...mapState('medicalCase', {
      medicalCasesLoaded: 'medicalCasesLoaded'
    }),
    ...mapState('user', {
      user: 'user'
    })
  },
  methods: {
    ...mapActions('medicalCase', {
      getMedicalCases: 'get_all'
    })
  }
})
export default class ReviewMedicalCases extends Vue {
  getMedicalCases!: () => void
  medicalCasesLoaded!: boolean
  loading = false
  loadMedicalCases () {
    this.loading = true
    this.getMedicalCases()
    this.loading = false
  }

  mounted () {
    if (!this.medicalCasesLoaded) {
      this.loadMedicalCases()
    }
  }
}
</script>
