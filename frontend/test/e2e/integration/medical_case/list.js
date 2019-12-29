/* eslint-disable no-undef */

describe('List medical case', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/list.json',
    })
    cy.login()
  })
  afterEach(() => {
    cy.logout()
  })
  it('lists all medical cases', () => {
    cy.visit('/medical_cases')
    cy.contains('2')
    cy.contains('test medical case 1')
    cy.contains('John Doe')
    cy.contains('test medical case 2')
    cy.contains('otherUser')
  })
  it('lists my presets', () => {
    cy.visit('/medical_cases/my')
    cy.contains('1')
    cy.contains('test medical case 1')
    cy.contains('John Doe')
    cy.get('test medical case 2').should('not.exist')
    cy.contains('otherUser').should('not.exist')
  })
})
