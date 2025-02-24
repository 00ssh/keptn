/// <reference types="cypress" />

class SettingsPage {
  GIT_USER_LOC = 'input[formcontrolname="gitUser"]';
  GIT_URL_INPUT_LOC = 'input[formcontrolname="gitUrl"]';

  inputGitUrl(GIT_URL: string): this {
    cy.get('input[formcontrolname="gitUrl"]').type(GIT_URL);
    return this;
  }

  inputGitUsername(GIT_USERNAME: string): this {
    cy.get('input[formcontrolname="gitUser"]').type(GIT_USERNAME);
    return this;
  }

  inputGitToken(GIT_TOKEN: string): this {
    cy.get('input[formcontrolname="gitToken"]').type(GIT_TOKEN);
    return this;
  }

  clickSaveChanges(): this {
    cy.get('.dt-button-primary > span.dt-button-label').contains('Save changes').click();
    return this;
  }

  getErrorMessageText(): unknown {
    return cy.get('.small');
  }

  clickDeleteProjectButton(): this {
    cy.get('span.dt-button-label').contains('Delete this project').click();
    return this;
  }

  typeProjectNameToDelete(projectName: string): this {
    const projectInputLoc = 'input[placeholder=proj_pattern]';
    cy.get(projectInputLoc.replace('proj_pattern', projectName)).click().type(projectName);
    return this;
  }

  submitDelete(): void {
    cy.get('span.dt-button-label').contains('I understand the consequences, delete this project').click();
  }
}

export default SettingsPage;
