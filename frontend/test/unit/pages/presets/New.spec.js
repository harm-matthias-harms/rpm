import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import NewPreset from '@/pages/presets/new.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('New Preset', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(NewPreset, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
