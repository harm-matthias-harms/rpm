/* eslint-disable no-undef */

describe('Create exercise', () => {
  let exercise
  beforeEach(() => {
    cy.fixture('exercise/create.json').then((json) => {
      exercise = json.data
    })
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/api/exercises',
      status: 200,
      response: 'fixture:exercise/create.json'
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

  it('creates exercise', () => {
    cy.visit('/exercises/new')
    cy.contains('New Exercise')
    cy.contains('div', 'Title').find('input').first().type(exercise.title)
    cy.get('.v-btn .v-icon.fa-plus').click({ multiple: true })

    cy.contains('div', 'Select a team')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', exercise.teams[0].team.title)
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Select a trainer')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', 'Autogenerate account')
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Select a role play manager')
      .first('div[role="button"]')
      .click({ force: true })
    cy.get('div:contains(Autogenerate account)')
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Name of make-up center')
      .find('input')
      .first()
      .type(exercise.makeupCenter[0].title)
    cy.contains('div', 'Select an account')
      .first('div[role="button"]')
      .click({ force: true })
    cy.get('div:contains(Autogenerate account)')
      .last('div[role="option"]')
      .click()

    cy.contains('create')// .click() cause of datepicker not tested
    // cy.contains(exercise.title)
    // cy.contains(exercise.author.username)
    // cy.url().should('include', '/exercises')
  })
})
