import { IUser } from '../model';
import { axios } from './axios.service';

export class UserService {
  public static async getUsers(): Promise<IUser[]> {
    const resp = await axios.get('/api/user');
    return resp.data.data;
  }

  public static async putUsers(users: IUser[]): Promise<IUser[]> {
    const resp = await axios.put('/api/user', { users });
    return resp.data.data;
  }
}
