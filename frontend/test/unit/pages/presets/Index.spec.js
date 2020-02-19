import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Presets from '@/pages/presets/index.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('Index Preset', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Presets, {
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
