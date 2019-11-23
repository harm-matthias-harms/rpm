import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../utils/vuex-store'
import ShowPreset from '@/pages/presets/_id.vue'

const $route = {
  path: '/presets/001',
  params: { id: '001' },
}

Vue.use(Vuetify)

describe('My Presets', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(ShowPreset, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      store,
      router,
      mocks: {
        $route,
      },
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
