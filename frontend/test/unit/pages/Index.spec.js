import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Index from '@/pages/index.vue'
import Vuetify from 'vuetify'

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    Vue.use(Vuetify)
    wrapper = mount(Index, {
      stubs: {
        NuxtLink: RouterLinkStub
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
