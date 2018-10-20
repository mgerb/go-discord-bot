import * as _ from 'lodash';
import { DateTime } from 'luxon';
import React from 'react';
import { ISound } from '../../model';

interface IProps {
  sounds: ISound[];
}

export const UploadHistory = ({ sounds }: IProps) => {
  const sortedSounds = _.orderBy(sounds, 'created_at', 'desc');
  return (
    <div className="card">
      <div className="card__header">Upload History</div>
      <table className="table">
        <thead>
          <tr>
            <th>Date</th>
            <th>Sound</th>
            <th className="hide-tiny">Ext</th>
            <th>Username</th>
            <th className="hide-tiny">Email</th>
          </tr>
        </thead>
        <tbody>
          {sortedSounds.map((s: ISound, i) => {
            const formattedDate = DateTime.fromISO(s.created_at).toLocaleString();
            return (
              <tr key={i}>
                <td title={formattedDate}>{formattedDate}</td>
                <td title={s.name}>{s.name}</td>
                <td className="hide-tiny" title={s.extension}>
                  {s.extension}
                </td>
                <td title={s.user.username}>{s.user.username}</td>
                <td className="hide-tiny" title={s.user.email}>
                  {s.user.email}
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};
