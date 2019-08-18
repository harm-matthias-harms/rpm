<template>
  <div>
    <v-card>
      <v-card-title primary-title>
        <h3 class="display-1">Sign In</h3>
      </v-card-title>
      <v-card-text>
        <v-form @submit.prevent="$store.dispatch('user/signin', user)">
          <v-text-field
            v-model="user.username"
            v-validate="{ required: true, min: 6, username: true }"
            :error-messages="errors.collect('username')"
            name="username"
            label="Username or Email"
            data-vv-as="username"
          ></v-text-field>
          <v-text-field
            ref="password"
            v-model="user.password"
            v-validate="{
              required: true,
              min: 6,
              passwordNoWhitespace: true
            }"
            :error-messages="errors.collect('password')"
            label="Password"
            name="password"
            type="password"
            data-vv-as="password"
          ></v-text-field>
          <v-btn
            :disabled="user.username && user.password && errors.any()"
            color="success"
            type="submit"
          >Sign In</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
    <div text-xs-center>
      <nuxt-link to="/sign_up">Create new account?</nuxt-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component
export default class SignIn extends Vue {
  user: object = {
    username: '',
    password: ''
  }

  created() {
    this.$validator.extend('username', {
      getMessage: () =>
        'The username can contain letters (a-z), numbers (0-9), and periods (.).',
      validate: value => !!value.match(/^[a-z1-9.]+$|/)
    })
    this.$validator.extend('passwordNoWhitespace', {
      getMessage: () => "The password can't contain a whitespace.",
      validate: value => !!value.match(/^\S+$/)
    })
  }
}
</script>
