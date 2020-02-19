/* eslint-disable no-undef */

describe('Show a team', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/teams',
      status: 200,
      response: 'fixture:team/list.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/teams/001',
      status: 200,
      response: 'fixture:team/team.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('shows a team', () => {
    cy.visit('/teams')
    cy.contains('test team 1').click()
    // has title
    cy.contains('Test team')
    // has author
    cy.contains('John Doe')
    cy.contains("Jan 01'70")
    // contains prefix
    cy.contains('Type:')
    cy.contains('Medivac:')
    // contains values
    cy.contains('EMT 1')
    cy.contains('done')
  })
})
