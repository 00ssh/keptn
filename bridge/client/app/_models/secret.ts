import { Secret as scrt, SecretKeyValuePair } from '../../../shared/interfaces/secret';
import { SecretScope } from '../../../shared/interfaces/secret-scope';

export class Secret implements scrt {
  name!: string;
  scope!: SecretScope;
  keys?: string[];
  data?: SecretKeyValuePair[];

  constructor() {
    this.scope = SecretScope.DEFAULT;
    this.data = [];
  }

  static fromJSON(data: unknown): Secret {
    return Object.assign(new this(), data);
  }

  setName(name: string): void {
    this.name = name;
  }

  setScope(scope: SecretScope): void {
    this.scope = scope;
  }

  addData(key: string, value: string): void {
    this.data?.push({ key, value });
  }

  getData(index: number): SecretKeyValuePair {
    if (!this.data) {
      this.data = [];
    }

    return this.data[index];
  }

  removeData(index: number): void {
    this.data?.splice(index, 1);
  }
}
