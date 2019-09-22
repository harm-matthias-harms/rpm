<template>
  <v-container>
    <v-row justify="center">
      <v-col md="4" sm="6">
        <h3 class="headline">
          <v-chip small label color="primary">
            {{ presetList.count }}
          </v-chip>Presets
        </h3>
      </v-col>
      <v-col md="4" sm="6" class="justify-end d-flex">
        <v-btn small color="primary" class="mr-2" to="/presets/new">
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
        <v-btn small color="primary" @click="getPresets()">
          <v-icon small>
            fas fa-redo
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="8" sm="12">
        <PresetTable :items="presetList.presets" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import PresetTable from '@/components/preset/table.vue'
@Component({
  components: {
    PresetTable
  },
  computed: {
    ...mapState('preset', {
      presetList: 'presetList',
      presetsLoaded: 'presetsLoaded'
    })
  },
  methods: {
    ...mapActions('preset', {
      getPresets: 'get_all'
    })
  }
})
export default class Presets extends Vue {
  getPresets!: () => void
  presetsLoaded!: boolean

  loading = false

  loadPresets () {
    this.loading = true
    this.getPresets()
    this.loading = false
  }

  mounted () {
    if (!this.presetsLoaded) {
      this.loadPresets()
    }
  }
}
</script>
