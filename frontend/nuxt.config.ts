import { Configuration } from '@nuxt/types'

const config: Configuration = {
  mode: 'spa',

  /*
   ** Headers of the page
   */
  head: {
    title: 'RPM',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: 'role play management for medical exercises',
      },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      {
        rel: 'stylesheet',
        href:
          'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons',
      },
    ],
  },

  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },

  /*
   ** Global CSS
   */
  css: ['~/assets/style/app.styl'],

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],

  /*
   ** Router middlewares
   */
  router: {
    middleware: ['auth'],
  },

  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    'nuxt-validate',
    '@nuxtjs/vuetify',
  ],
  /*
   ** Axios module configuration
   */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
    https: true,
    baseURL: 'http://localhost:3001',
    // xsrfCookieName: 'csrf',
    // xsrfHeaderName: 'X-CSRF-Token',
    credentials: true,
    // API_URL must be used for api url
  },
  /*
   ** vuetify module configuration
   */
  vuetify: {
    icons: {
      iconfont: 'fa',
    },
  },
  /*
   ** Build configuration
   */
  build: {
    transpile: [],
    plugins: [],
    loaders: {
      stylus: {
        import: [],
      },
    },
  },
  buildModules: ['@nuxt/typescript-build', '@nuxtjs/moment'],
}

export default config
