import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import { store } from '../../../utils/vuex-store'
import ShowPreset from '@/pages/presets/_id/index.vue'

Vue.use(Vuetify)
Vue.use(VueRouter)

describe('Show Presets', () => {
  let wrapper
  let router
  const storeCopy = store
  storeCopy.state.preset.preset = { id: '002' }
  beforeEach(() => {
    router = new VueRouter()
    wrapper = shallowMount(ShowPreset, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      store,
      router,
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
  test('open edit form', () => {
    wrapper.vm.editPreset({ id: '001' })
    expect(wrapper.vm.$route.path).toEqual('/presets/001/edit')
  })
})
