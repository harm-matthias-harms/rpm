import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import Edit from '@/pages/presets/_id/edit.vue'
import { store } from '../../../utils/vuex-store'

const $route = {
  path: '/presets/001',
  params: { id: '001' }
}

Vue.use(Vuetify)

describe('Edit Presets', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(Edit, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store,
      router,
      mocks: {
        $route
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
