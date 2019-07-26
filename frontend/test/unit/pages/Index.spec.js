import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Index from '@/pages/index.vue'

Vue.use(Vuetify)

describe('Index', () => {
  let wrapper
  beforeEach(() => {
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
