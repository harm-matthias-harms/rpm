<template>
  <div>
    <v-autocomplete
      v-model="team"
      :items="options"
      item-text="title"
      return-object
    >
    </v-autocomplete>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import { Team } from '~/store/team/type'

@Component({
  methods: {
    ...mapActions('team', {
      getTeams: 'get_all',
    }),
  },
})
export default class TeamSelector extends Vue {
  @Prop({ type: Object }) readonly value!: Team | undefined

  getTeams!: () => Promise<void>

  get team(): Team | undefined {
    return this.value
  }

  set team(team: Team | undefined) {
    this.$emit('input', team)
  }

  options: Team[] = []

  mounted() {
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