<template>
  <v-flex lg4 md8 xs12 text-center>
    <v-icon color="primary" size="120">fa-check-circle</v-icon>
    <h2>
      <br />Sign up successful
      <br />
      <nuxt-link to="/sign_in">Sign in here</nuxt-link>
    </h2>
  </v-flex>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component
export default class RegisterSuccess extends Vue {}
</script>
