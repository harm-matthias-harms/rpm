/* eslint-disable no-undef */

describe('Edit team', () => {
  let team
  beforeEach(() => {
    cy.fixture('team/edit.json').then(json => {
      team = json
    })
    cy.server()
    cy.route({
      method: 'PUT',
      url: 'http://localhost:3001/api/teams/001',
      status: 200,
      response: 'fixture:team/edit.json'
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

  it('updates team', () => {
    cy.visit('/teams/001/edit')
    cy.contains('Edit team')
    cy.contains('div', 'Medivac').click()
    cy.contains('edit').click()
    cy.contains(team.title)
    cy.contains(team.author.username)
    cy.contains(team.editor.username)
    cy.contains('clear')
    cy.url().should('include', '/teams/001')
  })
})
