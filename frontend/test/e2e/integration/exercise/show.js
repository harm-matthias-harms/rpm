/* eslint-disable no-undef */

describe('Show a exercise', () => {
  beforeEach(() => {
    cy.server()
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

  it('shows a exercise', () => {
    cy.visit('/exercises/001')
    // has title
    cy.contains('Test exercise')
    // has author
    cy.contains('John Doe')
    cy.contains("Jan 01'70")
    // has date
    cy.contains('3999-01-01 - 3999-01-04')
    // has teams
    cy.contains('test team 1')
    cy.contains('Team 1 - trainer')
    // has role play manager
    cy.contains('role play manager 1')
    // has make-up center
    cy.contains('MC 1')
    // shows code
    cy.contains('ABCDEF')
  })
})
