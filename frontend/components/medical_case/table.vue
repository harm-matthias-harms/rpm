<template>
  <div>
    <v-text-field v-model="search" append-icon="search" label="Search" single-line hide-details />
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
              {{ item.title }}
            </td>
            <td @click="openMedicalCase(item)">
              <v-chip v-for="(tag,i) in tags(item)" :key="i" class="ma-1">
                {{ tag }}
              </v-chip>
            </td>
            <td @click="openMedicalCase(item)">
              {{ item.author.username }}
            </td>
            <td class="px-0">
              <v-icon @click="editMedicalCase(item)">
                edit
              </v-icon>
              <DeleteButton
                v-if="!item.author.username || item.author.id == $store.state.user.user.id"
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
    { text: 'Tags', width: '40%', sortable: false, value: 'tags' },
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

  tags (medicalCase) {
    const tags: string[] = []
    if (medicalCase.generalInformation.usar) {
      tags.push('USAR')
    }
    if (medicalCase.generalInformation.medivac) {
      tags.push('MEDIVAC')
    }
    if (medicalCase.generalInformation.hospilisation) {
      tags.push('Need for hospilisation')
    }
    if (medicalCase.generalInformation.surgical) {
      tags.push('Surgical')
    }
    if (medicalCase.generalInformation.triage) {
      tags.push('Triage: ' + medicalCase.generalInformation.triage)
    }
    if (medicalCase.generalInformation.age) {
      tags.push('Age: ' + medicalCase.generalInformation.age)
    }
    if (medicalCase.generalInformation.gender) {
      tags.push('Gender: ' + medicalCase.generalInformation.gender)
    }
    return tags
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
        filterUSAR(item, search) ||
        filterMedivac(item, search) ||
        filterHospital(item, search) ||
        filterSurgical(item, search) ||
        filterAge(item, search) ||
        filterGender(item, search) ||
        filterTriage(item, search)
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

    function filterSurgical (item, search) {
      if ('surgical'.includes(search)) {
        return item.generalInformation.surgical
      }
      return false
    }

    function filterHospital (item, search) {
      if ('hospital'.includes(search)) {
        return item.generalInformation.hospilisation
      }
      return false
    }
    function filterUSAR (item, search) {
      if ('usar'.includes(search)) {
        return item.generalInformation.usar
      }
      return false
    }
    function filterMedivac (item, search) {
      if ('medivac'.includes(search)) {
        return item.generalInformation.medivac
      }
      return false
    }
    function filterAge (item, search) {
      const ageSpans = ['0-10', '11-17', '18-30', '31-60', '60+']
      if (item.generalInformation.age == null) {
        return false
      }
      for (let i = 0; i < ageSpans.length; i++) {
        const arr = ageSpans[i].split('-')
        if (arr.length > 1) {
          const min = arr[0]
          const max = arr[1]
          if (
            parseInt(search) >= parseInt(min) &&
            parseInt(search) <= parseInt(max)
          ) {
            return item.generalInformation.age === ageSpans[i]
          }
        } else if (arr[0] === '60+' && parseInt(search) >= 60) {
          return item.generalInformation.age === '60+'
        }
      }
    }

    function filterGender (item, search) {
      return item.generalInformation.gender.toLowerCase().includes(search)
    }

    function filterTriage (item, search) {
      const triages = ['Less Urgent', 'Urgent', 'Emergent']
      if (triages.some(triage => triage.toLowerCase().includes(search))) {
        return item.generalInformation.triage.toLowerCase().includes(search)
      }
      return false
    }
    return (
      value && search && needleArry.every(needle => filterAll(item, needle))
    )
  }
}
</script>
