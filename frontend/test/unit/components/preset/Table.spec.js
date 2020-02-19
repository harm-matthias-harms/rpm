import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import PresetTable from '@/components/preset/table.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Preset Table', () => {
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
  test('open edit form', () => {
    wrapper.vm.editPreset({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/presets/001/edit')
  })
  test('filter table', () => {
    wrapper.props().items.push({ title: 'test', author: { username: 'John' } })
    expect(
      wrapper.vm.filterPresets('test', 'John', wrapper.props().items[0])
    ).toBeTruthy()
    expect(
      wrapper.vm.filterPresets('test', 'test Jack', wrapper.props().items[0])
    ).toBeFalsy()
  })
})
