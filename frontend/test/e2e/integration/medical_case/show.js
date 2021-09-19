/* eslint-disable no-undef */

describe('Show a medical case', () => {
  beforeEach(() => {
    cy.server()
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
    // contains general
    cy.contains('General')
    cy.contains('Area/Discipline')
    cy.contains('Surgery')
    cy.contains('Context')
    cy.contains('Europe, LAMIC')
    cy.contains('Scenario')
    cy.contains('Conflict, Outbreak')
    // contains patient
    cy.contains("Patient's characteristics")
    cy.contains('Patient type')
    cy.contains('Acute')
    cy.contains('Triage')
    cy.contains('Red')
    cy.contains('Gender')
    cy.contains('Male')
    cy.contains('Age')
    cy.contains('Adult')
    // contains medical
    cy.contains('Medical assessment')
    cy.contains('S - Signs/Symptoms')
    cy.contains('signs')
    cy.contains('A - Allergies')
    cy.contains('allergies')
    cy.contains('M - Medications')
    cy.contains('medication')
    cy.contains('P - Past pertinent medical history')
    cy.contains('past medical history')
    cy.contains('L - Last oral intake')
    cy.contains('loi')
    cy.contains('E - Events leading up to present illness/injury')
    cy.contains('events')
    // contains makeup
    cy.contains('Make-up and attributes')
    cy.contains('Make-up instructions')
    cy.contains('make-up instructions')
    cy.contains('Acting instructions')
    cy.contains('acting instructions')
    // contains files
    cy.contains('Protokoll 1.11.2018 M-Lab.pdf')
    cy.contains('237.91 KB')
    // contains vital signs
    cy.contains('T0 - Start')
    cy.contains('T1 - Improvement')
    cy.contains('Onset of symptoms')
    cy.contains('AVPU')
    cy.contains('Mobility')
    cy.contains('Pulse')
    cy.contains('Blood pressure')
    cy.contains('Respiratory rate')
    cy.contains('Oxygen saturation')
    cy.contains('Capillary refill')
    cy.contains('Body temperatur')
    cy.contains('Findings on examination')
    cy.contains('Expected treatment')
    cy.contains('2 hours')
    cy.contains('U - Unresponsive')
    cy.contains('Walking')
    cy.contains('50 bpm')
    cy.contains('120/80')
    cy.contains('12 bpm')
    cy.contains('98 %')
    cy.contains('1.5 s')
    cy.contains('36.3 °C')
    cy.contains('foe')
    cy.contains('expected treatment')
    // contains default exams
    cy.contains('Normal exams')
  })
})
