import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../utils/vuex-store'
import Table from '@/components/medical_case/table.vue'

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
        RouterLink: RouterLinkStub,
      },
      propsData: {
        loading: false,
        items: [],
      },
      store,
      router,
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
})
