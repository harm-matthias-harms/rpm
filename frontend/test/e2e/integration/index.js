/* eslint-disable no-undef */

describe('Index Page', () => {
  it('it has the possibility to log in', () => {
    cy.visit('/')
    cy.get('.v-app-bar__nav-icon')
    cy.contains('.v-toolbar__items > .v-btn', 'sign in')
    cy.get('input[name="username"]')
    cy.get('input[name="password"]')
    cy.contains('.display-1', 'Sign In')
    cy.contains('a', 'Create new account?')
  })
  it('toogles the side bar', () => {
    cy.get('.v-app-bar__nav-icon').should('not.exist')
    cy.login()
    cy.get('.v-app-bar__nav-icon').click()
    cy.get('.v-navigation-drawer')
    cy.get('.v-list-item')
    cy.logout()
  })
  it('shows and closes cookie bar', () => {
    cy.clearLocalStorage()
    cy.contains('.v-snack__content', 'This site uses 🍪 for your security')
    cy.get('.v-snack__content > .v-btn > .v-btn__content > .v-icon').click()
    cy.get('.v-snack__content').should('not.exist')
  })
  it('no sign in if no connection', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: {},
      status: 404
    })
    cy.loginEnterForm()
    cy.contains("Couldn't connect to network.")
  })
  it('blocks unkown user', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: {},
      status: 401
    })
    cy.loginEnterForm()
    cy.contains('Wrong username or password')
  })
  it('signs in successfully', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: { success: true }
    })
    cy.loginEnterForm()
  })
})
