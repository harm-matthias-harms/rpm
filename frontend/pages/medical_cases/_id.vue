<template>
  <v-container>
    <v-row justify="center">
      <v-col
        lg="4"
        md="8"
        sm="12"
      >
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ medicalCase.title }}
            </h4>
            <v-row justify="center">
              <v-col class="col-auto">
                <Author
                  :author="medicalCase.author"
                  :created-at="medicalCase.createdAt"
                />
              </v-col>
              <v-col class="col-auto mr-auto">
                <Editor
                  v-if="medicalCase.editor.username"
                  :editor="medicalCase.editor"
                  :updated-at="medicalCase.updatedAt"
                />
              </v-col>
            </v-row>
          </v-card-text>
          <v-divider />
          <!-- <VitalSigns
            class="black--text"
            :vital-signs="medicalCase.vitalSigns"
          /> -->
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
  import VitalSigns from '@/components/preset/vitalSigns.vue'
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

  mounted () {
    const id = this.$route.params.id
    if (this.medicalCase.id !== id) {
      this.find(id)
    }
  }
  }
</script>
