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
    expect(wrapper.isVueInstance()).toBeTruthy()
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
  test('get tags', () => {
    const mc = { generalInformation: { surgical: true, usar: true, medivac: true, hospilisation: true, triage: 'Urgent', age: '18-30', gender: 'Male' } }
    expect(wrapper.vm.tags(mc)).toEqual(['USAR', 'MEDIVAC', 'Need for hospilisation', 'Surgical', 'Triage: Urgent', 'Age: 18-30', 'Gender: Male'])
  })
  test('filter table', () => {
    const mc = { author: {}, generalInformation: { surgical: true, usar: true, medivac: true, hospilisation: false, triage: 'Urgent', age: '18-30', gender: 'Male' } }
    expect(
      wrapper.vm.filterMedicalCases('test', 'surgical 18', mc)
    ).toBeTruthy()
    mc.generalInformation.age = '60+'
    expect(
      wrapper.vm.filterMedicalCases('test', 'surgical 70', mc)
    ).toBeTruthy()
    mc.generalInformation.age = null
    expect(
      wrapper.vm.filterMedicalCases('test', 'surgical usar medivac urgent 70', mc)
    ).toBeFalsy()
    expect(
      wrapper.vm.filterMedicalCases('test', 'hospital', mc)
    ).toBeFalsy()
  })
})
