import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../../utils/vuex-store'
import ShowMedicalCase from '@/pages/medical_cases/_id/index.vue'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Show Medical Case', () => {
  let wrapper
  let router
  store.state.medicalCase.medicalCase = {
    id: '002',
    author: { id: '001' },
    generalInformation: { usar: true },
    medicalHistory: { problems: 'problem' },
    expectations: { expectations: 'expectations' },
    vitalSigns: [{ title: true }]
  }
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(ShowMedicalCase, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store,
      router
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    expect(wrapper.vm.bytesToSize(10000)).toBe('9.77 KB')
  })
  test('open edit form', () => {
    wrapper.vm.editMedicalCase({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/medical_cases/001/edit')
  })
})
