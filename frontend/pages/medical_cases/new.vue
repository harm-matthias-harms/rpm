<template>
  <v-row justify="center">
    <v-col
      lg="6"
      md="10"
      sm="12"
    >
      <v-card>
        <v-card-title primary-title>
          New Medical Case
        </v-card-title>
        <v-card-text>
          <Form
            :medical-case="medicalCase"
            :at-submit="create"
            :is-new="true"
          />
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
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
      create: 'create'
    })
  }
})
export default class NewMedicalCase extends Vue {
  medicalCase = {
    general: {},
    patient: {},
    medical: {},
    makeup: {},
    vitalSigns: [],
    files: []
  }

  create!: (medicalCase, files) => void
}
</script>
