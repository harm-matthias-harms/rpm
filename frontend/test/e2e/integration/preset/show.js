/* eslint-disable no-undef */

describe('View a preset', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/presets',
      status: 200,
      response: 'fixture:preset/list.json'
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
  it('views a preset', () => {
    cy.visit('/presets')
    cy.contains('test preset 1').click()
    // has title
    cy.contains('Test preset')
    // has author
    cy.contains('John Doe')
    cy.contains("Jan 01'70")
    // contains prefix
    cy.contains('Onset of Symptoms:')
    cy.contains('AVPU:')
    cy.contains('Mobility:')
    cy.contains('Pulse:')
    cy.contains('Blood Pressure:')
    cy.contains('Respiratory Rate:')
    cy.contains('Oxygen Saturation:')
    cy.contains('Capillary Refill:')
    cy.contains('Temperatur:')
    cy.contains('Weight:')
    cy.contains('Height:')
    // contains values
    cy.contains('2 hours')
    cy.contains('unconscious')
    cy.contains('none')
    cy.contains('50 bpm')
    cy.contains('120/80')
    cy.contains('12 bpm')
    cy.contains('98 %')
    cy.contains('1.5 s')
    cy.contains('36.3 °C')
    cy.contains('90 kg')
    cy.contains('1.9 m')
  })
})
