/* eslint-disable no-undef */

describe('Edit Medical Case', () => {
  let medicalCase
  beforeEach(() => {
    cy.fixture('medical_case/edit.json').then((json) => {
      medicalCase = json
    })
    cy.server()
    cy.route({
      method: 'PUT',
      url: 'http://localhost:3001/api/medical_cases/001',
      status: 200,
      response: 'fixture:medical_case/edit.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases/001',
      status: 200,
      response: 'fixture:medical_case/medicalCase.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets',
      status: 200,
      response: 'fixture:preset/list.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('updates medical case', () => {
    cy.visit('/medical_cases/001/edit')
    cy.contains('Edit Medical Case')
    cy.contains('div', 'Signs')
      .find('input')
      .first()
      .clear()
      .type(medicalCase.medical.signs)
    cy.contains('edit').click()
    cy.contains(medicalCase.title)
    cy.contains(medicalCase.author.username)
    cy.contains(medicalCase.editor.username)
    cy.contains(medicalCase.medical.signs)
    cy.url().should('include', '/medical_cases/001')
  })
})
