import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import { store } from '../utils/vuex-store'
import Index from '@/pages/index.vue'

Vue.use(Vuex)
Vue.use(Vuetify)

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Index, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      store,
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
