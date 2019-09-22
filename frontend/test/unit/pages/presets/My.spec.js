import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import MyPresets from '@/pages/presets/my.vue'

Vue.use(Vuetify)

describe('My Presets', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(MyPresets, {
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
