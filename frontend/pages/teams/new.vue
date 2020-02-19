<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
      <v-card>
        <v-card-title primary-title>
          New Team
        </v-card-title>
        <v-card-text>
          <TeamForm :team="team" :at-submit="create" :is-new="true" />
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import TeamForm from '@/components/team/form.vue'

@Component({
  components: {
    TeamForm
  },
  methods: {
    ...mapActions('team', {
      create: 'create'
    })
  }
})
export default class NewTeam extends Vue {
  team = {}
  create!: (team) => void
}
</script>
