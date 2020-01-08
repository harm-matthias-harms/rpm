<template>
  <v-snackbar
    v-model="snackbar"
    bottom
    right
    :timeout="0"
  >
    This site uses 🍪 for your security
    <v-btn
      color="pink"
      text
      @click="snackbar = false"
    >
      <v-icon @click="acceptCookieHint">
        close
      </v-icon>
    </v-btn>
  </v-snackbar>
</template>

<script lang="ts">
  import { Component, Vue } from 'nuxt-property-decorator'
@Component
  export default class CookieHint extends Vue {
  snackbar: boolean = false

  acceptCookieHint () {
    localStorage.setItem('accepted_cookie_hint', 'true')
    this.snackbar = false
  }

  mounted () {
    this.snackbar = !(localStorage.getItem('accepted_cookie_hint') === 'true')
  }
  }
</script>
