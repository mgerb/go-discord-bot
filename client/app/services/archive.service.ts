import * as _ from 'lodash';
import { IVideoArchive } from '../model';
import { axios } from './axios.service';

export class ArchiveService {
  public static async getVideoArchive(): Promise<IVideoArchive[]> {
    const data = (await axios.get('/api/video-archive')).data.data;
    return _.orderBy(data, 'created_at', ['desc']);
  }

  public static postVideoArchive(data: any): Promise<any> {
    return axios.post('/api/video-archive', data);
  }
}
