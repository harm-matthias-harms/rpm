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
      :custom-filter="filterTeams"
      class="elevation-1"
      :loading="loading"
      loading-text="Loading... Please wait"
      multi-sort
    >
      <template v-slot:body="{ items }">
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td @click="openTeam(item)">
              {{ item.title }}
            </td>
            <td @click="openTeam(item)">
              {{ item.type }}
            </td>
            <td @click="openTeam(item)">
              <v-icon v-if="item.medivac" small>
                done
              </v-icon>
              <v-icon v-if="!item.medivac" small>
                clear
              </v-icon>
            </td>
            <td @click="openTeam(item)">
              {{ item.author.username }}
            </td>
            <td class="px-0">
              <v-icon @click="editTeam(item)">
                edit
              </v-icon>
              <DeleteButton
                v-if="item.author.id == $store.state.user.user.id"
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
import DeleteButton from '@/components/team/Delete.vue'

@Component({
  components: { DeleteButton }
})
export default class TeamTable extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>

  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Type', sortable: true, value: 'type' },
    { text: 'Medivac', sortable: true, value: 'medivac' },
    { text: 'Author', sortable: true, value: 'author.username' }
  ]

  search: string = ''

  openTeam (team) {
    this.$router.push('/teams/' + team.id)
  }

  editTeam (team) {
    this.$router.push('/teams/' + team.id + '/edit')
  }

  filterTeams (value, search, item) {
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
        filterType(item, search) ||
        filterMedivac(item, search)
      )
    }

    function filterTitle (item, search) {
      return item.title && item.title.toLowerCase().includes(search)
    }

    function filterType (item, search) {
      return item.type && item.type.toLowerCase().includes(search)
    }

    function filterMedivac (item, search) {
      if ('medivac'.includes(search)) {
        return item.medivac
      }
      return false
    }

    function filterAuthor (item, search) {
      return (
        item.author.username &&
        item.author.username.toLowerCase().includes(search)
      )
    }

    return (
      value && search && needleArry.every(needle => filterAll(item, needle))
    )
  }
}
</script>
