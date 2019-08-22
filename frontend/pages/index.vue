<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="4" md="8" sm="12" class="text-center">
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
    </v-row>

    <v-row justify="center">
      <v-col lg="4" md="8" sm="12">
        <SignIn v-if="!this.$store.state.user.isAuthenticated" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import SignIn from '@/components/auth/SignIn.vue'

@Component({
  components: { SignIn }
})
export default class Index extends Vue {}
</script>
