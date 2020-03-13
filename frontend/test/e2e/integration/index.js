/* eslint-disable no-undef */

describe('Index Page', () => {
  it('it has the possibility to log in', () => {
    cy.visit('/')
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
    cy.contains('John Doe')
    cy.contains('Preset').click()
    cy.get('[href="/presets"] > .v-list-item__content')
    cy.get('[href="/presets/new"] > .v-list-item__content')
    cy.contains('Medical Case').click()
    cy.get('[href="/medical_cases"] > .v-list-item__content')
    cy.get('[href="/medical_cases/new"] > .v-list-item__content')
    cy.contains('New')
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
    cy.loginEnterForm(false)
    cy.contains("Couldn't connect to network")
  })

  it('blocks unkown user', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: {},
      status: 401
    })
    cy.loginEnterForm(false)
    cy.contains('Wrong username, password or code')
  })

  it('signs in successfully', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: { success: true }
    })
    cy.loginEnterForm(true)
    cy.get('Sign In').should('not.exist')
  })
  it('signs in with code', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: { success: true }
    })
    cy.codeForm(true)
    cy.get('Sign In').should('not.exist')
  })
  it('fails with wrong code', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/authenticate',
      response: {},
      status: 401
    })
    cy.codeForm(false)
    cy.contains('Wrong username, password or code')
  })
})
