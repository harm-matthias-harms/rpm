/* eslint-disable no-undef */
describe('Index Page', () => {
  it('it has the possibility to log in', () => {
    cy.visit('/')
    cy.get('.v-toolbar__side-icon')
    cy.contains('.v-toolbar__items > .v-btn', 'sign in')
    cy.contains('button', 'sign in')
    cy.contains('button', 'register')
    cy.contains('button', 'enter code')
  })
  it('toogles the side bar', () => {
    cy.get('.v-toolbar__side-icon').click()
    cy.get('.v-navigation-drawer')
    cy.get('.v-toolbar__content > .v-list > [role="listitem"] > .v-list__tile')
  })
})
