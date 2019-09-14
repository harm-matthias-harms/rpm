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
  />
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class PresetTable extends Vue {
  headers = [
    { text: 'Title', align: 'left', sortable: true, value: 'title' },
    { text: 'Author', sortable: true, value: 'author.username' }
  ]

  @Prop({ type: Boolean, required: true }) readonly loading!: boolean
  @Prop({ type: Array, required: true }) readonly items!: Array<object>
}
</script>
