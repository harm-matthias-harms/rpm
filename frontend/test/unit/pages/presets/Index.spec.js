import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import Presets from '@/pages/presets/index.vue'

Vue.use(Vuetify)

describe('Index Preset', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Presets, {
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
