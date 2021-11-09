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
          v-if="isAdminOrRpm"
          small
          color="primary"
          class="mr-2"
          :to="`/exercises/${$route.params.id}/injects/new`"
        >
          <v-icon small> fas fa-plus </v-icon>
        </v-btn>
        <v-btn
          small
          color="primary"
          @click="getInjects({ exerciseID: $route.params.id })"
        >
          <v-icon small> fas fa-redo </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col xl="10" cols="12">
        <Table
          :items="injectList.injects"
          :exerciseID="this.$route.params.id"
          :loading="loading"
          :refresh-table="getInjects"
        />
      </v-col>
    </v-row>
    <v-row justify="center" class="mt-16">
      <v-col md="8" cols="12">
        <Statistics :items="injectList.injects" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Table from '~/components/inject/Table.vue'
import Statistics from '~/components/inject/Statistics.vue'
@Component({
  components: {
    Table,
    Statistics,
  },
  computed: {
    ...mapState('inject', {
      injectList: 'injectList',
      injectsLoaded: 'injectsLoaded',
      exerciseID: 'exerciseID',
    }),
  },
  methods: {
    ...mapActions('inject', {
      getInjects: 'get_all',
    }),
  },
})
export default class Injects extends Vue {
  getInjects!: (payload) => void
  injectsLoaded!: boolean
  injectList!: object
  exerciseID!: string

  loading = false

  loadInjects() {
    this.loading = true
    this.getInjects({ exerciseID: this.$route.params.id })
    this.loading = false
  }

  mounted() {
    if (!this.injectsLoaded || this.exerciseID !== this.$route.params.id) {
      this.loadInjects()
    }
  }

  get exerciseRoles() {
    return this.$store.state.user.user.roles.filter(
      (role) => role.exercise.id === this.exerciseID
    )
  }

  get isAdminOrRpm() {
    return this.exerciseRoles.some((role) =>
      ['admin', 'role play manager'].includes(role.role)
    )
  }
}
</script>
