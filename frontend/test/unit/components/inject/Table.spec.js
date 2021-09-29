import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import Table from '@/components/inject/Table'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Medical case table', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(Table, {
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
  test('open inject', () => {
    wrapper.vm.openInject({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/exercises/undefined/injects/001')
  })
  test('open edit inject', () => {
    wrapper.vm.editInject({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual(
      '/exercises/undefined/injects/001/edit'
    )
  })
  test('filter table', () => {
    const inject = {
      status: 'status',
      startTime: new Date('2020-12-17T02:00:00'),
      roleplayer: { fullName: 'John Doe' },
      medicalCase: { title: 'mc' },
      team: { title: 'team' },
      makeupCenterTitle: 'makeup'
    }
    expect(wrapper.vm.filterInjects('test', 'status', inject)).toBeTruthy()
    expect(wrapper.vm.filterInjects('test', 'John', inject)).toBeTruthy()
    expect(wrapper.vm.filterInjects('test', 'mc', inject)).toBeTruthy()
    expect(wrapper.vm.filterInjects('test', 'team', inject)).toBeTruthy()
    expect(wrapper.vm.filterInjects('test', 'makeup', inject)).toBeTruthy()
    expect(wrapper.vm.filterInjects('test', '02:00', inject)).toBeTruthy()
  })
})
