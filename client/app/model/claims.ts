import { Permissions } from './permissions';

// JWT claims
export interface IClaims {
  id: string;
  username: string;
  email: string;
  discriminator: string;
  permissions: Permissions;
  exp: number;
  iss: string; // issuer
}
