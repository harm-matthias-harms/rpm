import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import ShowMedicalCase from '@/pages/medical_cases/_id/index.vue'
import { store } from '../../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Show Medical Case', () => {
  let wrapper
  let router
  store.state.medicalCase.medicalCase = {
    id: '002',
    author: { id: '001' },
    general: { context: ['context'] },
    patient: { type: 'Acute' },
    medical: { signs: 'signs' },
    makeup: { makeup: 'makeup' },
    vitalSigns: [{ title: 'title' }]
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
    expect(wrapper).toBeTruthy()
    expect(wrapper.vm.bytesToSize(10000)).toBe('9.77 KB')
    expect(wrapper.vm.bytesToSize(0)).toBe('n/a')
    expect(wrapper.vm.bytesToSize(1)).toBe('1 Bytes')
  })

  test('open edit form', () => {
    wrapper.vm.editMedicalCase({ id: '001' })
    wrapper.vm.onCancel()
    expect(wrapper.vm.$route.path).toEqual('/medical_cases/001/edit')
  })
})
