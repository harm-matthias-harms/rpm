<template>
  <v-flex lg4 md8 xs12 text-center>
    <v-icon color="success" size="120">fa-check-circle</v-icon>
    <h2 style="margin-top: 10px;">You have signed up successfully</h2>
    <v-btn style="margin-top: 10px;" to="/sign_in" color="primary"
      >Sign in here</v-btn
    >
  </v-flex>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component
export default class RegisterSuccess extends Vue {}
</script>
