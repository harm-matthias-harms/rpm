import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import DeleteButton from '@/components/preset/Delete.vue'

Vue.use(Vuetify)

describe('Delete Button', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(DeleteButton, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      propsData: {
        item: { id: '001', vitalSigns: {} },
        goBack: false,
      },
      sync: false,
      store,
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    wrapper.vm.confirmDeletePreset({})
    wrapper.vm.deletePreset({ id: '001' })
    wrapper.vm.onCancel()
  })
})
