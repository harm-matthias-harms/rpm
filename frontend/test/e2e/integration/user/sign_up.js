/* eslint-disable no-undef */

describe('SignUp', () => {
  it('has link to sign in', () => {
    cy.visit('/sign_up')
    cy.contains('Sign in here').click()
    cy.url().should('include', '/')
  })
  it('error on no connection', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/register',
      response: {},
      status: 404
    })
    cy.register()
    cy.contains("Couldn't create new account.")
  })
  it('create account', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/register',
      response: { success: true }
    })
    cy.register()
    cy.get('div > .v-icon')
    cy.contains('a', 'Sign in here')
  })
  it('recognises if account already exists', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:3001/auth/register',
      response: {
        message: 'user already exists'
      },
      status: 400
    })
    cy.register()
    cy.url().should('include', '/')
  })
})
