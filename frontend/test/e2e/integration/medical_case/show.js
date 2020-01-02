/* eslint-disable no-undef */

describe('Show a medical case', () => {
  beforeEach(() => {
    cy.server()
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases',
      status: 200,
      response: 'fixture:medical_case/list.json',
    })
    cy.route({
      method: 'GET',
      url: 'http://localhost:3001/api/medical_cases/001',
      status: 200,
      response: 'fixture:medical_case/medicalCase.json',
    })
    cy.login()
  })
  afterEach(() => {
    cy.logout()
  })
  it('shows a medical Case', () => {
    cy.visit('/medical_cases')
    cy.contains('Test Medical Case').click()
    // has title
    cy.contains('Test Medical Case')
    // has author
    cy.contains('John Doe')
    cy.contains("Jan 01'70")
    // has editor
    cy.contains('Sam Who')
    cy.contains("Jan 01'01")
    // contains general Information
    cy.contains('General Information')
    cy.contains('USAR')
    cy.contains('MEDIVAC')
    cy.contains('Surgical')
    cy.contains('Need for Hospilisation')
    cy.contains('Short Summary:')
    cy.contains('short summary')
    cy.contains('Triage:')
    cy.contains('Urgent')
    cy.contains('Age:')
    cy.contains('0-10')
    cy.contains('Gender:')
    cy.contains('male')
    // contains medical history
    cy.contains('Medical History')
    cy.contains('Problems/Conditions:')
    cy.contains('problems')
    cy.contains('Vaccinations:').should('not.exist')
    cy.contains('Allergies:')
    cy.contains('allergies')
    cy.contains('Medication:')
    cy.contains('medication')
    cy.contains('Implantable Devices:')
    cy.contains('implanted devices')
    // contains expectations
    cy.contains('Expectations')
    cy.contains('General Status:')
    cy.contains('general status')
    cy.contains('On Examination:')
    cy.contains('on examination')
    cy.contains('Expectations:')
    cy.contains('expectations')
    // contains make-up etc.
    cy.contains('Other Information:')
    cy.contains('other information')
    cy.contains('Needed Make-Up and Attributes:')
    cy.contains('makeup')
    // contains files
    cy.contains('Protokoll 1.11.2018 M-Lab.pdf')
    cy.contains('237.91 KB')
    // contains vital signs
    cy.contains('starting treatment')
    cy.contains('ending treatment')
    cy.contains('successful treatment')
    cy.contains('Reason:')
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
    cy.contains('reason')
    cy.contains('2 hours')
    cy.contains('Unresponsive')
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