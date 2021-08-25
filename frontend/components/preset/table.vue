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
        itemsPerPageOptions: [25, 50, 100],
      }"
      :search="search"
      :custom-filter="filterPresets"
      class="elevation-1"
      :loading="loading"
      loading-text="Loading... Please wait"
      multi-sort
    >
      <template v-slot:body="{ items }">
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td @click="openPreset(item)">
              {{ item.title }}
            </td>
            <td @click="openPreset(item)">
              {{ item.author.username }}
            </td>
            <td class="px-0">
              <v-icon @click="editPreset(item)">
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
import DeleteButton from '@/components/preset/Delete.vue'
@Component({
  components: { DeleteButton }
})
export default class PresetTable extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>

  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Author', sortable: true, value: 'author.username' }
  ]

  search: string = ''

  openPreset (preset) {
    this.$router.push('/presets/' + preset.id)
  }

  editPreset (preset) {
    this.$router.push('/presets/' + preset.id + '/edit')
  }

  filterPresets (value, search, item) {
    // Split search into an array
    const needleArry = search
      .toString()
      .toLowerCase()
      .split(' ')
      .filter(x => x.trim().length > 0)

    function filterAll (item, search) {
      return filterTitle(item, search) || filterAuthor(item, search)
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

    return (
      value && search && needleArry.every(needle => filterAll(item, needle))
    )
  }
}
</script>
