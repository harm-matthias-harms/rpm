<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="8" md="10" sm="12">
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ medicalCase.title }}
              <v-icon color="primary" @click="editMedicalCase(medicalCase)">
                edit
              </v-icon>
              <DeleteButton
                v-if="medicalCase.author.id == $store.state.user.user.id"
                :item="medicalCase"
                :go-back="true"
              />
            </h4>
            <v-row>
              <v-col v-if="medicalCase.author && medicalCase.author.username" class="col-auto">
                <Author :author="medicalCase.author" :created-at="medicalCase.createdAt" />
              </v-col>
              <v-col
                v-if="medicalCase.editor && medicalCase.editor.username"
                class="col-auto mr-auto"
              >
                <Editor :editor="medicalCase.editor" :updated-at="medicalCase.editedAt" />
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-chip v-for="(tag,i) in tags(medicalCase)" :key="i" class="ma-1">
                  {{ tag }}
                </v-chip>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <p class="body-1 py-3 black--text">
                  {{ medicalCase.generalInformation.shortSummary }}
                </p>
              </v-col>
            </v-row>
          </v-card-text>
          <v-expansion-panels v-model="expansionPanel" multiple>
            <v-expansion-panel>
              <v-expansion-panel-header>Medical history</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <div v-if="medicalCase.medicalHistory.problems">
                    <v-list-item>
                      <v-list-item-content>Problems/conditions:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.medicalHistory.problems }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <div v-if="medicalCase.medicalHistory.vaccinations">
                    <v-list-item>
                      <v-list-item-content>Vaccinations:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.medicalHistory.vaccinations }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <div v-if="medicalCase.medicalHistory.allergies">
                    <v-list-item>
                      <v-list-item-content>Allergies:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.medicalHistory.allergies }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <div v-if="medicalCase.medicalHistory.medication">
                    <v-list-item>
                      <v-list-item-content>Medication:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.medicalHistory.medication }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <v-list-item v-if="medicalCase.medicalHistory.implantedDevices">
                    <v-list-item-content>Implantable devices:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.medicalHistory.implantedDevices }}
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
              <v-expansion-panel-header>Expectations</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <div v-if="medicalCase.expectations.generalStatus">
                    <v-list-item>
                      <v-list-item-content>General status:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.expectations.generalStatus }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <div v-if="medicalCase.expectations.onExamination">
                    <v-list-item>
                      <v-list-item-content>On examination:</v-list-item-content>
                      <v-list-item-content
                        class="align-end"
                      >
                        {{ medicalCase.expectations.onExamination }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <v-list-item v-if="medicalCase.expectations.expectations">
                    <v-list-item-content>Expectations:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.expectations.expectations }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
          <v-list dense>
            <v-list-item v-if="medicalCase.otherInformation">
              <v-list-item-content>Other information:</v-list-item-content>
              <v-list-item-content class="align-end">
                {{ medicalCase.otherInformation }}
              </v-list-item-content>
            </v-list-item>
            <v-divider v-if="medicalCase.makeup && medicalCase.otherInformation" />
            <v-list-item v-if="medicalCase.makeup">
              <v-list-item-content>Needed make-up and attributes:</v-list-item-content>
              <v-list-item-content class="align-end">
                {{ medicalCase.makeup }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-row class="px-4">
            <v-col v-for="(file, i) in medicalCase.files" :key="i" sm="12" md="4" lg="3">
              <v-card>
                <v-icon>attach_file</v-icon>
                <v-card-title class="body-2">
                  <a
                    :href="($axios.defaults.baseURL === '/' ? '' : $axios.defaults.baseURL) + '/api/medical_cases/' + medicalCase.id + '/documents/' + file.id"
                  >
                    {{ file.name }}
                    <br>
                    ({{ bytesToSize(file.size) }})
                  </a>
                </v-card-title>
                <v-card-actions v-if="medicalCase.author.id == $store.state.user.user.id">
                  <v-btn text color="red" @click="$set(deleteFileDialog, i, true)">
                    Delete
                  </v-btn>
                  <Confirm
                    text="Are you sure you want to delete this file?"
                    :dialog.sync="deleteFileDialog[i]"
                    :item="{mcID: medicalCase.id, id: file.id}"
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
    Confirm
  },
  computed: {
    ...mapState('medicalCase', {
      medicalCase: 'medicalCase'
    })
  },
  methods: {
    ...mapActions('medicalCase', {
      find: 'find',
      deleteFile: 'deleteFile'
    })
  }
})
export default class ShowMedicalCase extends Vue {
  find!: (id) => void
  deleteFile!: ({ mcID, id }) => void
  deleteFileDialog: any = new Array(100).fill(false)
  medicalCase!: any
  expansionPanel: Array<number> = []

  mounted () {
    const id = this.$route.params.id
    if (this.medicalCase.id !== id) {
      this.find(id)
    }
    this.setExpansionPanel()
  }

  editMedicalCase (medicalCase) {
    this.$router.push('/medical_cases/' + medicalCase.id + '/edit')
  }

  onCancel () {
    this.deleteFileDialog = new Array(100).fill(false)
  }

  setExpansionPanel () {
    if (this.medicalCase) {
      if (this.anyFieldPresent(this.medicalCase.generalInformation)) {
        this.expansionPanel.push(0)
      }
      if (this.anyFieldPresent(this.medicalCase.medicalHistory)) {
        this.expansionPanel.push(1)
      }
      if (
        this.medicalCase.vitalSigns &&
        this.medicalCase.vitalSigns.length > 0 &&
        this.anyFieldPresent(this.medicalCase.vitalSigns[0])
      ) {
        this.expansionPanel.push(2)
      }
      if (this.anyFieldPresent(this.medicalCase.expectations)) {
        this.expansionPanel.push(3)
      }
    }
  }

  anyFieldPresent (object: Object) {
    const exist = (elem) => {
      if (elem) {
        return true
      }
    }
    return Object.values(object).some(exist)
  }

  tags (medicalCase) {
    const tags: string[] = []
    if (medicalCase.generalInformation.usar) { tags.push('USAR') }
    if (medicalCase.generalInformation.medivac) { tags.push('MEDIVAC') }
    if (medicalCase.generalInformation.hospilisation) { tags.push('Need for hospilisation') }
    if (medicalCase.generalInformation.surgical) { tags.push('Surgical') }
    if (medicalCase.generalInformation.triage) { tags.push('Triage: ' + medicalCase.generalInformation.triage) }
    if (medicalCase.generalInformation.age) { tags.push('Age: ' + medicalCase.generalInformation.age) }
    if (medicalCase.generalInformation.gender) { tags.push('Gender: ' + medicalCase.generalInformation.gender) }
    return tags
  }

  bytesToSize (bytes) {
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
