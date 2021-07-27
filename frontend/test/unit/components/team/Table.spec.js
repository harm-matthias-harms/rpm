import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import TeamTable from '@/components/team/table.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Team table', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(TeamTable, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        loading: false,
        items: []
      },
      store,
      router
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
    expect(wrapper.props().loading).toBeFalsy()
    expect(wrapper.props().items).toMatchObject([])
  })
  test('opens team', () => {
    wrapper.vm.openTeam({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/teams/001')
  })
  test('open edit form', () => {
    wrapper.vm.editTeam({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/teams/001/edit')
  })
  test('filter table', () => {
    wrapper.props().items.push({ title: 'test', author: { username: 'John' }, type: 'EMT 1', medivac: true })
    expect(
      wrapper.vm.filterTeams('test', 'John medivac EMT 1', wrapper.props().items[0])
    ).toBeTruthy()
    expect(
      wrapper.vm.filterTeams('test', 'test Jack', wrapper.props().items[0])
    ).toBeFalsy()
  })
})
