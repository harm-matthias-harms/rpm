<template>
  <v-container>
    <v-row justify="center">
      <v-col md="4" sm="6">
        <h3 class="headline">
          <v-chip small label color="primary">
            {{ count(user.username) }}
          </v-chip>Presets
        </h3>
      </v-col>
      <v-col md="4" sm="6" class="justify-end d-flex">
        <v-btn small color="primary" class="mr-2">
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
        <v-btn small color="primary">
          <v-icon small @click="loadPresets">
            fas fa-redo
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="8" sm="12">
        <PresetTable :items="items(user.username)" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapGetters, mapState, mapActions } from 'vuex'
import PresetTable from '@/components/preset/table.vue'
@Component({
  components: {
    PresetTable
  },
  computed: {
    ...mapGetters('preset', {
      items: 'myOwn',
      count: 'myOwnCount'
    }),
    ...mapState('preset', {
      presetsLoaded: 'presetsLoaded'
    }),
    ...mapState('user', {
      user: 'user'
    })
  },
  methods: {
    ...mapActions('preset', {
      getPresets: 'get_all'
    })
  }
})
export default class OwnPresets extends Vue {
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
