<template>
  <div>
    <div v-for="(medicalCase, index) in medicalCases" :key="medicalCase.id">
      <v-card class="elevation-0">
        <v-card-title>
          {{ medicalCase.title }}
        </v-card-title>
        <v-card-text>
          <RoleplayerForm
            v-model="roleplayer[index]"
            :medicalCase="medicalCase"
          />
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Watch, Component, Vue } from 'vue-property-decorator'
import RoleplayerForm from './RoleplayerForm.vue'
import { Roleplayer } from '~/store/inject/type'
import { MedicalCaseShort } from '~/store/medicalCase/type'

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
    this.$emit('input', roleplayer)
  }

  @Watch('medicalCases', { deep: true })
  onMedicalCasesChanged() {
    while (this.roleplayer.length < this.medicalCases.length) {
      this.roleplayer.push({})
    }
  }
}
</script>
