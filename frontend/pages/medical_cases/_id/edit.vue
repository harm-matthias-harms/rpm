<template>
  <v-container>
    <v-row justify="center">
      <v-col
        lg="8"
        md="10"
        sm="12"
      >
        <v-card>
          <v-card-title primary-title>
            Edit Medical Case
          </v-card-title>
          <v-card-text>
            <Form
              :medical-case="medicalCase"
              :at-submit="update"
              :is-new="false"
            />
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import Form from '@/components/medical_case/form.vue'

@Component({
  components: {
    Form
  },
  methods: {
    ...mapActions('medicalCase', {
      find: 'find',
      update: 'update'
    })
  }
})
export default class ShowMedicalCase extends Vue {
  find!: (id) => Promise<any>
  create!: (medicalCase) => void
  medicalCase: any = JSON.parse(JSON.stringify(this.$store.state.medicalCase.medicalCase))
  expansionPanel: Array<number> = []

  mounted () {
    const id = this.$route.params.id
    if (this.medicalCase.id !== id) {
      this.find(id).then(() => {
        this.medicalCase = JSON.parse(JSON.stringify(this.$store.state.medicalCase.medicalCase))
      })
    }
  }
}
</script>
