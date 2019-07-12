import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Index from '@/pages/index.vue'

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
