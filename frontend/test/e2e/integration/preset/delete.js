/* eslint-disable no-undef */

describe('Delete preset', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'DELETE',
      url: 'http://localhost:3001/api/presets/001',
      status: 200,
      response: { status: 200 },
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets',
      status: 200,
      response: 'fixture:preset/list.json',
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets/001',
      status: 200,
      response: 'fixture:preset/preset.json',
    })
    cy.login()
  })
  afterEach(() => {
    cy.logout()
  })

  it('updates preset', () => {
    cy.visit('/presets')
    cy.get('.v-icon').contains('delete')
    cy.visit('/presets/001')
    cy.get('.v-icon').contains('delete').click()
    cy.contains('Are you sure you want to delete this preset?')
    cy.get('.v-card__actions').contains('cancel').click()
    cy.get('.v-icon').contains('delete').click()
    cy.contains('Are you sure you want to delete this preset?')
    cy.get('.v-card__actions').contains('ok').click()
    cy.url().should('include', '/presets')
  })
})
