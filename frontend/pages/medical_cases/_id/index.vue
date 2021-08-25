<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="8" md="10" sm="12">
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ medicalCase.title }}
              <v-icon
                v-if="!medicalCase.approved"
                color="green"
                @click="approve(medicalCase)"
              >
                check
              </v-icon>
              <v-icon color="primary" @click="editMedicalCase(medicalCase)">
                edit
              </v-icon>
              <DeleteButton
                v-if="
                  !medicalCase.author.username ||
                  medicalCase.author.id == $store.state.user.user.id
                "
                :item="medicalCase"
                :go-back="true"
              />
            </h4>
            <v-row>
              <v-col
                v-if="medicalCase.author && medicalCase.author.username"
                class="col-auto"
              >
                <Author
                  :author="medicalCase.author"
                  :created-at="medicalCase.createdAt"
                />
              </v-col>
              <v-col
                v-if="medicalCase.editor && medicalCase.editor.username"
                class="col-auto mr-auto"
              >
                <Editor
                  :editor="medicalCase.editor"
                  :edited-at="medicalCase.editedAt"
                />
              </v-col>
            </v-row>
          </v-card-text>
          <v-expansion-panels v-model="expansionPanel" multiple>
            <v-expansion-panel>
              <v-expansion-panel-header>General</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.general.discipline">
                    <v-list-item-content>Area/Discipline:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.general.discipline.join(', ') }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.general.context">
                    <v-list-item-content>Context:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.general.context.join(', ') }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.general.scenario">
                    <v-list-item-content>Scenario:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.general.scenario.join(', ') }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>
                Patient's characteristics
              </v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.patient.type">
                    <v-list-item-content>Patient type:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.patient.type }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.patient.triage">
                    <v-list-item-content>Triage:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.patient.triage }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.patient.gender">
                    <v-list-item-content>Gender:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.patient.gender.join(', ') }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.patient.age">
                    <v-list-item-content>Age:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.patient.age }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>
                Medical assessment
              </v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.medical.signs">
                    <v-list-item-content>
                      S - Signs/Symptoms:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.signs }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medical.allergies">
                    <v-list-item-content>A - Allergies:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.allergies }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medical.medication">
                    <v-list-item-content>M - Medications:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.medication }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medical.past">
                    <v-list-item-content>
                      P - Past pertinent medical history:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.past }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medical.loi">
                    <v-list-item-content>
                      L - Last oral intake:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.loi }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medical.events">
                    <v-list-item-content>
                      E - Events leading up to present illness/injury:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medical.events }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>Vital signs</v-expansion-panel-header>
              <v-expansion-panel-content>
                <VitalSigns :vital-signs="medicalCase.vitalSigns" />
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>
                Make-up and attributes
              </v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.makeup.makeup">
                    <v-list-item-content>
                      Make-up instructions:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.makeup.makeup }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.makeup.acting">
                    <v-list-item-content>
                      Acting instructions:
                    </v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.makeup.acting }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
          <v-row class="px-4">
            <v-col
              v-for="(file, i) in medicalCase.files"
              :key="i"
              sm="12"
              md="4"
              lg="3"
            >
              <v-card>
                <v-icon>attach_file</v-icon>
                <v-card-title class="body-2">
                  <a
                    :href="
                      ($axios.defaults.baseURL === '/'
                        ? ''
                        : $axios.defaults.baseURL) +
                      '/api/medical_cases/' +
                      medicalCase.id +
                      '/documents/' +
                      file.id
                    "
                  >
                    {{ file.name }}
                    <br />
                    ({{ bytesToSize(file.size) }})
                  </a>
                </v-card-title>
                <v-card-actions
                  v-if="
                    !medicalCase.author.username ||
                    medicalCase.author.id == $store.state.user.user.id
                  "
                >
                  <v-btn
                    text
                    color="red"
                    @click="$set(deleteFileDialog, i, true)"
                  >
                    Delete
                  </v-btn>
                  <Confirm
                    text="Are you sure you want to delete this file?"
                    :dialog.sync="deleteFileDialog[i]"
                    :item="{ mcID: medicalCase.id, id: file.id }"
                    :at-submit="deleteFile"
                    :at-cancel="onCancel"
                  />
                </v-card-actions>
              </v-card>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Author from '@/components/utils/Author.vue'
import Editor from '@/components/utils/Editor.vue'
import VitalSigns from '@/components/medical_case/vital_signs/show.vue'
import DeleteButton from '@/components/medical_case/Delete.vue'
import Confirm from '@/components/utils/Confirm.vue'

@Component({
  components: {
    Author,
    Editor,
    VitalSigns,
    DeleteButton,
    Confirm,
  },
  computed: {
    ...mapState('medicalCase', {
      medicalCase: 'medicalCase',
    }),
  },
  methods: {
    ...mapActions('medicalCase', {
      find: 'find',
      approve: 'approve',
      deleteFile: 'deleteFile',
    }),
  },
})
export default class ShowMedicalCase extends Vue {
  find!: (id) => void
  deleteFile!: ({ mcID, id }) => void
  deleteFileDialog: any = new Array(100).fill(false)
  medicalCase!: any
  expansionPanel: Array<number> = []

  mounted() {
    const id = this.$route.params.id
    if (this.medicalCase.id !== id) {
      this.find(id)
    }
    this.setExpansionPanel()
  }

  editMedicalCase(medicalCase) {
    this.$router.push('/medical_cases/' + medicalCase.id + '/edit')
  }

  onCancel() {
    this.deleteFileDialog = new Array(100).fill(false)
  }

  setExpansionPanel() {
    if (this.medicalCase) {
      if (this.anyFieldPresent(this.medicalCase.general)) {
        this.expansionPanel.push(0)
      }
      if (this.anyFieldPresent(this.medicalCase.patient)) {
        this.expansionPanel.push(1)
      }
      if (this.anyFieldPresent(this.medicalCase.medical)) {
        this.expansionPanel.push(2)
      }
      if (
        this.medicalCase.vitalSigns &&
        this.medicalCase.vitalSigns.length > 0 &&
        this.anyFieldPresent(this.medicalCase.vitalSigns[0])
      ) {
        this.expansionPanel.push(3)
      }
      if (this.anyFieldPresent(this.medicalCase.makeup)) {
        this.expansionPanel.push(4)
      }
    }
  }

  anyFieldPresent(object: Object) {
    const exist = (elem) => {
      if (elem) {
        return true
      }
    }
    return Object.values(object).some(exist)
  }

  bytesToSize(bytes) {
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
    if (bytes === 0) {
      return 'n/a'
    }
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    if (i === 0) {
      return bytes + ' ' + sizes[i]
    }
    return (bytes / 1024 ** i).toFixed(2) + ' ' + sizes[i]
  }
}
</script>

<style scoped>
a {
  text-decoration: none;
}
p,
.v-list-item__content {
  white-space: pre-line;
}
</style>
