<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
      <v-card>
        <v-card-title primary-title>
          New Case
        </v-card-title>
        <v-card-text>
          <Form :at-submit="create" />
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import Form from '@/components/inject/Form.vue'

@Component({
  components: {
    Form
  },
  methods: {
    ...mapActions('inject', {
      create: 'create'
    })
  }
})
export default class NewInject extends Vue {
  create!: (payload) => void
}
</script>
