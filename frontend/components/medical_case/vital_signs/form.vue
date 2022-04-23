<template>
  <v-expansion-panel>
    <v-expansion-panel-header>
      <div>
        <v-chip :color="levelColor()" class="mr-2">{{ level }}</v-chip>
        {{ vitalSign.title ? vitalSign.title : 'No title' }}
        <v-btn
          small
          class="float-end mr-5"
          color="error"
          @click="$emit('remove')"
        >
          delete
        </v-btn>
      </div>
    </v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-select
        v-model="vitalSign.title"
        :items="titleOptions()"
        label="Title"
      />
      <Form :vital-signs.sync="vitalSign.data" />
      <v-btn class="mb-4" @click="addChild"> Add Step </v-btn>
      <v-expansion-panels multiple class="mb-4">
        <VitalSign
          v-for="(vs, i) in vitalSign.childs"
          :key="i"
          :vital-sign.sync="vitalSign.childs[i]"
          :level="level + 1"
          :isPrehospital="isPrehospital"
          @remove="vitalSign.childs.splice(i, 1)"
        />
      </v-expansion-panels>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
import Form from '@/components/vital_signs/form.vue'
@Component({
  name: 'VitalSign',
  components: { Form },
})
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSign!: any
  @Prop({ type: Number, required: true }) readonly level!: number
  @Prop({ type: Boolean, required: true }) readonly isPrehospital!: Boolean
  @Watch('vitalSign', { immediate: true, deep: true })
  updateVitalSignChanged(val: any) {
    this.$emit('update:vitalSign', val)
  }

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
        treatmentExpected: undefined,
      },
    },
    childs: [],
  }

  addChild() {
    if (this.vitalSign.childs && this.vitalSign.childs.length >= 2) return
    this.vitalSign.childs.push(JSON.parse(JSON.stringify(this.emptyVitalSign)))
  }

  titleOptions() {
    if (this.level === 0) return ['T0']

    if (this.isPrehospital && this.level === 1) {
      return ['A Pre', 'B Pre']
    }

    if (this.isPrehospital) return ['T1 A', 'T2 A', 'T1 B', 'T2 B']
    return ['T1', 'T2']
  }

  levelColor() {
    const colors = ['primary', 'success', 'warning', 'error']

    return colors[this.level % 4]
  }
}
</script>
