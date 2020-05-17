import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import Form from '@/components/exercise/Form.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Exercise form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Form, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        exercise: { title: 'some random title', teams: [], roleplayManager: [], makeupCenter: [] },
        atSubmit: () => {},
        isNew: false
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        }
      },
      sync: false,
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
  test('autogenerate trainer', () => {
    wrapper.vm.exercise.teams.push({ team: { title: 'Team 1' }, trainer: {} })
    wrapper.vm.autoGenerateTrainer(wrapper.vm.exercise.teams[0])
    expect(wrapper.vm.exercise.teams[0].trainer.username).toBe(
      'SRT - Team 1 - trainer'
    )
    expect(wrapper.vm.exercise.teams[0].trainer.code).not.toBeUndefined()
    expect(wrapper.vm.exercise.teams[0].trainer.code.length).toBe(6)
    expect(wrapper.vm.trainer.length).toBe(1)
  })
  test('autogenerate role play manager', () => {
    wrapper.vm.exercise.roleplayManager.push({})
    wrapper.vm.autoGenerateRoleplayManager(
      wrapper.vm.exercise.roleplayManager,
      0
    )
    expect(wrapper.vm.exercise.roleplayManager[0].username).toBe(
      'SRT - role play manager 1'
    )
    expect(wrapper.vm.exercise.roleplayManager[0].code).not.toBeUndefined()
    expect(wrapper.vm.exercise.roleplayManager[0].code.length).toBe(6)
    expect(wrapper.vm.roleplayManager.length).toBe(1)
  })
  test('autogenerate make-up center', () => {
    wrapper.vm.exercise.makeupCenter.push({ title: 'Test', account: {} })
    wrapper.vm.autoGenerateMakeUpCenter(wrapper.vm.exercise.makeupCenter[0])
    expect(wrapper.vm.exercise.makeupCenter[0].account.username).toBe('SRT - Test')
    expect(wrapper.vm.exercise.makeupCenter[0].account.code).not.toBeUndefined()
    expect(wrapper.vm.exercise.makeupCenter[0].account.code.length).toBe(6)
    expect(wrapper.vm.makeupCenter.length).toBe(1)
  })
  test('get right exercise prefix', () => {
    expect(wrapper.vm.getExercisePrefix()).toBe('SRT')
  })
})
