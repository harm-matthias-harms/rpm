import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import Table from '@/components/medical_case/table.vue'
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
  test('open preset', () => {
    wrapper.vm.openMedicalCase({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/medical_cases/001')
  })
  test('open edit form', () => {
    wrapper.vm.editMedicalCase({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/medical_cases/001/edit')
  })
  test('filter table', () => {
    const mc = { author: {}, general: { discipline: 'disc', context: ['context'], scenario: ['scenario'] }, patient: { triage: 'Red' } }
    expect(
      wrapper.vm.filterMedicalCases('test', 'disc', mc)
    ).toBeTruthy()
    expect(
      wrapper.vm.filterMedicalCases('test', 'context', mc)
    ).toBeTruthy()
    expect(
      wrapper.vm.filterMedicalCases('test', 'something', mc)
    ).toBeFalsy()
    expect(
      wrapper.vm.filterMedicalCases('test', 'scenario', mc)
    ).toBeTruthy()
    expect(
      wrapper.vm.filterMedicalCases('test', 'red', mc)
    ).toBeTruthy()
  })
})
