<template>
  <v-form @submit.prevent="submit()">
    <v-stepper v-model="stepper">
      <v-stepper-header>
        <v-stepper-step :complete="stepper > 1" step="1" editable>
          Select a time
        </v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step :complete="stepper > 2" step="2" editable>
          Select a team
        </v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step :complete="stepper > 3" step="3" editable>
          Select a make-up center
        </v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step :complete="stepper > 4" step="4" editable>
          Select medical cases
        </v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="5" editable> Select roleplayer </v-stepper-step>
      </v-stepper-header>
      <v-stepper-items>
        <v-stepper-content step="1">
          <StartTime v-model="startTime"></StartTime>
          <v-btn color="primary" @click="stepper = 2" :disabled="!startTime">
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content step="2">
          <TeamSelector v-model="team" />
          <v-btn color="primary" @click="stepper = 3" :disabled="!team">
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content step="3">
          <MakeupCenterSelector v-model="makeupCenter" />
          <v-btn color="primary" @click="stepper = 4" :disabled="!makeupCenter">
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content step="4">
          <MedicalCaseSelector v-model="medicalCases"></MedicalCaseSelector>
          <v-btn
            color="primary"
            @click="stepper = 5"
            :disabled="medicalCases.length == 0"
          >
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content step="5">
          <RoleplayerCreator
            :medicalCases="medicalCases"
            v-model="roleplayer"
          ></RoleplayerCreator>
          <v-btn
            color="primary"
            type="submit"
            :disabled="roleplayer.some((rp) => !Boolean(rp.fullName))"
          >
            Create
          </v-btn>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </v-form>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import { Inject, Roleplayer } from '~/store/inject/type'
import { Team } from '~/store/team/type'
import { MedicalCaseShort } from '~/store/medicalCase/type'
import StartTime from './StartTime.vue'
import TeamSelector from './TeamSelector.vue'
import MakeupCenterSelector from './MakeupCenterSelector.vue'
import MedicalCaseSelector from './MedicalCaseSelector.vue'
import RoleplayerCreator from './RoleplayerCreator.vue'

import _ from 'lodash'

@Component({
  components: {
    StartTime,
    TeamSelector,
    MakeupCenterSelector,
    MedicalCaseSelector,
    RoleplayerCreator,
  },
  methods: {
    ...mapActions('medicalCase', {
      findMedicalCase: 'find',
    }),
  },
})
export default class Form extends Vue {
  @Prop({ type: Function, required: true }) readonly atSubmit!: (
    payload
  ) => void

  findMedicalCase!: (id) => void

  stepper: number = 1
  startTime: Date | null = null
  team: Team | null = null
  makeupCenter: string | null = null
  medicalCases: MedicalCaseShort[] = []
  roleplayer: Roleplayer[] = []

  async submit() {
    const injects = await this.createInjects()

    this.atSubmit({ injects: injects, exerciseID: this.$route.params.id })
  }

  async createInjects(): Promise<Inject[]> {
    const injects: Inject[] = []
    for (const [index, medicalCase] of this.medicalCases.entries()) {
      await this.findMedicalCase(medicalCase.id)
      injects.push({
        exerciseID: this.$route.params.id,
        status: 'Waiting for Makeup',
        startTime: this.startTime,
        makeupCenterTitle: this.makeupCenter,
        team: this.team,
        medicalCase: this.$store.state.medicalCase.medicalCase,
        roleplayer: this.roleplayer[index],
        author: null,
        editor: null,
      })
    }
    return injects
  }
}
</script>
