<template>
  <v-app>
    <v-app-bar
      app
      dense
      dark
      color="primary"
    >
      <v-app-bar-nav-icon
        v-if="isAuthenticated"
        @click="drawer = !drawer"
      />
      <v-toolbar-title>RPM</v-toolbar-title>
      <v-spacer />
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn
          v-if="!isAuthenticated"
          text
          to="/"
        >
          sign in
        </v-btn>
        <v-btn
          v-if="isAuthenticated"
          text
          @click="signout"
        >
          sign out
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>

    <v-navigation-drawer
      v-if="isAuthenticated"
      v-model="drawer"
      temporary
      app
    >
      <v-list
        nav
        class="px-0"
      >
        <v-list-item>
          <v-list-item-avatar>
            <v-icon>fa-user</v-icon>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>{{ user.username }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-divider />
      <v-list
        dense
        nav
      >
        <v-list-group
          v-for="item in items"
          :key="item.name"
          :prepend-icon="item.icon"
          no-action
        >
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="item.name" />
            </v-list-item-content>
          </template>

          <v-list-item
            v-for="subItem in item.items"
            :key="subItem.name"
            :to="subItem.url"
          >
            <v-list-item-icon>
              <v-icon>{{ subItem.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title v-text="subItem.name" />
            </v-list-item-content>
          </v-list-item>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>

    <CookieHint />
    <Snackbar />

    <v-content>
      <Nuxt v-if="!isLoading" />
      <v-overlay :value="isLoading">
        <v-progress-circular
          indeterminate
          size="64"
        />
      </v-overlay>
    </v-content>

    <v-footer app>
      <v-spacer />
      <a
        target="_blank"
        href="https://github.com/harm-matthias-harms/rpm"
      >
        <v-icon>fab fa-github</v-icon>
      </a>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapMutations } from 'vuex'
import CookieHint from '@/components/utils/CookieHint.vue'
import Snackbar from '@/components/utils/Snackbar.vue'

@Component({
  components: {
    CookieHint,
    Snackbar
  },
  computed: {
    ...mapState('loader', {
      isLoading: 'isLoading'
    }),
    ...mapState('user', {
      user: 'user',
      isAuthenticated: 'isAuthenticated'
    })
  },
  methods: {
    ...mapMutations('user', {
      logout: 'LOGOUT'
    })
  }
})
export default class Default extends Vue {
  logout!: () => void

  items = [
    {
      name: 'Medical Case',
      icon: 'fas fa-file-medical',
      items: [
        {
          name: 'List',
          icon: 'fas fa-table',
          url: '/medical_cases'
        },
        {
          name: 'New',
          icon: 'fas fa-plus',
          url: '/medical_cases/new'
        }
      ]
    },
    {
      name: 'Preset',
      icon: 'fas fa-syringe',
      items: [
        {
          name: 'List',
          icon: 'fas fa-table',
          url: '/presets'
        },
        {
          name: 'New',
          icon: 'fas fa-plus',
          url: '/presets/new'
        }
      ]
    }
  ]

  drawer: boolean = false

  signout () {
    this.logout()
    this.$router.push('/')
  }
}
</script>
