/* eslint-disable no-undef */

describe('List medical case', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/list.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('lists all medical cases', () => {
    cy.visit('/medical_cases')
    cy.contains('2')
    cy.contains('Test Medical Case')
    cy.contains('John Doe')
    cy.contains('test medical case 2')
    cy.contains('otherUser')
  })

  it('filters medical cases', () => {
    cy.visit('/medical_cases')
    cy.contains('div', 'Search')
      .find('input')
      .first()
      .type('John')
    cy.contains('John Doe')
    cy.contains('otherUser').should('not.exist')
  })
})
