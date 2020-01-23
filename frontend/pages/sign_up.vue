<template>
  <v-container>
    <v-row
      v-if="!registerSuccess"
      justify="center"
    >
      <v-col
        lg="4"
        md="8"
        sm="12"
      >
        <v-card>
          <v-card-title primary-title>
            <h3 class="display-1">
              Sign Up
            </h3>
          </v-card-title>
          <v-card-text>
            <v-form @submit.prevent="register(user)">
              <v-text-field
                v-model="user.username"
                v-validate="{ required: true, min: 6, username: true }"
                :error-messages="errors.collect('username')"
                name="username"
                label="Username"
                data-vv-as="username"
              />
              <v-text-field
                v-model="user.email"
                v-validate="'required|email'"
                :error-messages="errors.collect('email')"
                name="email"
                label="E-mail address"
                data-vv-as="email"
              />
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
              />
              <v-text-field
                v-model="user.passwordConfirm"
                v-validate="'required|confirmed:password'"
                :error-messages="errors.collect('password_confirm')"
                label="Confirm password"
                name="password_confirm"
                type="password"
                data-vv-as="password"
              />
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
              >
                Sign Up
              </v-btn>
            </v-form>
            <div>
              By clicking sign up, you agree that your information will
              be stored for further authentification.
            </div>
          </v-card-text>
        </v-card>
        <div text-xs-center>
          Already have an account?
          <nuxt-link to="/">
            Sign In here
          </nuxt-link>
        </div>
      </v-col>
    </v-row>
    <v-row
      v-if="registerSuccess"
      justify="center"
    >
      <v-col
        lg="4"
        md="8"
        sm="12"
        class="text-center"
      >
        <RegisterSuccess />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator'
  import { mapState, mapActions } from 'vuex'
  import RegisterSuccess from '@/components/auth/RegisterSuccess.vue'

@Component({
  components: { RegisterSuccess },
  computed: {
    ...mapState('user', {
      registerError: 'registerError',
      registerErrorReason: 'registerErrorReason',
      storedUser: 'user',
      registerSuccess: 'registerSuccess',
    }),
  },
  methods: {
    ...mapActions('user', {
      register: 'register',
    }),
  },
})
  export default class SignUp extends Vue {
  registerError!: boolean
  registerErrorReason!: string
  storedUser!: any

  user: object = {
    username: '',
    email: '',
    password: '',
    passwordConfirm: '',
  }

  created () {
    this.$validator.extend('username', {
      getMessage: () =>
        'The username can contain letters (a-z), numbers (0-9), and periods (.)',
      validate: value => !!value.match(/^[a-z1-9.]+$/),
    })
    this.$validator.extend('passwordNoWhitespace', {
      getMessage: () => "The password can't contain a whitespace",
      validate: value => !!value.match(/^\S+$/),
    })
  }

  mounted () {
    if (this.registerError) {
      if (this.registerErrorReason === 'already exists') {
        this.$router.push('/')
      }
      this.user = {
        username: this.storedUser.username,
        email: this.storedUser.email,
        password: this.storedUser.password,
      }
    }
  }
  }
</script>
