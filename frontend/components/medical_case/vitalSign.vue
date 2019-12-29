<template>
  <v-expansion-panel>
    <v-expansion-panel-header>{{vitalSign.title ? vitalSign.title : 'No title set'}}</v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-text-field v-model="vitalSign.title" label="Title"></v-text-field>
      <v-text-field v-model="vitalSign.reason" label="Reason"></v-text-field>
      <Form :vitalSigns.sync="vitalSign.data" />
      <v-btn @click="addChild" class="mb-4">Add Step</v-btn>
      <v-expansion-panels multiple class="mb-4">
        <VitalSign v-for="(vs, i) in vitalSign.childs" :key="i" :vitalSign.sync="vs" />
      </v-expansion-panels>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
import Form from '@/components/vital_signs/form.vue'
@Component({
  components: {
    Form
  }
})
export default class VitalSign extends Vue {
  @Prop({ type: Object, required: true }) readonly vitalSign!: any
  @Watch('vitalSign', { immediate: true, deep: true })
  updateVitalSignChanged(val: any, oldVal: any) {
    this.$emit('update:vitalSign', val)
  }

  emptyVitalSign: object = {
    title: undefined,
    reason: undefined,
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
      weight: undefined,
      height: undefined
    },
    childs: []
  }

  addChild() {
    this.vitalSign.childs.push(this.emptyVitalSign)
  }
}
</script>