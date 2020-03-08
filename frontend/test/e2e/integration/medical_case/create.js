/* eslint-disable no-undef */

describe('Create medical case', () => {
  let medicalCase
  let general
  let patient
  let medical
  let makeup
  let vitalSigns
  let child
  beforeEach(() => {
    cy.fixture('medical_case/create.json').then((json) => {
      medicalCase = json.data
      general = medicalCase.general
      patient = medicalCase.patient
      medical = medicalCase.medical
      makeup = medicalCase.makeup
      vitalSigns = medicalCase.vitalSigns[0]
      child = vitalSigns.childs[0]
    })
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/create.json'
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/list.json'
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

  it('creates medical case', () => {
    cy.visit('/medical_cases/new')
    cy.contains('New Medical Case')
    // general Information
    cy.contains('div', 'Title')
      .find('input')
      .first()
      .type(medicalCase.title)

    // general
    cy.contains('div', 'Area/Discipline')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', general.discipline)
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Context')
      .first('div[role="button"]')
      .click({ force: true })
    general.context.forEach((e) => {
      cy.contains('div', e)
        .last('div[role="option"]')
        .click()
    })

    cy.contains('div', 'Scenario')
      .first('div[role="button"]')
      .click({ force: true })
    general.scenario.forEach((e) => {
      cy.contains('div', e)
        .last('div[role="option"]')
        .click()
    })

    // patient
    cy.contains('div', 'Patient type')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', patient.type)
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Triage')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', patient.triage)
      .last('div[role="option"]')
      .click()

    cy.contains('div', 'Gender')
      .first('div[role="button"]')
      .click({ force: true })
    patient.gender.forEach((e) => {
      cy.contains('div', e)
        .last('div[role="option"]')
        .click()
    })

    cy.contains('div', 'Age')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', patient.age)
      .last('div[role="option"]')
      .click()

    // medical
    cy.contains('div', 'S – Signs/Symptoms')
      .find('input')
      .first()
      .type(medical.signs, { force: true })

    cy.contains('div', 'A – Allergies')
      .find('input')
      .first()
      .type(medical.allergies, { force: true })

    cy.contains('div', 'M – Medications')
      .find('input')
      .first()
      .type(medical.medication, { force: true })

    cy.contains('div', 'P – Past pertinent medical history')
      .find('input')
      .first()
      .type(medical.past, { force: true })

    cy.contains('div', 'L – Last oral intake')
      .find('input')
      .first()
      .type(medical.loi, { force: true })

    cy.contains('div', 'E – Events leading up to present illness/injury')
      .find('input')
      .first()
      .type(medical.events, { force: true })

    // makeup
    cy.contains('div', 'Make-up instructions')
      .find('textarea')
      .first()
      .type(makeup.makeup, { force: true })

    cy.contains('div', 'Acting instructions')
      .find('textarea')
      .first()
      .type(makeup.acting, { force: true })

    // Vital Signs

    // add first vital sign
    cy.contains('button', 'Vital Signs')
      .find('.fa-plus')
      .first()
      .click()

    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Title')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', vitalSigns.title)
      .first('div[role="option"]')
      .click()
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Onset of symptoms')
      .find('input')
      .first()
      .type(vitalSigns.data.oos)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'AVPU')
      .first('div[role="button"]')
      .click()
    cy.contains('div', vitalSigns.data.avpu)
      .first('div[role="option"]')
      .click()
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Mobility')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', vitalSigns.data.mobility)
      .first('div[role="option"]')
      .click()
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Pulse')
      .find('input')
      .first()
      .type(vitalSigns.data.pulse)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure systolic')
      .find('input')
      .first()
      .type(vitalSigns.data.bloodPressureSystolic)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure diastolic')
      .find('input')
      .first()
      .type(vitalSigns.data.bloodPressureDiastolic)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Respiratory rate')
      .find('input')
      .first()
      .type(vitalSigns.data.respiratoryRate)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Oxygen saturation')
      .find('input')
      .first()
      .type(vitalSigns.data.oxygenSaturation)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Capillary refill')
      .find('input')
      .first()
      .type(vitalSigns.data.capillaryRefill)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Body temperature')
      .find('input')
      .first()
      .type(vitalSigns.data.temperature)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Findings on examination')
      .find('textarea')
      .first()
      .type(vitalSigns.data.expectations.foe)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Expected treatment')
      .find('textarea')
      .first()
      .type(vitalSigns.data.expectations.treatmentExpected)

    // select predefined preset
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Select a preset')
      .find('input')
      .first()
      .type('test preset 1')

    cy.contains('div', 'test preset 1').click()

    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Body temperature')
      .find('input')
      .first()
      .should('have.value', vitalSigns.data.temperature.toString())

    // add sub vital sign
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('Add Step')
      .click()
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('No title set')
      .click()

    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Title')
      .first('div[role="button"]')
      .click({ force: true })
    cy.contains('div', child.title)
      .first('div[role="option"]')
      .click({ force: true })

    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Onset of symptoms')
      .find('input')
      .first()
      .type(child.data.oos)
    // currently not working because cypress is not finding the right document
    // cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
    //   .contains('div', 'AVPU')
    //   .first('div[role="button"]')
    //   .click({force: true})
    // cy.contains('div', child.data.avpu)
    //   .last('div[role="option"]')
    //   .click({force: true})
    // cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
    //   .contains('div', 'Mobility')
    //   .first('div[role="button"]')
    //   .click({force: true})
    // cy.contains('div', child.data.mobility)
    //   .first('div[role="option"]')
    //   .click({force: true})

    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Pulse')
      .find('input')
      .first()
      .type(child.data.pulse)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure systolic')
      .find('input')
      .first()
      .type(child.data.bloodPressureSystolic)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure diastolic')
      .find('input')
      .first()
      .type(child.data.bloodPressureDiastolic)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Respiratory rate')
      .find('input')
      .first()
      .type(child.data.respiratoryRate)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Oxygen saturation')
      .find('input')
      .first()
      .type(child.data.oxygenSaturation)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Capillary refill')
      .find('input')
      .first()
      .type(child.data.capillaryRefill)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Body temperature')
      .find('input')
      .first()
      .type(child.data.temperature)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Findings on examination')
      .find('textarea')
      .first()
      .type(child.data.expectations.foe)
    cy.get(':nth-child(5) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Expected treatment')
      .find('textarea')
      .first()
      .type(child.data.expectations.treatmentExpected)

    cy.contains('create').click()
    cy.contains(medicalCase.title)
    cy.contains(medicalCase.author.username)
    cy.url().should('include', '/medical_cases')
  })
})
