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
          v-model="filters.status"
          :items="states"
          clearable
          label="State"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          v-model="filters.team"
          :items="selectableTeams"
          clearable
          label="Team"
        />
      </v-col>
      <v-col cols="2">
        <v-select
          v-model="filters.makeupCenter"
          :items="selectableMakeupCenter"
          clearable
          label="Make-up center"
        />
      </v-col>
      <v-col cols="auto">
        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              :disabled="!selectedInjects.length"
              v-bind="attrs"
              v-on="on"
            >
              Print
              {{
                selectedInjects.length ? `(${selectedInjects.length})` : null
              }}
              <v-icon small right>fas fa-chevron-down</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item
              :to="{
                path: `/exercises/${$route.params.id}/injects/print/vital_signs`,
                query: {
                  injects: selectedInjects.map((inject) => inject.id),
                },
              }"
              target="_blank"
            >
              <v-list-item-title> Vital signs </v-list-item-title>
            </v-list-item>
            <v-list-item
              :to="{
                path: `/exercises/${$route.params.id}/injects/print/makeup`,
                query: {
                  injects: selectedInjects.map((inject) => inject.id),
                },
              }"
              target="_blank"
            >
              <v-list-item-title> Makeup</v-list-item-title>
            </v-list-item>
            <v-list-item
              :to="{
                path: `/exercises/${this.$route.params.id}/injects/print/roleplayer`,
                query: {
                  injects: selectedInjects.map((inject) => inject.id),
                },
              }"
              target="_blank"
            >
              <v-list-item-title>Roleplayer</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-col>
    </v-row>
    <v-data-table
      v-model="selectedInjects"
      :headers="headers"
      :items="itemsForRole"
      :items-per-page="50"
      :footer-props="{
        itemsPerPageOptions: [10, 25, 50, 100],
      }"
      :search="search"
      :custom-filter="filterInjects"
      class="elevation-1"
      :loading="loading"
      loading-text="Loading... Please wait"
      show-select
      multi-sort
    >
      <template v-slot:[`item.status`]="{ item }">
        <v-select
          dense
          :items="states"
          :value="item.status"
          @change="
            (selected) => {
              updateStatus(selected, item)
            }
          "
        >
          <template v-slot:selection="{ item }">
            <v-badge dot inline left :color="getStatusColor(item)">
              {{ item }}
            </v-badge>
          </template>
          <template v-slot:item="{ item }">
            <v-badge dot inline left :color="getStatusColor(item)">
              {{ item }}
            </v-badge>
          </template>
        </v-select>
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
      <template v-slot:[`item.action`]="{ item }">
        <v-icon @click="openMedicalCase(item)">
          fas fa-external-link-alt
        </v-icon>
        <DeleteButton
          v-if="isAdminOrRpm && item.status === states[0]"
          :item="item"
          :go-back="false"
          :exercise-id="$route.params.id"
        />
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import DeleteButton from './Delete.vue'
import { Inject, InjectShort } from '~/store/inject/type'
import { MedicalCase } from '~/store/medicalCase/type'
@Component({
  components: {
    DeleteButton,
  },
  methods: {
    ...mapActions('inject', {
      findInject: 'find',
      updateInject: 'update',
    }),
  },
})
export default class Table extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: String, required: true }) readonly exerciseID!: string
  @Prop({ type: Array, required: true }) readonly items!: InjectShort[]
  @Prop({ type: Function, required: true }) readonly refreshTable!: ({
    exerciseID,
  }) => void

  findInject!: ({ id, exerciseID, disableLoader }) => void
  updateInject!: ({ inject, showInject }) => void

  headers = [
    { text: 'State', sortable: true, value: 'status', width: '220px' },
    { text: 'Roleplayer', sortable: true, value: 'roleplayer.fullName' },
    { text: 'Medical case', sortable: true, value: 'medicalCase.title' },
    { text: 'Team', sortable: true, value: 'team.title' },
    { text: 'Make-up center', sortable: true, value: 'makeupCenterTitle' },
    { text: 'Start time', sortable: true, value: 'startTime' },
    { text: 'Actions', sortable: false, value: 'action' },
  ]

  search: string = ''
  filters = {
    status: '',
    team: '',
    makeupCenter: '',
  }
  states: string[] = [
    'Waiting for makeup',
    'In makeup',
    'Ready to play',
    'Playing',
    'Finished',
  ]

  selectedInjects: InjectShort[] = []

  get exerciseRoles() {
    return this.$store.state.user.user.roles.filter(
      (role) => role.exercise.id === this.exerciseID
    )
  }

  get itemsForRole(): InjectShort[] {
    const isMakeupCenter =
      this.exerciseRoles.some((role) => role.role === 'make-up center') &&
      !this.exerciseRoles.some((role) =>
        ['admin', 'role play manager', 'trainer'].includes(role.role)
      )

    let filteredItems = this.items
    if (this.filters.status) {
      filteredItems = filteredItems.filter(
        (item) => item.status === this.filters.status
      )
    }
    if (this.filters.team) {
      filteredItems = filteredItems.filter(
        (item) => item.team.title === this.filters.team
      )
    }
    if (this.filters.makeupCenter) {
      filteredItems = filteredItems.filter(
        (item) => item.makeupCenterTitle === this.filters.makeupCenter
      )
    }

    if (isMakeupCenter) {
      return filteredItems.filter((item) =>
        ['Waiting for makeup', 'In makeup', 'Ready to play'].includes(
          item.status!
        )
      )
    }

    return filteredItems
  }

  get selectableTeams() {
    return [...new Set(this.items.map((item) => item.team.title))]
  }

  get selectableMakeupCenter() {
    return [...new Set(this.items.map((item) => item.makeupCenterTitle))]
  }

  get isAdminOrRpm() {
    return this.exerciseRoles.some((role) =>
      ['admin', 'role play manager'].includes(role.role)
    )
  }

  openMedicalCase(inject: InjectShort) {
    window.open('/medical_cases/' + inject.medicalCase.id, '_blank')
  }

  async updateStatus(newStatus: string, inject: InjectShort) {
    await this.findInject({
      id: inject.id,
      exerciseID: this.$route.params.id,
      disableLoader: true,
    })

    const updatedInject: Inject = Object.assign(
      {},
      this.$store.state.inject.inject
    )
    updatedInject.status = newStatus

    await this.updateInject({ inject: updatedInject, showInject: false })
    await this.refreshTable({ exerciseID: this.$route.params.id })
    this.updateSelected()
  }

  updateSelected() {
    this.selectedInjects = this.selectedInjects
      ?.map(
        (inject) =>
          this.itemsForRole?.find((item: InjectShort) => item.id === inject.id)!
      )
      .filter((e) => e)!
  }

  getStatusColor(status: string): string {
    const colors = {
      'Waiting for makeup': 'error',
      'In makeup': 'primary',
      'Ready to play': 'warning',
      Playing: 'success',
      Finished: 'black',
    }
    if (status in colors) {
      return colors[status]
    }
    return 'gray'
  }

  getTriageColor(medicalCase: MedicalCase): String {
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
    return 'grey'
  }

  filterInjects(value, search, item: InjectShort) {
    // Split search into an array
    const needleArry = search
      .toString()
      .toLowerCase()
      .split(' ')
      .filter((x) => x.trim().length > 0)

    function filterAll(item: InjectShort, search: string) {
      return (
        filterStatus(item, search) ||
        filterRoleplayer(item, search) ||
        filterMedicalCase(item, search) ||
        filterTeam(item, search) ||
        filterMakeupCenter(item, search) ||
        filterStartTime(item, search)
      )
    }

    function filterStatus(item: InjectShort, search: string) {
      return item.status?.toLowerCase().includes(search)
    }

    function filterRoleplayer(item: InjectShort, search: string) {
      return item.roleplayer.fullName?.toLowerCase().includes(search)
    }

    function filterMedicalCase(item: InjectShort, search: string) {
      return item.medicalCase.title?.toLowerCase().includes(search)
    }

    function filterTeam(item: InjectShort, search: string) {
      return item.team?.title?.toLowerCase().includes(search)
    }

    function filterMakeupCenter(item: InjectShort, search: string) {
      return item.makeupCenterTitle?.toLowerCase().includes(search)
    }

    function filterStartTime(item: InjectShort, search: string) {
      return item.startTime?.toString().toLowerCase().includes(search)
    }

    return (
      value && search && needleArry.every((needle) => filterAll(item, needle))
    )
  }
}
</script>
