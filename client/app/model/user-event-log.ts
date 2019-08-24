import { IUser } from './user';

export interface IUserEventLog {
  content: string;
  created_at: string;
  deleted_at?: string;
  id: number;
  updated_at: string;
  user: IUser;
  user_id: string;
}
