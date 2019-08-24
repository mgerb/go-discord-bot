import * as _ from 'lodash';
import { DateTime } from 'luxon';
import React from 'react';
import { ISound } from '../../model';
import { ClipPlayerControl } from '../clip-player-control/clip-player-control';

interface IProps {
  sounds: ISound[];
  showDiscordPlay?: boolean;
}

export const UploadHistory = ({ sounds, showDiscordPlay }: IProps) => {
  const sortedSounds = _.orderBy(sounds, 'created_at', 'desc');
  return (
    <div className="card">
      <div className="card__header">Upload History</div>
      <table className="table table--ellipsis">
        <thead>
          <tr>
            <th className="hide-small">Date</th>
            <th>Sound</th>
            <th className="hide-small">Ext</th>
            <th>User</th>
            <th className="hide-small">Email</th>
          </tr>
        </thead>
        <tbody>
          {sortedSounds.map((s: ISound, i) => {
            const formattedDate = DateTime.fromISO(s.created_at).toLocaleString();
            return (
              <tr key={i}>
                <td className="hide-small" title={formattedDate}>
                  {formattedDate}
                </td>
                <td title={s.name}>{s.name}</td>
                <td className="hide-small" title={s.extension}>
                  {s.extension}
                </td>
                <td title={s.user.username}>{s.user.username}</td>
                <td className="hide-small" title={s.user.email}>
                  {s.user.email}
                </td>
                <td>
                  <ClipPlayerControl showDiscordPlay={showDiscordPlay} sound={s} type="sounds"></ClipPlayerControl>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};
