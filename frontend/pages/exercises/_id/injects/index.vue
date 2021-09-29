<template>
  <v-container>
    <v-row justify="center">
      <v-col md="5" sm="6">
        <h3 class="headline">
          <v-chip small label color="primary">
            {{ injectList.count }}
          </v-chip>
          Cases
        </h3>
      </v-col>
      <v-col md="5" sm="6" class="justify-end d-flex">
        <v-btn
          small
          color="primary"
          class="mr-2"
          :to="`/exercises/${$route.params.id}/injects/new`"
        >
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
        <v-btn
          small
          color="primary"
          @click="getInjects({ exerciseID: $route.params.id })"
        >
          <v-icon small>
            fas fa-redo
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="10" sm="12">
        <Table :items="injectList.injects" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Table from '~/components/inject/Table.vue'
@Component({
  components: {
    Table,
  },
  computed: {
    ...mapState('inject', {
      injectList: 'injectList',
      injectsLoaded: 'injectsLoaded',
    }),
  },
  methods: {
    ...mapActions('inject', {
      getInjects: 'get_all',
    }),
  },
})
export default class MedicalCases extends Vue {
  getInjects!: (payload) => void
  injectsLoaded!: boolean
  injectList!: object

  loading = false

  loadInjects() {
    this.loading = true
    this.getInjects({ exerciseID: this.$route.params.id })
    this.loading = false
  }

  mounted() {
    if (!this.injectsLoaded) {
      this.loadInjects()
    }
  }
}
</script>
