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
      :headers="headers"
      :items="items"
      :items-per-page="50"
      :footer-props="{
        itemsPerPageOptions: [10, 25, 50, 100],
      }"
      :search="search"
      :custom-filter="filterInjects"
      class="elevation-1"
      :loading="loading"
      loading-text="Loading... Please wait"
      multi-sort
    >
      <template v-slot:[`item.status`]="{ item }">
        <v-badge dot inline left color="primary">
          {{ item.status }}
        </v-badge>
      </template>
      <template v-slot:[`item.medicalCase.title`]="{ item }">
        <v-badge dot inline left :color="getTriageColor(item.medicalCase)">
          {{ item.medicalCase.title }}
        </v-badge>
      </template>
      <template v-slot:[`item.startTime`]="{ item }">
        {{
          item.startTime
            ? new Date(item.startTime).toLocaleString([], {
              timeZone: 'UTC',
            })
            : '-'
        }}
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { Inject } from '~/store/inject/type'
import { MedicalCase } from '~/store/medicalCase/type'
// import DeleteButton from '@/components/medical_case/Delete.vue'
@Component({
  // components: {
  //   DeleteButton
  // }
})
export default class Table extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>

  headers = [
    { text: 'State', sortable: true, value: 'status' },
    { text: 'Roleplayer', sortable: true, value: 'roleplayer.fullName' },
    { text: 'Medical case', sortable: true, value: 'medicalCase.title' },
    { text: 'Team', sortable: true, value: 'team.title' },
    { text: 'Make-up center', sortable: true, value: 'makeupCenterTitle' },
    { text: 'Start time', sortable: true, value: 'startTime' },
    { text: 'Actions', sortable: false, value: 'action' }
  ]

  search: string = ''

  openInject (inject) {
    this.$router.push(
      `/exercises/${this.$route.params.id}/injects/${inject.id}`
    )
  }

  editInject (inject) {
    this.$router.push(
      `/exercises/${this.$route.params.id}/injects/${inject.id}/edit`
    )
  }

  getTriageColor (medicalCase: MedicalCase) {
    if (!medicalCase.patient.triage) {
      return 'grey'
    }
    if (medicalCase.patient.triage === 'Red') {
      return 'error'
    }
    if (medicalCase.patient.triage === 'Yellow') {
      return 'warning'
    }
    if (medicalCase.patient.triage === 'Green') {
      return 'success'
    }
    if (medicalCase.patient.triage === 'Deceased/Unsalvageable') {
      return 'black'
    }
  }

  filterInjects (value, search, item: Inject) {
    // Split search into an array
    const needleArry = search
      .toString()
      .toLowerCase()
      .split(' ')
      .filter(x => x.trim().length > 0)

    function filterAll (item: Inject, search: string) {
      return (
        filterStatus(item, search) ||
        filterRoleplayer(item, search) ||
        filterMedicalCase(item, search) ||
        filterTeam(item, search) ||
        filterMakeupCenter(item, search) ||
        filterStartTime(item, search)
      )
    }

    function filterStatus (item: Inject, search: string) {
      return item.status?.toLowerCase().includes(search)
    }

    function filterRoleplayer (item: Inject, search: string) {
      return item.roleplayer.fullName?.toLowerCase().includes(search)
    }

    function filterMedicalCase (item: Inject, search: string) {
      return item.medicalCase.title?.toLowerCase().includes(search)
    }

    function filterTeam (item: Inject, search: string) {
      return item.team?.title?.toLowerCase().includes(search)
    }

    function filterMakeupCenter (item: Inject, search: string) {
      return item.makeupCenterTitle?.toLowerCase().includes(search)
    }

    function filterStartTime (item: Inject, search: string) {
      return item.startTime?.toString().toLowerCase().includes(search)
    }

    return (
      value && search && needleArry.every(needle => filterAll(item, needle))
    )
  }
}
</script>
