import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import DeleteButton from '@/components/team/Delete.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('Delete Button', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(DeleteButton, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        item: { id: '001', vitalSigns: {} },
        goBack: false
      },
      sync: false,
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
    wrapper.vm.deleteTeam({ id: '001' })
    wrapper.vm.onCancel()
  })
})
