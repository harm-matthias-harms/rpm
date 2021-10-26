<template>
  <div>
    <v-text-field
      v-model="search"
      append-icon="search"
      label="Search"
      single-line
      hide-details
    />
    <v-data-table
      v-model="medicalCases"
      show-select
      :headers="headers"
      :items="options"
      :items-per-page="25"
      :footer-props="{
        itemsPerPageOptions: [10, 25, 50],
      }"
      :search="search"
      :custom-filter="filterMedicalCases"
      multi-sort
    >
      <template v-slot:[`item.title`]="{ item }">
        <v-badge dot inline left :color="getTriageColor(item)">
          {{ item.title }}
        </v-badge>
      </template>
      <template v-slot:[`item.general.discipline`]="{ item }">
        {{ item.general.discipline ? item.general.discipline.join(', ') : '' }}
      </template>
      <template v-slot:[`item.general.context`]="{ item }">
        {{ item.general.context ? item.general.context.join(', ') : '' }}
      </template>
      <template v-slot:[`item.general.scenario`]="{ item }">
        {{ item.general.scenario ? item.general.scenario.join(', ') : '' }}
      </template>
      <template v-slot:[`item.general.preHospital`]="{ item }">
        <v-icon :color="item.general.preHospital ? 'success' : 'error'">
          {{ item.general.preHospital ? 'fa-check-circle' : 'fa-times-circle' }}
        </v-icon>
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import { MedicalCaseShort } from '~/store/medicalCase/type'

@Component({
  methods: {
    ...mapActions('medicalCase', {
      getMedicalCases: 'get_all',
    }),
  },
})
export default class MedicalCaseSelector extends Vue {
  @Prop({ type: Array }) readonly value!: MedicalCaseShort[]

  getMedicalCases!: () => Promise<void>

  get medicalCases(): MedicalCaseShort[] {
    return this.value
  }

  set medicalCases(medicalCases: MedicalCaseShort[]) {
    this.$emit('input', medicalCases)
  }

  options: MedicalCaseShort[] = []
  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Area', sortable: true, value: 'general.discipline' },
    { text: 'Context', sortable: true, value: 'general.context' },
    { text: 'Scenario', sortable: true, value: 'general.scenario' },
    { text: 'PreHospital', sortable: true, value: 'general.preHospital' },
  ]

  search: string = ''

  mounted() {
    if (this.$store.state.medicalCase.medicalCasesLoaded) {
      this.options = this.$store.state.medicalCase.medicalCasesList.medicalCases
    } else {
      this.getMedicalCases().then(() => {
        this.options =
          this.$store.state.medicalCase.medicalCasesList.medicalCases
      })
    }
  }

  getTriageColor(item) {
    if (!item.patient.triage) {
      return 'grey'
    }
    if (item.patient.triage === 'Red') {
      return 'error'
    }
    if (item.patient.triage === 'Yellow') {
      return 'warning'
    }
    if (item.patient.triage === 'Green') {
      return 'success'
    }
    if (item.patient.triage === 'Deceased/Unsalvageable') {
      return 'black'
    }
  }

  filterMedicalCases(value, search, item) {
    // Split search into an array
    const needleArry = search
      .toString()
      .toLowerCase()
      .split(' ')
      .filter((x) => x.trim().length > 0)

    function filterAll(item, search) {
      return (
        filterTitle(item, search) ||
        filterAuthor(item, search) ||
        filterGeneral(item, search) ||
        filterPatient(item, search)
      )
    }

    function filterTitle(item, search) {
      return item.title && item.title.toLowerCase().includes(search)
    }

    function filterAuthor(item, search) {
      return (
        item.author.username &&
        item.author.username.toLowerCase().includes(search)
      )
    }

    function filterGeneral(item, search) {
      return (
        item.general.discipline?.some((e) =>
          e.toLowerCase().includes(search)
        ) ||
        item.general.context?.some((e) => e.toLowerCase().includes(search)) ||
        item.general.scenario?.some((e) => e.toLowerCase().includes(search)) ||
        ('prehospital'.includes(search) && item.general.preHospital)
      )
    }

    function filterPatient(item, search) {
      return (
        item.patient.triage &&
        item.patient.triage.toLowerCase().includes(search)
      )
    }

    return (
      value && search && needleArry.every((needle) => filterAll(item, needle))
    )
  }
}
</script>
