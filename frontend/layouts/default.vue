<template>
  <v-app>
    <v-toolbar dark dense color="primary">
      <v-toolbar-side-icon @click="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>RPM</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat>sign in</v-btn>
      </v-toolbar-items>
    </v-toolbar>

    <v-navigation-drawer v-model="drawer" temporary absolute>
      <v-toolbar flat class="transparent">
        <v-list class="pa-0">
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <img src="https://randomuser.me/api/portraits/men/85.jpg" />
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>John Leider</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-toolbar>
      <v-list class="pt-0" dense>
        <v-divider></v-divider>
        <v-list-tile v-for="item in items" :key="item.title">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <CookieHint />
    <Snackbar />

    <v-content>
      <v-container fill-height>
        <Nuxt v-if="!$store.state.loader.isLoading" />
        <v-layout
          v-if="$store.state.loader.isLoading"
          align-center
          justify-center
          fill-height
        >
          <div>
            <v-progress-circular
              :size="100"
              :width="7"
              color="primary"
              indeterminate
            ></v-progress-circular>
          </div>
        </v-layout>
      </v-container>
    </v-content>

    <v-footer class="pa-3">
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
