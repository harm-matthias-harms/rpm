<template>
  <v-form @submit.prevent="atSubmit(exercise)">
    <v-text-field
      v-model="exercise.title"
      v-validate="'required|max:50'"
      :error-messages="errors.collect('title')"
      label="Title"
      required
      data-vv-name="title"
    />
    <v-row>
      <v-col>
        <v-menu
          v-model="startDateMenu"
          :close-on-content-click="false"
          :nudge-right="40"
          transition="scale-transition"
          offset-y
          min-width="290px"
        >
          <template v-slot:activator="{ on }">
            <v-text-field
              v-model="exercise.startTime"
              label="Starting date"
              prepend-icon="event"
              readonly
              required
              v-on="on"
            />
          </template>
          <v-date-picker
            v-model="exercise.startTime"
            :min="new Date().toISOString().slice(0,10)"
            @input="startDateMenu = false"
          />
        </v-menu>
      </v-col>
      <v-col>
        <v-menu
          v-model="endDateMenu"
          :close-on-content-click="false"
          :nudge-right="40"
          transition="scale-transition"
          offset-y
          min-width="290px"
        >
          <template v-slot:activator="{ on }">
            <v-text-field
              v-model="exercise.endTime"
              label="Ending date"
              prepend-icon="event"
              readonly
              required
              v-on="on"
            />
          </template>
          <v-date-picker
            v-model="exercise.endTime"
            :min="new Date().toISOString().slice(0,10)"
            @input="endDateMenu = false"
          />
        </v-menu>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p class="subtitle-1">
          Teams
        </p>
      </v-col>
      <v-col class="justify-end d-flex">
        <v-btn small color="primary" @click="exercise.teams.push(Object.assign({}, emptyTeam))">
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row v-for="(team, index) in exercise.teams" :key="'team' + index">
      <v-col>
        <v-select
          v-model="team.team"
          :items="teamList.teams"
          label="Select a team"
          item-text="title"
          required
          return-object
        />
      </v-col>
      <v-col>
        <v-autocomplete
          v-model="team.trainer"
          :items="trainer"
          label="Select a trainer"
          item-text="username"
          return-object
          required
        >
          <template v-slot:prepend-item>
            <v-list-item ripple @click="autoGenerateTrainer(team)">
              <v-list-item-content>
                <v-list-item-title>Autogenerate account</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider class="mt-2" />
          </template>
          <template v-slot:item="data">
            <v-list-item-content>
              <v-list-item-title>{{ data.item.username }}</v-list-item-title>
              <v-list-item-subtitle>{{ data.item.email }}</v-list-item-subtitle>
            </v-list-item-content>
          </template>
        </v-autocomplete>
      </v-col>
      <v-col class="d-flex" cols="1" @click="exercise.teams.splice(index, 1)">
        <v-icon color="red">
          delete
        </v-icon>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p class="subtitle-1">
          Role play manager
        </p>
      </v-col>
      <v-col class="justify-end d-flex">
        <v-btn small color="primary" @click="exercise.roleplayManager.push({})">
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row v-for="(rpm, index) in exercise.roleplayManager" :key="'rpm' + index">
      <v-col>
        <v-autocomplete
          v-model="exercise.roleplayManager[index]"
          :items="roleplayManager"
          label="Select a role play manager"
          item-text="username"
          return-object
          required
        >
          <template v-slot:prepend-item>
            <v-list-item
              ripple
              @click="autoGenerateRoleplayManager(exercise.roleplayManager, index)"
            >
              <v-list-item-content>
                <v-list-item-title>Autogenerate account</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider class="mt-2" />
          </template>
          <template v-slot:item="data">
            <v-list-item-content>
              <v-list-item-title>{{ data.item.username }}</v-list-item-title>
              <v-list-item-subtitle>{{ data.item.email }}</v-list-item-subtitle>
            </v-list-item-content>
          </template>
        </v-autocomplete>
      </v-col>
      <v-col class="d-flex" cols="1" @click="exercise.roleplayManager.splice(index, 1)">
        <v-icon color="red">
          delete
        </v-icon>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p class="subtitle-1">
          Make-up center
        </p>
      </v-col>
      <v-col class="justify-end d-flex">
        <v-btn
          small
          color="primary"
          @click="exercise.makeupCenter.push(Object.assign({}, emptyMakeUpCenter))"
        >
          <v-icon small>
            fas fa-plus
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row v-for="(mc, index) in exercise.makeupCenter" :key="'mc' + index">
      <v-col>
        <v-text-field v-model="mc.title" label="Name of make-up center" required />
      </v-col>
      <v-col>
        <v-autocomplete
          v-model="mc.account"
          :items="makeupCenter"
          label="Select an account"
          item-text="username"
          return-object
          required
        >
          <template v-slot:prepend-item>
            <v-list-item ripple @click="autoGenerateMakeUpCenter(mc, index)">
              <v-list-item-content>
                <v-list-item-title>Autogenerate account</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider class="mt-2" />
          </template>
          <template v-slot:item="data">
            <v-list-item-content>
              <v-list-item-title>{{ data.item.username }}</v-list-item-title>
              <v-list-item-subtitle>{{ data.item.email }}</v-list-item-subtitle>
            </v-list-item-content>
          </template>
        </v-autocomplete>
      </v-col>
      <v-col class="d-flex" cols="1" @click="exercise.makeupCenter.splice(index, 1)">
        <v-icon color="red">
          delete
        </v-icon>
      </v-col>
    </v-row>
    <v-btn
      :disabled="!exercise.title || errors.any()"
      class="mr-4"
      type="submit"
      color="primary"
    >
      {{ isNew ? "create" : "edit" }}
    </v-btn>
    <v-btn @click="$router.back()">
      cancel
    </v-btn>
  </v-form>
</template>
<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions, mapState } from 'vuex'
@Component({
  methods: {
    ...mapActions('team', {
      getTeams: 'get_all'
    }),
    ...mapActions('user', {
      getUser: 'get_all'
    })
  },
  computed: {
    ...mapState('team', {
      teamList: 'teamList'
    }),
    ...mapState('user', {
      userList: 'userList'
    })
  }
})
export default class Form extends Vue {
  @Prop({ type: Object, required: true }) readonly exercise!: any
  @Prop({ type: Function, required: true }) readonly atSubmit!: (
    payload
  ) => void

  @Prop({ type: Boolean, required: true }) readonly isNew!: boolean
  startDateMenu: boolean = false
  endDateMenu: boolean = false
  getTeams!: (payload) => void
  getUser!: (payload) => Promise<any>
  userList!: Array<Object>
  trainer: Array<Object> = []
  roleplayManager: Array<Object> = []
  makeupCenter: Array<Object> = []
  emptyTeam: object = { team: {}, trainer: {} }
  emptyMakeUpCenter: object = { title: '', account: {} }

  autoGenerateTrainer (team) {
    if (Object.keys(team.trainer).length === 0) {
      const randomTrainer = {
        username: this.getExercisePrefix() + ' - ' + team.team.title + ' - trainer',
        code: Math.random().toString(36).substring(2, 8).toUpperCase()
      }
      team.trainer = randomTrainer
      this.trainer.unshift(randomTrainer)
    }
  }

  autoGenerateRoleplayManager (rolePlayManager, index) {
    if (Object.keys(rolePlayManager[index]).length === 0) {
      const user = {
        username: this.getExercisePrefix() + ' - role play manager ' + (index + 1),
        code: Math.random().toString(36).substring(2, 8).toUpperCase()
      }
      rolePlayManager[index] = user
      this.roleplayManager.unshift(user)
    }
  }

  autoGenerateMakeUpCenter (mc) {
    if (Object.keys(mc.account).length === 0) {
      const account = {
        username: this.getExercisePrefix() + ' - ' + mc.title,
        code: Math.random().toString(36).substring(2, 8).toUpperCase()
      }
      mc.account = account
      this.makeupCenter.unshift(account)
    }
  }

  getExercisePrefix () {
    return this.exercise.title.match(/\b(\w)/g).join('').toUpperCase()
  }

  mounted () {
    this.getTeams({ disableLoader: true })
    this.getUser({ disableLoader: true }).then(() => {
      this.trainer = [...this.userList]
      this.roleplayManager = [...this.userList]
      this.makeupCenter = [...this.userList]
    })
  }
}
</script>
