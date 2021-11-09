<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="8" md="10" sm="12">
        <v-card>
          <v-card-text v-if="exercise.id === $route.params.id">
            <h4 class="display-1 font-weight-light mb-0 black--text">
              {{ exercise.title }}
              <v-icon
                v-if="exercise.author.id == $store.state.user.user.id"
                color="primary"
                @click="editExercise(exercise)"
              >
                edit
              </v-icon>
              <DeleteButton
                v-if="exercise.author.id == $store.state.user.user.id"
                :item="exercise"
              />
              <v-btn
                class="float-right"
                color="primary"
                @click="openInjects(exercise)"
              >
                Enter
              </v-btn>
            </h4>
            <p class="body-2 ml-2 mb-4">
              {{ exercise.startTime.slice(0, 10) }} -
              {{ exercise.endTime.slice(0, 10) }}
            </p>

            <v-row
              v-if="exercise.author.id == $store.state.user.user.id"
              class="mb-2"
            >
              <v-col
                v-if="exercise.author && exercise.author.username"
                class="col-auto"
              >
                <Author
                  :author="exercise.author"
                  :created-at="exercise.createdAt"
                />
              </v-col>
            </v-row>
            <div v-if="exercise.author.id == $store.state.user.user.id">
              <v-card class="mb-2">
                <v-card-text>
                  <h2 class="font-weight-light black--text">Teams</h2>
                  <v-row
                    v-for="(team, index) in exercise.teams"
                    :key="'team' + index"
                  >
                    <v-col>
                      <h3 class="font-weight-light black--text">
                        {{ team.team.title }}
                      </h3>
                      <p>
                        {{ team.team.type }}
                        {{ team.team.medivac ? '(MEDIVAC)' : null }}
                      </p>
                    </v-col>
                    <v-col>
                      <h3 class="font-weight-light black--text">Trainer</h3>
                      <p>
                        {{ team.trainer.username }}
                        <br />
                        {{
                          team.trainer.email
                            ? team.trainer.email
                            : 'Code: ' + team.trainer.code
                        }}
                      </p>
                    </v-col>
                  </v-row>
                </v-card-text>
              </v-card>
              <v-card class="mb-2">
                <v-card-text>
                  <h2 class="font-weight-light black--text">
                    Role Play Manager
                  </h2>
                  <v-row
                    v-for="(rpm, index) in exercise.roleplayManager"
                    :key="'rpm' + index"
                  >
                    <v-col>
                      <h3 class="font-weight-light black--text">
                        {{ rpm.username }}
                        <br />
                        {{ rpm.email ? rpm.email : 'Code: ' + rpm.code }}
                      </h3>
                    </v-col>
                  </v-row>
                </v-card-text>
              </v-card>
              <v-card class="mb-2">
                <v-card-text>
                  <h2 class="font-weight-light black--text">Make-Up Center</h2>
                  <v-row
                    v-for="(mc, index) in exercise.makeupCenter"
                    :key="'team' + index"
                  >
                    <v-col>
                      <h3 class="font-weight-light black--text">
                        {{ mc.title }}
                      </h3>
                    </v-col>
                    <v-col>
                      <h3 class="font-weight-light black--text">Account</h3>
                      <p>
                        {{ mc.account.username }}
                        <br />
                        {{
                          mc.account.email
                            ? mc.account.email
                            : 'Code: ' + mc.account.code
                        }}
                      </p>
                    </v-col>
                  </v-row>
                </v-card-text>
              </v-card>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Author from '@/components/utils/Author.vue'
import DeleteButton from '@/components/exercise/Delete.vue'
@Component({
  components: {
    Author,
    DeleteButton,
  },
  computed: {
    ...mapState('exercise', {
      exercise: 'exercise',
    }),
  },
  methods: {
    ...mapActions('exercise', {
      find: 'find',
    }),
  },
})
export default class ShowExercise extends Vue {
  find!: (id) => void
  exercise!: any

  async mounted() {
    const id = this.$route.params.id
    if (this.exercise.id !== id) {
      await this.find({ id })
    }
    if (
      this.exercise.id &&
      this.exercise.author.id !== this.$store.state.user.user.id
    ) {
      this.openInjects(this.exercise)
    }
  }

  editExercise(exercise) {
    this.$router.push('/exercises/' + exercise.id + '/edit')
  }

  openInjects(exercise) {
    this.$router.push(`/exercises/${exercise.id}/injects`)
  }
}
</script>
