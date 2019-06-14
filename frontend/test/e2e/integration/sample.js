/* eslint-disable */

describe('Sample tests', () => {
    it('Visits index page', () => {
      cy.visit('/');
      cy.contains('.v-card__title', 'Nuxt.js');
    });
  });