<template>
  <v-layout wrap justify-center align-center>
    <v-flex lg4 md8 xs12>
      <v-card>
        <v-card-title primary-title>
          <h3 class="display-1">Sign Up</h3>
        </v-card-title>
        <v-card-text>
          <v-form @submit.prevent="$store.dispatch('user/register', user)">
            <v-text-field
              v-model="user.username"
              v-validate="'required|min:6'"
              :error-messages="errors.collect('username')"
              name="username"
              label="Username"
              data-vv-as="username"
            ></v-text-field>
            <v-text-field
              v-model="user.email"
              v-validate="'required|email'"
              :error-messages="errors.collect('email')"
              name="email"
              label="Email Adress"
              data-vv-as="email"
            ></v-text-field>
            <v-text-field
              ref="password"
              v-model="user.password"
              v-validate="'required|min:6'"
              :error-messages="errors.collect('password')"
              label="Password"
              name="password"
              type="password"
              data-vv-as="password"
            ></v-text-field>
            <v-text-field
              v-model="user.passwordConfirm"
              v-validate="'required|confirmed:password'"
              :error-messages="errors.collect('password_confirm')"
              label="Confirm Password"
              name="password_confirm"
              type="password"
              data-vv-as="password"
            ></v-text-field>
            <v-btn
              :disabled="
                user.username &&
                  user.email &&
                  user.password &&
                  user.passwordConfirm &&
                  errors.any()
              "
              color="success"
              type="submit"
              >Sign Up</v-btn
            >
          </v-form>
          <div>
            By clicking Sign Up button, you agree that your information will be
            stored for further authentification.
          </div>
        </v-card-text>
      </v-card>
      <div text-xs-center>
        Already have an account?
        <nuxt-link to="/sign_in">Sign In here</nuxt-link>
      </div>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class SignUp extends Vue {
  user: object = {
    username: '',
    email: '',
    password: '',
    passwordConfirm: ''
  }

  mounted() {
    if (this.$store.state.user.registerError) {
      if (this.$store.state.user.registerErrorReason === 'already exists') {
        this.$router.push('/sign_in')
      }
      this.user = {
        username: this.$store.state.user.user.username,
        email: this.$store.state.user.user.email,
        password: this.$store.state.user.user.password
      }
    }
    if (this.$store.state.user.registerSuccess) {
      this.$router.push('/account_created')
    }
  }
}
</script>
