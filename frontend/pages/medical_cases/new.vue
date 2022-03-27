<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
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
import { mapActions, mapState } from 'vuex'
import Form from '@/components/medical_case/form.vue'
import { MedicalCase } from '~/store/medicalCase/type'
import { defaultMedicalCase } from '~/store/medicalCase/state'

@Component({
  components: {
    Form,
  },
  computed: {
    ...mapState('medicalCase', {
      clone: 'medicalCase',
    }),
  },
  methods: {
    ...mapActions('medicalCase', {
      create: 'create',
      find: 'find',
    }),
  },
})
export default class NewMedicalCase extends Vue {
  medicalCase: MedicalCase = Object.assign({}, defaultMedicalCase)
  create!: (medicalCase, files) => void
  find!: (id) => void

  async mounted() {
    const cloneId = this.$route.query.cloneId
    if (cloneId) {
      let clone: MedicalCase = await this.$axios.$get(
        `/api/medical_cases/${cloneId}`
      )
      clone = {
        ...clone,
        id: undefined,
        author: {},
        editor: {},
        createdAt: undefined,
        editedAt: undefined,
        files: [],
      }
      this.medicalCase = Object.assign({}, clone)
    }
  }
}
</script>
