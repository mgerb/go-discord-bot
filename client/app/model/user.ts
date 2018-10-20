import { Permissions } from './permissions';

export interface IUser {
  avatar: string;
  bot: boolean;
  created_at: string;
  deleted_at?: string;
  discriminator: string;
  email: string;
  id: string;
  mfa_enabled: boolean;
  permissions: Permissions;
  token: string;
  updated_at: string;
  username: string;
  verified: boolean;
}
