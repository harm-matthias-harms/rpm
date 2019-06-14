import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Logo from '@/components/Logo.vue'
import Vuetify from 'vuetify'

describe('Logo', () => {
  let wrapper
  beforeEach(() => {
    Vue.use(Vuetify)
    wrapper = mount(Logo, {
      stubs: {
        NuxtLink: RouterLinkStub
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
