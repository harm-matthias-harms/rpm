<template>
  <v-container>
    <v-row v-if="isAuthenticated && isLoaded" align="center">
      <v-col cols="12" lg="3" md="4" sm="6" class="text-center">
        <v-card @click="$router.push('/exercises/new')">
          <v-card-text>
            <v-icon class="my-3">
              fa fa-plus
            </v-icon>
            <p class="subtitle-1">
              New Exercise
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <div v-if="user.roles">
        <v-col
          v-for="(role, index) in user.roles"
          :key="index"
          cols="12"
          lg="3"
          md="4"
          sm="6"
        >
          <nuxt-link :to="'/exercises/' + role.exercise.id">
            <v-card>
              <v-card-text>
                <p class="title mb-0 text--primary">
                  {{ role.exercise.title }}
                </p>
                <p>{{ role.exercise.startTime.slice(0,10) }} - {{ role.exercise.endTime.slice(0,10) }}</p>
                <p class="mb-0">
                  {{ role.role.charAt(0).toUpperCase() + role.role.slice(1) }}
                </p>
              </v-card-text>
            </v-card>
          </nuxt-link>
        </v-col>
      </div>
    </v-row>

    <v-row v-if="!isAuthenticated" justify="center" style="margin-top: auto; margin-bottom: auto;">
      <v-col cols="12" class="text-center">
        <v-icon color="primary" size="120">
          fas fa-file-medical
        </v-icon>
        <div>
          <h1 class="display-3 mt-3 mb-6 font-weight-light">
            Role Play
            <br>Management
          </h1>
        </div>
      </v-col>
      <v-col lg="4" md="5" sm="12" cols="12">
        <SignIn />
      </v-col>
      <v-col lg="1" md="1" sm="1" cols="0" class="text-center">
        <v-divider vertical class="d-none d-md-inline-flex" />
      </v-col>
      <v-col lg="4" md="5" sm="12" cols="12">
        <Code />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { mapState } from 'vuex'
import SignIn from '@/components/auth/SignIn.vue'
import Code from '@/components/auth/Code.vue'

@Component({
  components: { SignIn, Code },
  computed: {
    ...mapState('user', {
      isAuthenticated: 'isAuthenticated',
      user: 'user',
      isLoaded: 'isLoaded'
    })
  }
})
export default class Index extends Vue {}
</script>

<style scoped>
a {
  color: inherit;
  text-decoration: inherit;
}
</style>
