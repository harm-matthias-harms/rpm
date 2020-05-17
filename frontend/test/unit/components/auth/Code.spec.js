import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Code from '@/components/auth/Code.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuex)
Vue.use(VueRouter)
Vue.use(Vuetify)

describe('SignIn', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = mount(Code, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      sync: false,
      store,
      router
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
