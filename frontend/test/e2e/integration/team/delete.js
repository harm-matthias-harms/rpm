/* eslint-disable no-undef */

describe('Delete team', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'DELETE',
      url: 'http://localhost:3001/api/teams/001',
      status: 200,
      response: { status: 200 }
    })
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

  /*
  until cypress fixes the icon bug
  it('deletes team', () => {
    cy.visit('/teams')
    cy.get('i').contains('delete')
    cy.visit('/teams/001')
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this team?')
    cy.get('.v-card__actions').contains('cancel').click()
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this team?')
    cy.get('.v-card__actions').contains('ok').click()
    cy.url().should('include', '/teams')
  }) */
})
