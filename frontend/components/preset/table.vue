<template>
  <v-data-table
    :headers="headers"
    :items="items"
    :items-per-page="25"
    :footer-props="{
      itemsPerPageOptions: [10, 25, 50],
    }"
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
        </tr>
      </tbody>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class PresetTable extends Vue {
  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>

  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Author', sortable: true, value: 'author.username' }
  ]

  openPreset (preset) {
    this.$router.push('/presets/' + preset.id)
  }
}
</script>
