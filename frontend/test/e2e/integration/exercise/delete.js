/* eslint-disable no-undef */

describe('Delete exercise', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'DELETE',
      url: 'http://localhost:3001/api/exercises/001',
      status: 200,
      response: { status: 200 }
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/exercises/001',
      status: 200,
      response: 'fixture:exercise/exercise.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  /*
  until cypress fixes the icon bug
  it('deletes exercise', () => {
    cy.visit('/exercises/001')
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this exercise?')
    cy.get('.v-card__actions').contains('cancel').click()
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this exercise?')
    cy.get('.v-card__actions').contains('ok').click()
    cy.url().should('include', '/')
  }) */
})
