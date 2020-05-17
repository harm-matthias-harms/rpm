<template>
  <v-expansion-panel>
    <v-expansion-panel-header>{{ vitalSign.title ? vitalSign.title : 'No title' }}</v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-select
        v-model="vitalSign.title"
        :items="['T0 - Start', 'T1 - Improvement', 'T2 - Deterioration']"
        label="Title"
      />
      <v-autocomplete
        v-model="presetID"
        :items="presets.presets"
        :loading="loading"
        :item-text="item => item.title + ' - ' + item.author.username"
        item-value="id"
        label="Select a preset"
      >
        <template v-slot:item="data">
          <v-list-item-content>
            <v-list-item-title>{{ data.item.title }}</v-list-item-title>
            <v-list-item-subtitle>{{ data.item.author.username }}</v-list-item-subtitle>
          </v-list-item-content>
        </template>
      </v-autocomplete>
      <Form :vital-signs.sync="vitalSign.data" />
      <v-btn
        class="mb-4"
        @click="addChild"
      >
        Add Step
      </v-btn>
      <v-expansion-panels
        multiple
        class="mb-4"
      >
        <VitalSign
          v-for="(vs, i) in vitalSign.childs"
          :key="i"
          :vital-sign.sync="vitalSign.childs[i]"
        />
      </v-expansion-panels>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Form from '@/components/vital_signs/form.vue'
@Component({
  name: 'VitalSign',
  components: { Form },
  computed: {
    ...mapState('preset', {
      presetsLoaded: 'presetsLoaded',
      presets: 'presetList',
      preset: 'preset'
    })
  },
  methods: {
    ...mapActions('preset', {
      getPresets: 'get_all',
      findPreset: 'find'
    })
  }
})
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSign!: any
  @Watch('vitalSign', { immediate: true, deep: true })
  updateVitalSignChanged (val: any) {
    this.$emit('update:vitalSign', val)
  }

  @Watch('presetID')
  presetIDChanged (val: string) {
    if (val) {
      this.findPreset({ id: val, disableLoader: true }).then((res) => {
        this.vitalSign.data = res.vitalSigns
      })
    }
  }

  presetsLoaded!: boolean
  getPresets!: (payload) => void
  findPreset!: (payload) => Promise<any>
  presets!: []
  preset: any
  loading = false
  presetID: string = ''

  emptyVitalSign: object = {
    title: undefined,
    data: {
      oos: undefined,
      avpu: undefined,
      mobility: undefined,
      respiratoryRate: undefined,
      pulse: undefined,
      temperature: undefined,
      capillaryRefill: undefined,
      bloodPressureSystolic: undefined,
      bloodPressureDiastolic: undefined,
      oxygenSaturation: undefined,
      expectations: {
        foe: undefined,
        treatmentExpected: undefined
      }
    },
    childs: []
  }

  addChild () {
    this.vitalSign.childs.push(Object.assign({}, this.emptyVitalSign))
  }

  loadPresets () {
    this.loading = true
    this.getPresets({ disableLoader: true })
    this.loading = false
  }

  mounted () {
    if (!this.presetsLoaded) {
      this.loadPresets()
    }
  }
}
</script>
