<template>
  <v-layout mt-12 column justify-center align-center>
    <v-icon color="primary" size="120">fas fa-file-medical</v-icon>
    <div class="text-center">
      <h1 class="display-3 mt-3 mb-6 font-weight-light">
        Role Play
        <br />Management
      </h1>
    </div>

    <v-layout justify-space-around>
      <v-flex xs12>
        <v-btn to="/sign_up" color="primary">sign up</v-btn>
        <v-btn to="/sign_in" color="secondary">sign in</v-btn>
        <v-btn to="/enter_code" color="primary">enter code</v-btn>
      </v-flex>
    </v-layout>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class Index extends Vue {}
</script>
