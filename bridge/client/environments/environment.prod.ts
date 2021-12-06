import { DynamicEnvironment } from './environment.dynamic';

class Environment extends DynamicEnvironment {
  constructor() {
    super();
    this.production = true;
    this.appConfigUrl = 'assets/branding/app-config.json';
    this.baseUrl = '/bridge';
  }
}

export const environment = new Environment();
