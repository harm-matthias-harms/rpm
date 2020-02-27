import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import Edit from '@/pages/teams/_id/edit.vue'
import { store } from '../../../utils/vuex-store'

const $route = {
  path: '/teams/001',
  params: { id: '001' }
}

Vue.use(Vuetify)

describe('Edit team', () => {
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
