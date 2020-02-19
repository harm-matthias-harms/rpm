/* eslint-disable no-undef */

describe('Create medical case', () => {
  let medicalCase
  let generalInformation
  let medicalHistory
  let expectations
  let vitalSigns
  let child
  beforeEach(() => {
    cy.fixture('medical_case/create.json').then((json) => {
      medicalCase = json.data
      generalInformation = medicalCase.generalInformation
      medicalHistory = medicalCase.medicalHistory
      expectations = medicalCase.expectations
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
    cy.contains('div', 'make-up')
      .find('textarea')
      .first()
      .type(medicalCase.makeup)
    cy.contains('div', 'Other information')
      .find('textarea')
      .first()
      .type(medicalCase.otherInformation)

    // general information
    cy.contains('Surgical').click()
    cy.contains('Need for hospilisation').click()
    cy.contains('USAR').click()
    cy.contains('MEDIVAC').click()

    cy.contains('div', 'Triage')
      .first('div[role="button"]')
      .click()
    cy.contains('div', generalInformation.triage)
      .first('div[role="option"]')
      .click()

    cy.contains('div', 'Short summary')
      .find('textarea')
      .first()
      .type(generalInformation.shortSummary)

    cy.contains('div', 'Age')
      .first('div[role="button"]')
      .click()
    cy.contains('div', generalInformation.age)
      .first('div[role="option"]')
      .click()

    cy.contains('div', 'Gender')
      .first('div[role="button"]')
      .click()
    cy.contains('div', generalInformation.gender)
      .first('div[role="option"]')
      .click()

    // medical history
    cy.contains('Medical history').click()

    cy.contains('div', 'Problems')
      .find('input')
      .first()
      .type(medicalHistory.problems)

    cy.contains('div', 'Vaccinations')
      .find('input')
      .first()
      .type(medicalHistory.vaccinations)

    cy.contains('div', 'Allergies')
      .find('input')
      .first()
      .type(medicalHistory.allergies)

    cy.contains('div', 'Medication')
      .find('input')
      .first()
      .type(medicalHistory.medication)

    cy.contains('div', 'Implantable devices')
      .find('input')
      .first()
      .type(medicalHistory.implantedDevices)

    // expectations
    cy.contains('div', 'General status')
      .find('textarea')
      .first()
      .type(expectations.generalStatus)

    cy.contains('div', 'On examination')
      .find('textarea')
      .first()
      .type(expectations.onExamination)

    cy.contains('div', 'Expectations')
      .find('textarea')
      .first()
      .type(expectations.expectations)

    // Vital Signs

    // add first vital sign
    cy.contains('button', 'Vital Signs')
      .find('.fa-plus')
      .first()
      .click()

    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Title')
      .find('input')
      .last()
      .type(vitalSigns.title)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Reason')
      .find('input')
      .last()
      .type(vitalSigns.reason)
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
      .find('input')
      .first()
      .type(vitalSigns.data.mobility)
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
      .contains('div', 'Temperature')
      .find('input')
      .first()
      .type(vitalSigns.data.temperature)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Weight')
      .find('input')
      .first()
      .type(vitalSigns.data.weight)
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Height')
      .find('input')
      .first()
      .type(vitalSigns.data.height - 10)

    // select predefined preset
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Select a preset')
      .find('input')
      .first()
      .type('test preset 1')

    cy.contains('div', 'test preset 1').click()

    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Height')
      .find('input')
      .first()
      .should('have.value', vitalSigns.data.height.toString())

    // add sub vital sign
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('Add Step')
      .click()
    cy.get(':nth-child(1) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('No title set')
      .click()

    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Title')
      .find('input')
      .last()
      .type(child.title)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Reason')
      .find('input')
      .last()
      .type(child.reason)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Onset of symptoms')
      .find('input')
      .first()
      .type(child.data.oos)
    // currently not working because cypress is not finding the right document
    // cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
    //   .contains('div', 'AVPU')
    //   .first('div[role="button"]')
    //   .click()
    // cy.contains('div', child.data.avpu)
    //   .last('div[role="option"]')
    //   .click()
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Mobility')
      .find('input')
      .first()
      .type(child.data.mobility)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Pulse')
      .find('input')
      .first()
      .type(child.data.pulse)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure systolic')
      .find('input')
      .first()
      .type(child.data.bloodPressureSystolic)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Blood pressure diastolic')
      .find('input')
      .first()
      .type(child.data.bloodPressureDiastolic)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Respiratory rate')
      .find('input')
      .first()
      .type(child.data.respiratoryRate)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Oxygen saturation')
      .find('input')
      .first()
      .type(child.data.oxygenSaturation)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Capillary refill')
      .find('input')
      .first()
      .type(child.data.capillaryRefill)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Temperature')
      .find('input')
      .first()
      .type(child.data.temperature)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Weight')
      .find('input')
      .first()
      .type(child.data.weight)
    cy.get(':nth-child(6) > .v-expansion-panel > .v-expansion-panel-content')
      .contains('div', 'Height')
      .find('input')
      .first()
      .type(child.data.height)

    cy.contains('create').click()
    cy.contains(medicalCase.title)
    cy.contains(medicalCase.author.username)
    cy.url().should('include', '/medical_cases')
  })
})
