import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import Injects from '@/pages/exercises/_id/injects/index.vue'
import { store } from '../../../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('List injects', () => {
  let wrapper
  let router
  store.state.inject.injectList = {
    injects: [],
    count: 0
  }
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(Injects, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store,
      router
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
