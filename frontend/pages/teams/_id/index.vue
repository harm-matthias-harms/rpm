<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="8" md="10" sm="12">
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ team.title }}
              <v-icon color="primary" @click="editTeam(team)">
                edit
              </v-icon>
              <DeleteButton
                v-if="team.author.id == $store.state.user.user.id"
                :item="team"
                :go-back="true"
              />
            </h4>

            <v-row>
              <v-col
                v-if="team.author && team.author.username"
                class="col-auto"
              >
                <Author :author="team.author" :created-at="team.createdAt" />
              </v-col>
              <v-col
                v-if="team.editor && team.editor.username"
                class="col-auto mr-auto"
              >
                <Editor :editor="team.editor" :edited-at="team.editedAt" />
              </v-col>
            </v-row>
          </v-card-text>
          <v-list dense class="body-1">
            <v-list-item>
              <v-list-item-content>Type:</v-list-item-content>
              <v-list-item-content>{{ team.type }}</v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>Medivac:</v-list-item-content>
              <v-list-item-content>
                <v-icon v-if="team.medivac">
                  done
                </v-icon>
                <v-icon v-if="!team.medivac">
                  clear
                </v-icon>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Author from '@/components/utils/Author.vue'
import Editor from '@/components/utils/Editor.vue'
import DeleteButton from '@/components/team/Delete.vue'
@Component({
  components: {
    Author,
    Editor,
    DeleteButton,
  },
  computed: {
    ...mapState('team', {
      team: 'team',
    }),
  },
  methods: {
    ...mapActions('team', {
      find: 'find',
    }),
  },
})
export default class ShowTeam extends Vue {
  find!: (id) => void
  team!: any

  mounted() {
    const id = this.$route.params.id
    if (this.team.id !== id) {
      this.find({ id })
    }
  }

  editTeam(team) {
    this.$router.push('/teams/' + team.id + '/edit')
  }
}
</script>
<style scoped>
i {
  align-items: inherit !important;
  display: inline !important;
}
</style>
