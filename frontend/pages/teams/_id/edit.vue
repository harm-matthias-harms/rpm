<template>
  <v-row justify="center">
    <v-col lg="6" md="10" sm="12">
      <v-card>
        <v-card-title primary-title>
          Edit team
        </v-card-title>
        <v-card-text>
          <TeamForm :team="team" :at-submit="update" :is-new="false" />
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
    TeamForm,
  },
  methods: {
    ...mapActions('team', {
      find: 'find',
      update: 'update',
    }),
  },
})
export default class Edit extends Vue {
  find!: (id) => Promise<any>
  update!: (team) => void
  team: any = JSON.parse(JSON.stringify(this.$store.state.team.team))

  mounted() {
    const id = this.$route.params.id
    if (this.team.id !== id) {
      this.find({ id }).then(() => {
        this.team = JSON.parse(JSON.stringify(this.$store.state.team.team))
      })
    }
  }
}
</script>
