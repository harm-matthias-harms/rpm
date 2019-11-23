<template>
  <v-snackbar
    v-model="show"
    bottom
    right
    :timeout="6000"
  >
    {{ message }}
    <v-btn
      color="pink"
      text
      @click="show = false"
    >
      <v-icon>close</v-icon>
    </v-btn>
  </v-snackbar>
</template>

<script lang="ts">
  import { Component, Vue } from 'nuxt-property-decorator'

@Component
  export default class Snackbar extends Vue {
  show: boolean = false
  message: string = ''

  created () {
    this.$store.watch(
      state => state.snackbar.message,
      () => {
        const msg = this.$store.state.snackbar.message
        if (msg !== '') {
          this.show = true
          this.message = this.$store.state.snackbar.message
        }
      },
    )
  }
  }
</script>
