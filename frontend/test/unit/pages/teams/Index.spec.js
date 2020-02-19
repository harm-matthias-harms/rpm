import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import Teams from '@/pages/teams/index.vue'

Vue.use(Vuetify)

describe('Index team', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Teams, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
