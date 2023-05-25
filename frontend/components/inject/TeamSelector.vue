<template>
  <div>
    <v-autocomplete
      v-model="team"
      :items="filteredTeams"
      item-text="title"
      return-object
    />
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import { Team } from '~/store/team/type'

@Component({
  methods: {
    ...mapActions('team', {
      getTeams: 'get_all'
    })
  }
})
export default class TeamSelector extends Vue {
  @Prop({ type: Object }) readonly value!: Team | undefined

  getTeams!: () => Promise<void>

  get team (): Team | undefined {
    return this.value
  }

  set team (team: Team | undefined) {
    this.$emit('input', team)
  }

  options: Team[] = []

  get filteredTeams(): Team[] {
    return this.options.filter((team) => {
      for(const exerciseTeam of this.$store.state.exercise.exercise.teams) {
        if (exerciseTeam.team.id == team.id) return true
      };

      return false
    })
  }

  mounted () {
    if (this.$store.state.team.teamsLoaded) {
      this.options = this.$store.state.team.teamList.teams
    } else {
      this.getTeams().then(() => {
        this.options = this.$store.state.team.teamList.teams
      })
    }
  }
}
</script>
