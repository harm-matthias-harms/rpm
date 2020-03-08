/* eslint-disable no-undef */

describe('Create preset', () => {
  let preset
  let vitalSigns
  beforeEach(() => {
    cy.fixture('preset/create.json').then((json) => {
      preset = json.data
      vitalSigns = preset.vitalSigns
    })
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/api/presets',
      status: 200,
      response: 'fixture:preset/create.json'
    })
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

  it('creates preset', () => {
    cy.visit('/presets/new')
    cy.contains('New Preset')
    cy.contains('div', 'Title').find('input').first().type(preset.title)
    cy.contains('div', 'Onset of symptoms').find('input').first().type(vitalSigns.oos)
    cy.contains('div', 'AVPU').first('div[role="button"]').click()
    cy.contains('div', vitalSigns.avpu).first('div[role="option"]').click()
    cy.contains('div', 'Mobility').first('div[role="button"]').click()
    cy.contains('div', vitalSigns.mobility).first('div[role="option"]').click()
    cy.contains('div', 'Pulse').find('input').first().type(vitalSigns.pulse)
    cy.contains('div', 'Blood pressure systolic').find('input').first().type(vitalSigns.bloodPressureSystolic)
    cy.contains('div', 'Blood pressure diastolic').find('input').first().type(vitalSigns.bloodPressureDiastolic)
    cy.contains('div', 'Respiratory rate').find('input').first().type(vitalSigns.respiratoryRate)
    cy.contains('div', 'Oxygen saturation').find('input').first().type(vitalSigns.oxygenSaturation)
    cy.contains('div', 'Capillary refill').find('input').first().type(vitalSigns.capillaryRefill)
    cy.contains('div', 'Body temperature').find('input').first().type(vitalSigns.temperature)
    cy.contains('div', 'Findings on examination').find('textarea').first().type(vitalSigns.expectations.foe)
    cy.contains('div', 'Expected treatment').find('textarea').first().type(vitalSigns.expectations.treatmentExpected)
    cy.contains('create').click()
    cy.contains(preset.title)
    cy.contains(preset.author.username)
    cy.url().should('include', '/presets')
  })
})
