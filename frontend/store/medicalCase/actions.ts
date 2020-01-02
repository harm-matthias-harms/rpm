import { ActionTree } from 'vuex/types'
import { State } from './type'
import { State as RootState } from '@/store/root'

export const actions: ActionTree<State, RootState> = {
  create ({ commit }, payload) {
    const medicalCase = payload.medicalCase
    const files = payload.files
    const form = new FormData()
    form.append('medicalCase', JSON.stringify(medicalCase))
    files.forEach((file) => {
      form.append('files', file)
    })
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/api/medical_cases', form)
      .then((response) => {
        commit('SET_MEDICAL_CASE_TO_LIST', response.data)
        this.$router.push('/medical_cases')
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create medical case.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  get_all ({ commit }) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$get('/api/medical_cases')
      .then((response) => {
        commit('SET_MEDICAL_CASE_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load medical cases.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  find ({ commit }, id) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$get('/api/medical_cases/' + id)
      .then((response) => {
        commit('SET_MEDICAL_CASE', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find medical case.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
}

export default actions
