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
      :items-per-page="25"
      :footer-props="{
        itemsPerPageOptions: [10, 25, 50],
      }"
      :search="search"
      :custom-filter="filterMedicalCases"
      class="elevation-1"
      :loading="loading"
      loading-text="Loading... Please wait"
      multi-sort
    >
      <template v-slot:body="{ items }">
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td @click="openMedicalCase(item)">
              <v-badge dot inline left :color="getTriageColor(item)">
                {{ item.title }}
              </v-badge>
            </td>
            <td @click="openMedicalCase(item)">
              {{
                item.general.discipline
                  ? item.general.discipline.join(', ')
                  : ''
              }}
            </td>
            <td @click="openMedicalCase(item)">
              {{ item.general.context ? item.general.context.join(', ') : '' }}
            </td>
            <td @click="openMedicalCase(item)">
              {{
                item.general.scenario ? item.general.scenario.join(', ') : ''
              }}
            </td>
            <td @click="openMedicalCase(item)">
              {{ item.author.username }}
            </td>
            <td class="px-0">
              <v-icon @click="editMedicalCase(item)">
                edit
              </v-icon>
              <DeleteButton
                v-if="
                  !item.author.username ||
                    item.author.id == $store.state.user.user.id
                "
                :item="item"
                :go-back="false"
              />
            </td>
          </tr>
        </tbody>
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import DeleteButton from '@/components/medical_case/Delete.vue'
@Component({
  components: {
    DeleteButton
  }
})
export default class Table extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>

  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Area', sortable: true, value: 'general.area' },
    { text: 'Context', sortable: true, value: 'general.context' },
    { text: 'Scenario', sortable: true, value: 'general.scenario' },
    { text: 'Author', sortable: true, value: 'author.username' },
    { text: 'Actions', sortable: false, value: 'action' }
  ]

  search: string = ''

  openMedicalCase (medicalCase) {
    this.$router.push('/medical_cases/' + medicalCase.id)
  }

  editMedicalCase (medicalCase) {
    this.$router.push('/medical_cases/' + medicalCase.id + '/edit')
  }

  getTriageColor (item) {
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

  filterMedicalCases (value, search, item) {
    // Split search into an array
    const needleArry = search
      .toString()
      .toLowerCase()
      .split(' ')
      .filter(x => x.trim().length > 0)

    function filterAll (item, search) {
      return (
        filterTitle(item, search) ||
        filterAuthor(item, search) ||
        filterGeneral(item, search) ||
        filterPatient(item, search)
      )
    }

    function filterTitle (item, search) {
      return item.title && item.title.toLowerCase().includes(search)
    }

    function filterAuthor (item, search) {
      return (
        item.author.username &&
        item.author.username.toLowerCase().includes(search)
      )
    }

    function filterGeneral (item, search) {
      return (
        item.general.discipline?.some(e =>
          e.toLowerCase().includes(search)
        ) ||
        item.general.context?.some(e => e.toLowerCase().includes(search)) ||
        item.general.scenario?.some(e => e.toLowerCase().includes(search))
      )
    }

    function filterPatient (item, search) {
      return (
        item.patient.triage &&
        item.patient.triage.toLowerCase().includes(search)
      )
    }

    return (
      value && search && needleArry.every(needle => filterAll(item, needle))
    )
  }
}
</script>
