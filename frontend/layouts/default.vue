<template>
  <v-app>
    <v-app-bar app dense dark color="primary">
      <v-app-bar-nav-icon v-if="this.$store.state.user.isAuthenticated" @click="drawer = !drawer" />
      <v-toolbar-title>RPM</v-toolbar-title>
      <v-spacer />
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn v-if="!this.$store.state.user.isAuthenticated" text to="/">
          sign in
        </v-btn>
        <v-btn v-if="this.$store.state.user.isAuthenticated" text @click="signout">
          sign out
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>

    <v-navigation-drawer
      v-if="this.$store.state.user.isAuthenticated"
      v-model="drawer"
      temporary
      absolute
      app
    >
      <v-toolbar text class="transparent">
        <v-list class="pa-0">
          <v-list-item>
            <v-list-item-avatar>
              <v-icon>fa-user</v-icon>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ this.$store.state.user.user.username }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-toolbar>
      <v-list class="pt-0" dense>
        <v-divider />
        <v-list-item v-for="item in items" :key="item.title">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <CookieHint />
    <Snackbar />

    <v-content>
      <Nuxt v-if="!$store.state.loader.isLoading" />
      <v-overlay :value="$store.state.loader.isLoading">
        <v-progress-circular indeterminate size="64" />
      </v-overlay>
    </v-content>

    <v-footer app>
      <v-spacer />
      <a target="_blank" href="https://github.com/harm-matthias-harms/rpm">
        <v-icon>fab fa-github</v-icon>
      </a>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import CookieHint from '@/components/utils/CookieHint.vue'
import Snackbar from '@/components/utils/Snackbar.vue'

interface SidebarItem {
  readonly title: string
  readonly icon: string
}

@Component({
  components: {
    CookieHint,
    Snackbar
  }
})
export default class Default extends Vue {
  items: SidebarItem[] = [
    {
      title: 'Home',
      icon: 'dashboard'
    }
  ]
  drawer: boolean = false
  signout () {
    this.$store.commit('user/LOGOUT')
    this.$router.push('/')
  }
}
</script>
