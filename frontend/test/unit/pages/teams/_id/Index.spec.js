import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import ShowTeam from '@/pages/teams/_id/index.vue'
import { store } from '../../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Show team', () => {
  let wrapper
  let router
  const storeCopy = store
  storeCopy.state.team.team = { id: '002', author: { id: '001', username: 'username' }, editor: { id: '001', username: 'username' } }
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(ShowTeam, {
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
  test('open edit form', () => {
    wrapper.vm.editTeam({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/teams/001/edit')
  })
})
