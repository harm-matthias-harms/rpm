import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../utils/vuex-store'
import PresetTable from '@/components/preset/table.vue'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('My Presets', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(PresetTable, {
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
    wrapper.vm.openPreset({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/presets/001')
  })
})
