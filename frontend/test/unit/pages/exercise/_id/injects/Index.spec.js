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
  store.state.exercise.exercise = {
    id: '002',
    author: { id: '001' },
    startTime: '2001-01-01',
    endTime: '2001-01-01',
    teams: [],
    roleplayManager: [],
    makeupCenter: []
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
