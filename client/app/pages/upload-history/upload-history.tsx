import orderBy from 'lodash/orderBy';
import { DateTime } from 'luxon';
import { inject, observer } from 'mobx-react';
import React from 'react';
import { ClipPlayerControl } from '../../components/clip-player-control/clip-player-control';
import { ISound } from '../../model';
import { SoundService } from '../../services';
import { AppStore } from '../../stores';

interface IProps {
  appStore: AppStore;
}

interface IState {
  sounds: ISound[];
}

@inject('appStore')
@observer
export class UploadHistory extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      sounds: [],
    };
  }

  componentDidMount() {
    SoundService.getSounds().then(sounds => {
      this.setState({ sounds });
    });
  }

  renderUploadHistory = (sounds: ISound[], showDiscordPlay: boolean) => {
    const sortedSounds = orderBy(sounds, 'created_at', 'desc');
    return (
      <div className="card">
        <div className="card__header">Upload History</div>
        <table className="table table--ellipsis table--fixed">
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
  public render() {
    const { hasModPermissions } = this.props.appStore;
    const { sounds } = this.state;
    return <div className="content">{this.renderUploadHistory(sounds, hasModPermissions())}</div>;
  }
}
