<template>
  <v-layout mt-5 column justify-center align-center>
    <v-icon color="primary">fas fa-10x fa-file-medical</v-icon>
    <div class="text-xs-center text-lg-left">
      <h1 class="display-3 mb-4 font-weight-light">Role Play Management</h1>
    </div>

    <v-layout justify-center wrap>
      <v-btn to="/sign_up" color="primary">sign up</v-btn>
      <v-btn to="/sign_in" color="secondary">sign in</v-btn>
      <v-btn to="/enter_code" color="primary">enter code</v-btn>
    </v-layout>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class Index extends Vue {}
</script>
