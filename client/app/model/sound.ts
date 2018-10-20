import { IUser } from './user';

export interface ISound {
  created_at: string;
  deleted_at?: string;
  extension: string;
  id: number;
  name: string;
  updated_at: string;
  user: IUser;
  user_id: string;
}
