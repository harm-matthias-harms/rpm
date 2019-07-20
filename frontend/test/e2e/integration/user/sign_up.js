describe('SignUp', () => {
  it('has link to sign in', () => {
    cy.visit('/sign_up')
    cy.contains('Sign In here').click()
    cy.url().should('include', '/sign_in')
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
    cy.url().should('include', '/account_created')
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
    cy.url().should('include', '/sign_in')
  })
})
