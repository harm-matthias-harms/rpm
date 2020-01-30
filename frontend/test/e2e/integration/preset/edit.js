/* eslint-disable no-undef */

describe('Edit preset', () => {
  let preset
  let vitalSigns
  beforeEach(() => {
    cy.fixture('preset/edit.json').then((json) => {
      preset = json
      vitalSigns = preset.vitalSigns
    })
    cy.server()
    cy.route({
      method: 'PUT',
      url: 'http://localhost:3001/api/presets/001',
      status: 200,
      response: 'fixture:preset/edit.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets/001',
      status: 200,
      response: 'fixture:preset/preset.json'
    })
    cy.login()
  })

  afterEach(() => {
    cy.logout()
  })

  it('updates preset', () => {
    cy.visit('/presets/001/edit')
    cy.contains('Edit Preset')
    cy.contains('div', 'Height').find('input').first().clear().type(vitalSigns.height)
    cy.contains('edit').click()
    cy.contains(preset.title)
    cy.contains(preset.author.username)
    cy.contains(preset.editor.username)
    cy.contains('1.95 m')
    cy.url().should('include', '/presets/001')
  })
})
