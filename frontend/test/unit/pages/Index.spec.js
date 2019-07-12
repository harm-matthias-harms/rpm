import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Index from '@/pages/index.vue'
import Vuetify from 'vuetify'

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    Vue.use(Vuetify)
    wrapper = shallowMount(Index, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
