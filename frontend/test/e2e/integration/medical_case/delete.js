/* eslint-disable no-undef */

describe('Delete Medical Case', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'DELETE',
      url: 'http://localhost:3001/api/medical_cases/001',
      status: 200,
      response: { status: 200 }
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/list.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases/001',
      status: 200,
      response: 'fixture:medical_case/medicalCase.json'
    })
    cy.login()
  })
  afterEach(() => {
    cy.logout()
  })

  /*
    until cypress fixes the icon bug
    it('deletes medical case', () => {
    cy.visit('/medical_cases')
    cy.get('i').contains('delete')
    cy.visit('/medical_cases/001')
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this medical case?')
    cy.get('.v-card__actions').contains('cancel').click()
    cy.get('i').contains('delete').click()
    cy.contains('Are you sure you want to delete this medical case?')
    cy.get('.v-card__actions').contains('ok').click()
    cy.url().should('include', '/medical_cases')
  }) */
})
