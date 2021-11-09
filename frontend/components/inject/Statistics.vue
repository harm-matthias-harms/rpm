<template>
  <div>
    <h3 class="headline">Statistics</h3>
    <div
      v-for="[date, teams] in Object.entries(countByStatus)"
      :key="date"
      class="mt-10"
    >
      <h4>{{ date }}</h4>
      <v-data-table
        :items="teams"
        :headers="headers"
        item-key="name"
        :items-per-page="-1"
        sort-by="name"
        hide-default-footer
        class="elevation-1 mt-5"
      >
        <template v-slot:item="{ item }">
          <tr>
            <td>{{ item.name }}</td>
            <td>{{ item['Waiting for makeup'] || 0 }}</td>
            <td>{{ item['In makeup'] || 0 }}</td>
            <td>{{ item['Ready to play'] || 0 }}</td>
            <td>{{ item['Playing'] || 0 }}</td>
            <td>{{ item['Finished'] || 0 }}</td>
          </tr>
        </template>
      </v-data-table>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { groupBy, countBy } from 'lodash'

import { InjectShort } from '~/store/inject/type'
@Component
export default class Statistics extends Vue {
  @Prop({ type: Array, required: true }) readonly items!: InjectShort[]

  headers = [
    { text: 'Team', sortable: true, value: 'name' },
    { text: 'Waiting for makeup', sortable: true, value: 'Waiting for makeup' },
    { text: 'In makeup', sortable: true, value: 'In makeup' },
    { text: 'Ready to play', sortable: true, value: 'Ready to play' },
    { text: 'Playing', sortable: true, value: 'Playing' },
    { text: 'Finished', sortable: true, value: 'Finished' },
  ]

  get groupedByDay() {
    const grouped = groupBy(this.items, (item: InjectShort) => {
      return new Date(item.startTime!).toISOString().slice(0, 10)
    })
    const ordered = {}
    Object.keys(grouped)
      .sort()
      .forEach((key) => (ordered[key] = grouped[key]))
    return ordered
  }

  get groupedByTeam() {
    const group = this.groupedByDay
    for (const [key, items] of Object.entries(group)) {
      group[key] = groupBy(items, (item: InjectShort) => {
        return item.team.title
      })
    }

    return group
  }

  get countByStatus() {
    const group = this.groupedByTeam
    for (const [key, teams] of Object.entries(group)) {
      group[key] = []
      for (const [team, items] of Object.entries(new Object(teams))) {
        group[key].push({
          name: team,
          ...countBy(items, (item: InjectShort) => {
            return item.status
          }),
        })
      }
    }
    return group
  }
}
</script>
