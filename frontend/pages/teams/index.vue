<template>
  <v-container>
    <v-row justify="center">
      <v-col md="5" sm="6">
        <h3 class="headline">
          <v-chip small label color="primary">{{ teamList.count }}</v-chip>Teams
        </h3>
      </v-col>
      <v-col md="5" sm="6" class="justify-end d-flex">
        <v-btn small color="primary" class="mr-2" to="/teams/new">
          <v-icon small>fas fa-plus</v-icon>
        </v-btn>
        <v-btn small color="primary" @click="getTeams()">
          <v-icon small>fas fa-redo</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="10" sm="12">
        <TeamTable :items="teamList.teams" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import TeamTable from '@/components/team/table.vue'
@Component({
  components: {
    TeamTable
  },
  computed: {
    ...mapState('team', {
      teamList: 'teamList',
      teamsLoaded: 'teamsLoaded'
    })
  },
  methods: {
    ...mapActions('team', {
      getTeams: 'get_all'
    })
  }
})
export default class Teams extends Vue {
  getTeams!: () => void
  teamsLoaded!: boolean

  loading: boolean = false

  loadTeams() {
    this.loading = true
    this.getTeams()
    this.loading = false
  }

  mounted() {
    if (!this.teamsLoaded) {
      this.loadTeams()
    }
  }
}
</script>
