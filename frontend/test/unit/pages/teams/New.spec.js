import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import NewTeam from '@/pages/teams/new.vue'

Vue.use(Vuetify)

describe('New team', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(NewTeam, {
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
