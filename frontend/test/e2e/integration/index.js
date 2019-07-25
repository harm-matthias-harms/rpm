/* eslint-disable no-undef */

describe('Index Page', () => {
  it('it has the possibility to log in', () => {
    cy.visit('/')
    cy.get('.v-app-bar__nav-icon')
    cy.contains('.v-toolbar__items > .v-btn', 'sign in')
    cy.contains('a', 'sign in')
    cy.contains('a', 'sign up')
    cy.contains('a', 'enter code')
  })
  it('toogles the side bar', () => {
    cy.get('.v-app-bar__nav-icon').click()
    cy.get('.v-navigation-drawer')
    cy.get('.v-list-item')
  })
  it('shows and closes cookie bar', () => {
    cy.clearLocalStorage()
    cy.contains('.v-snack__content', 'This site uses 🍪 for your security')
    cy.get('.v-snack__content > .v-btn > .v-btn__content > .v-icon').click()
    cy.get('.v-snack__content').should('not.exist')
  })
})
