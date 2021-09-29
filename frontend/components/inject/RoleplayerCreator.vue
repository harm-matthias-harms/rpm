<template>
  <div>
    <div v-for="(medicalCase, index) in medicalCases" :key="medicalCase.id">
      <v-card class="elevation-0">
        <v-card-title>
          {{ medicalCase.title }}
        </v-card-title>
        <v-card-text>
          <RoleplayerForm v-model="roleplayer[index]"></RoleplayerForm>
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
import { Roleplayer } from '~/store/inject/type'
import { MedicalCaseShort } from '~/store/medicalCase/type'
import RoleplayerForm from './RoleplayerForm.vue'

@Component({
  components: {
    RoleplayerForm,
  },
})
export default class RoleplayerCreator extends Vue {
  @Prop({ type: Array }) readonly value!: Roleplayer[]
  @Prop({ type: Array }) readonly medicalCases!: MedicalCaseShort[]

  get roleplayer(): Roleplayer[] {
    return this.value
  }

  set roleplayer(roleplayer: Roleplayer[]) {
    console.log(roleplayer)
    this.$emit('input', roleplayer)
  }

  @Watch('medicalCases', { deep: true })
  onMedicalCasesChanged(val: MedicalCaseShort, oldVal: MedicalCaseShort) {
    while (this.roleplayer.length < this.medicalCases.length) {
      this.roleplayer.push({})
    }
  }
}
</script>