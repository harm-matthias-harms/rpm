<template>
  <v-container>
    <v-row justify="center">
      <v-col
        lg="4"
        md="8"
        sm="12"
        class="text-center"
      >
        <v-icon
          color="primary"
          size="120"
        >
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

    <v-row v-if="!isAuthenticated" justify="center">
      <v-col
        lg="4"
        md="5"
        sm="12"
        cols="12"
      >
        <SignIn />
      </v-col>
      <v-col
        lg="1"
        md="1"
        sm="1"
        cols="0"
        class="text-center"
      >
        <v-divider vertical class="d-none d-md-inline-flex" />
      </v-col>
      <v-col
        lg="4"
        md="5"
        sm="12"
        cols="12"
      >
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
      isAuthenticated: 'isAuthenticated'
    })
  }
})
export default class Index extends Vue {}
</script>
