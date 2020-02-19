import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../../utils/vuex-store'
import Edit from '@/pages/medical_cases/_id/edit.vue'

const $route = {
  path: '/medical_cases/001',
  params: { id: '001' }
}

Vue.use(Vuetify)

describe('Edit Medical Case', () => {
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
