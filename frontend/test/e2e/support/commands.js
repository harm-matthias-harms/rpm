// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This is will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })
Cypress.Commands.add('register', () => {
  cy.visit('/')
  cy.contains('sign up').click()
  // sign up button should be disables
  cy.get('.v-form > .v-btn').should('be.disabled')
  // enter username
  cy.get('input[name=username]').type('user')
  cy.contains('The username field must be at least 6 characters.')
  cy.get('input[name=username]').type('name')
  // enter email
  cy.get('input[name=email]').type('user')
  cy.contains('The email field must be a valid email.')
  cy.get('input[name=email]').type('@mail.com')
  // enter password
  cy.get('input[name=password]').type('123')
  cy.contains('The password field must be at least 6 characters.')
  cy.get('input[name=password]').type('456')
  cy.get('input[name=password_confirm]').type('123')
  cy.contains('The password confirmation does not match.')
  cy.get('input[name=password_confirm]').type('456')
  // submit form
  cy.get('form').submit()
})
