/* eslint-disable no-undef */

describe('Create team', () => {
  let team
  beforeEach(() => {
    cy.fixture('team/create.json').then((json) => {
      team = json.data
    })
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/api/teams',
      status: 200,
      response: 'fixture:team/create.json'
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

  it('creates team', () => {
    cy.visit('/teams/new')
    cy.contains('New Team')
    cy.contains('div', 'Title').find('input').first().type(team.title)
    cy.contains('div', 'Type').first('div[role="button"]').click()
    cy.contains('div', team.type).first('div[role="option"]').click()
    cy.contains('div', 'Medivac').click()
    cy.contains('create').click()
    cy.contains(team.title)
    cy.contains(team.author.username)
    cy.url().should('include', '/teams')
  })
})
