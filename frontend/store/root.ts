import { GetterTree, ActionTree, MutationTree } from 'vuex'

export const types = {}

export interface State {}

export interface RootState {}

export const state = (): State => ({})

export const getters: GetterTree<State, State> = {}

export interface Actions<S, R> extends ActionTree<S, R> {}

export const actions: Actions<State, State> = {}

export const mutations: MutationTree<State> = {}
