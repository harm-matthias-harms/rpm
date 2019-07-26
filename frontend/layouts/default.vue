<template>
  <v-app>
    <v-app-bar app dense dark color="primary">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>RPM</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn text to="/signin">sign in</v-btn>
      </v-toolbar-items>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" temporary absolute app>
      <v-toolbar text class="transparent">
        <v-list class="pa-0">
          <v-list-item>
            <v-list-item-avatar>
              <img src="https://randomuser.me/api/portraits/men/85.jpg" />
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>John Leider</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-toolbar>
      <v-list class="pt-0" dense>
        <v-divider></v-divider>
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
      <v-container fill-height>
        <Nuxt v-if="!$store.state.loader.isLoading" />
        <v-overlay :value="$store.state.loader.isLoading">
          <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>
      </v-container>
    </v-content>

    <v-footer app>
      <v-spacer></v-spacer>
      <a href="https://github.com/harm-matthias-harms/rpm">
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
}
</script>
