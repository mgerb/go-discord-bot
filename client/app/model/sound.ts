import { IUser } from './user';

export type SoundListType = 'sounds' | 'clips' | 'favorites';

export interface SoundType {
  extension: string;
  name: string;
  prefix?: string;
}

// sound from database
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
