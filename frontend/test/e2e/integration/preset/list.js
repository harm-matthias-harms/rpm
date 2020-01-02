/* eslint-disable no-undef */

describe('List presets', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets',
      status: 200,
      response: 'fixture:preset/list.json',
    })
    cy.login()
  })
  afterEach(() => {
    cy.logout()
  })
  it('lists all presets', () => {
    cy.visit('/presets')
    cy.contains('2')
    cy.contains('test preset 1')
    cy.contains('John Doe')
    cy.contains('test preset 2')
    cy.contains('otherUser')
  })
  it('lists my presets', () => {
    cy.visit('/presets/my')
    cy.contains('1')
    cy.contains('test preset 1')
    cy.contains('John Doe')
    cy.get('test preset 2').should('not.exist')
    cy.contains('otherUser').should('not.exist')
  })
})
