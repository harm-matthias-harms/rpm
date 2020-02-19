/* eslint-disable no-undef */

describe('List teams', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/teams',
      status: 200,
      response: 'fixture:team/list.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('lists all teams', () => {
    cy.visit('/teams')
    cy.contains('2')
    cy.contains('test team 1')
    cy.contains('John Doe')
    cy.contains('test team 2')
    cy.contains('otherUser')
  })

  it('filters teams', () => {
    cy.visit('/teams')
    cy.contains('div', 'Search')
      .find('input')
      .first()
      .type('John')
    cy.contains('test team 1')
    cy.contains('John Doe')
    cy.get('test team 2').should('not.exist')
    cy.contains('otherUser').should('not.exist')
  })
})
