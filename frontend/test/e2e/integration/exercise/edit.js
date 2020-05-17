/* eslint-disable no-undef */

describe('Edit exercise', () => {
  let exercise
  beforeEach(() => {
    cy.fixture('exercise/edit.json').then((json) => {
      exercise = json
    })
    cy.server()
    cy.route({
      method: 'PUT',
      url: 'http://localhost:3001/api/exercises/001',
      status: 200,
      response: 'fixture:exercise/edit.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/exercises/001',
      status: 200,
      response: 'fixture:exercise/exercise.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/teams',
      status: 200,
      response: 'fixture:team/list.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/user',
      status: 200,
      response: []
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('updates exercise', () => {
    cy.visit('/exercises/001/edit')
    cy.contains('Edit Exercise')
    cy.contains('div', 'Title').find('input').first().clear().type(exercise.title)
    cy.contains('edit').click()
    cy.contains(exercise.title)
    cy.contains(exercise.author.username)
    cy.url().should('include', '/exercises/001')
  })
})
