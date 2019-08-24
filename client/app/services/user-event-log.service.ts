import { IUserEventLog } from '../model';
import { axios } from './axios.service';

export class UserEventLogService {
  public static async getUserEventLogs(): Promise<IUserEventLog[]> {
    const resp = await axios.get('/api/user-event-log');
    return resp.data.data;
  }
}
