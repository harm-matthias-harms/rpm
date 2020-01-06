/* eslint-disable no-undef */

describe('Create preset', () => {
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
        response: 'fixture:preset/edit.json',
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

    it('creates preset', () => {
      cy.visit('/presets')
      cy.contains('edit')
      cy.visit('/presets/001')
      cy.contains('edit').click()
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
