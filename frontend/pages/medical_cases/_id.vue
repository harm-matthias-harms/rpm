<template>
  <v-container>
    <v-row justify="center">
      <v-col
        lg="8"
        md="10"
        sm="12"
      >
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ medicalCase.title }}
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
                  :updated-at="medicalCase.editedAt"
                />
              </v-col>
            </v-row>
          </v-card-text>
          <v-expansion-panels
            v-model="expansionPanel"
            multiple
          >
            <v-expansion-panel>
              <v-expansion-panel-header>General Information</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-chip v-if="medicalCase.generalInformation.usar">
                  USAR
                </v-chip>
                <v-chip v-if="medicalCase.generalInformation.medivac">
                  MEDIVAC
                </v-chip>
                <v-chip v-if="medicalCase.generalInformation.surgical">
                  Surgical
                </v-chip>
                <v-chip v-if="medicalCase.generalInformation.hospilisation">
                  Need for Hospilisation
                </v-chip>
                <v-list dense>
                  <v-list-item v-if="medicalCase.generalInformation.shortSummary">
                    <v-list-item-content>Short Summary:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.generalInformation.shortSummary }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.generalInformation.triage">
                    <v-list-item-content>Triage:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.generalInformation.triage }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.generalInformation.age">
                    <v-list-item-content>Age:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.generalInformation.age }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.generalInformation.gender">
                    <v-list-item-content>Gender:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.generalInformation.gender }}
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>Medical History</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.medicalHistory.problems ">
                    <v-list-item-content>Problems/Conditions:</v-list-item-content>
                    <v-list-item-content class="align-end">
                      {{ medicalCase.medicalHistory.problems }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medicalHistory.vaccinations">
                    <v-list-item-content>Vaccinations:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.medicalHistory.vaccinations }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medicalHistory.allergies">
                    <v-list-item-content>Allergies:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.medicalHistory.allergies }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medicalHistory.medication">
                    <v-list-item-content>Medication:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.medicalHistory.medication }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.medicalHistory.implantedDevices">
                    <v-list-item-content>Implantable Devices:</v-list-item-content>
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
              <v-expansion-panel-header>Vital Signs</v-expansion-panel-header>
              <v-expansion-panel-content>
                <VitalSigns :vital-signs="medicalCase.vitalSigns" />
              </v-expansion-panel-content>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-header>Expectations</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list dense>
                  <v-list-item v-if="medicalCase.expectations.generalStatus">
                    <v-list-item-content>General Status:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.expectations.generalStatus }}
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item v-if="medicalCase.expectations.onExamination">
                    <v-list-item-content>On Examination:</v-list-item-content>
                    <v-list-item-content
                      class="align-end"
                    >
                      {{ medicalCase.expectations.onExamination }}
                    </v-list-item-content>
                  </v-list-item>
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
              <v-list-item-content>Other Information:</v-list-item-content>
              <v-list-item-content class="align-end">
                {{ medicalCase.otherInformation }}
              </v-list-item-content>
            </v-list-item>
            <v-list-item v-if="medicalCase.makeup">
              <v-list-item-content>Needed Make-Up and Attributes:</v-list-item-content>
              <v-list-item-content class="align-end">
                {{ medicalCase.makeup }}
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-row class="px-4">
            <v-col
              v-for="(file, i) in medicalCase.files"
              :key="i"
              sm="12"
              md="4"
              lg="3"
            >
              <v-card
                :href="$axios.defaults.baseURL + '/api/medical_cases/' + medicalCase.id + '/documents/' + file.id"
              >
                <v-icon>attach_file</v-icon>
                <v-card-title class="body-2">
                  {{ file.name }}
                  <br>
                  ({{ bytesToSize(file.size) }})
                </v-card-title>
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

@Component({
  components: {
    Author,
    Editor,
    VitalSigns,
  },
  computed: {
    ...mapState('medicalCase', {
      medicalCase: 'medicalCase',
    }),
  },
  methods: {
    ...mapActions('medicalCase', {
      find: 'find',
    }),
  },
})
  export default class ShowMedicalCase extends Vue {
  find!: (id) => void
  medicalCase!: any
  expansionPanel: Array<number> = []

  mounted () {
    const id = this.$route.params.id
    if (this.medicalCase.id !== id) {
      this.find(id)
    }
    this.setExpansionPanel()
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

  bytesToSize (bytes) {
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
    if (bytes === 0) return 'n/a'
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    if (i === 0) return bytes + ' ' + sizes[i]
    return (bytes / 1024 ** i).toFixed(2) + ' ' + sizes[i]
  }
  }
</script>
