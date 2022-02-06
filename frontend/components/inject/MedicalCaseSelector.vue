<template>
  <div>
    <v-row dense>
      <v-col>
        <v-text-field
          v-model="search"
          append-icon="search"
          label="Search"
          single-line
          hide-details
        />
      </v-col>
      <v-col cols="2">
        <v-select
          clearable
          label="Triage"
          v-model="filters.triage"
          :items="['Deceased/Unsalvageable', 'Red', 'Yellow', 'Green']"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          clearable
          label="Area / Discipline"
          v-model="filters.area"
          :items="[
            'Internal med',
            'Surgery',
            'Gyn/Obs',
            'Infectious diseases',
            'Trauma',
            'Public health',
          ]"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          clearable
          label="Scenario"
          v-model="filters.scenario"
          :items="[
            'Conflict',
            'Natural disaster',
            'Man-made disaster',
            'Mass-Casualty-Incident',
            'Outbreak',
            'People displacement/Refugees',
          ]"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          clearable
          label="Pre hospital"
          v-model="filters.preHospital"
          :items="[true, false]"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          clearable
          label="MEDEVAC"
          v-model="filters.medevac"
          :items="[true, false]"
        />
      </v-col>
    </v-row>
    <v-data-table
      v-model="medicalCases"
      show-select
      :headers="headers"
      :items="filteredItems"
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
      <template v-slot:[`item.general.medevac`]="{ item }">
        <v-icon :color="item.general.medevac ? 'success' : 'error'">
          {{ item.general.medevac ? 'fa-check-circle' : 'fa-times-circle' }}
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
    { text: 'Pre Hospital', sortable: true, value: 'general.preHospital' },
    { text: 'Medevac', sortable: true, value: 'general.medevac' },
  ]

  search: string = ''

  filters = {
    triage: '',
    area: '',
    scenario: '',
    preHospital: undefined,
    medevac: undefined,
  }

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

  get filteredItems(): MedicalCaseShort[] {
    let items = this.options

    if (this.filters.triage) {
      items = items.filter((item) => {
        return item.patient.triage === this.filters.triage
      })
    }

    if (this.filters.area) {
      items = items.filter((item) => {
        return item.general.discipline?.includes(this.filters.area)
      })
    }

    if (this.filters.scenario) {
      items = items.filter((item) => {
        return item.general.scenario?.includes(this.filters.scenario)
      })
    }

    if ([true, false].includes(this.filters.preHospital!)) {
      items = items.filter((item) => {
        return item.general.preHospital === this.filters.preHospital
      })
    }

    if ([true, false].includes(this.filters.medevac!)) {
      items = items.filter((item) => {
        return item.general.medevac === this.filters.medevac
      })
    }

    return items
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
        ('prehospital'.includes(search) && item.general.preHospital) ||
        ('medevac'.includes(search) && item.general.medevac)
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
