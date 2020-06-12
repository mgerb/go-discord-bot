import React from 'react';
import { SoundListType, SoundType } from '../../model';
import { SoundService } from '../../services';
import { ClipPlayerControl } from '../clip-player-control/clip-player-control';
import './sound-list.scss';

interface Props {
  soundList: SoundType[];
  type: SoundListType;
  title: string;
  onPlayDiscord?: (sound: SoundType) => void;
  hasModPermissions: boolean;
  onFavorite?: (sound: SoundType) => void;
  deleteFavorite?: (sound: SoundType) => void;
}

interface State {
  showAudioControls: boolean[];
}

export class SoundList extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      showAudioControls: [],
    };
  }

  checkExtension(extension: string) {
    switch (extension) {
      case 'wav':
        return true;
      case 'mp3':
        return true;
      case 'mpeg':
        return true;
      default:
        return false;
    }
  }

  render() {
    const { hasModPermissions, onFavorite, soundList, title, type } = this.props;

    return (
      <div className="card">
        <div className="card__header" style={{ display: 'flex' }}>
          <div>
            <span>{title}</span>
            <i className="fa fa fa-volume-up" aria-hidden="true" />
          </div>
          <div style={{ flex: 1 }} />
          <div>({soundList.length})</div>
        </div>

        {soundList.length > 0
          ? soundList.map((sound: SoundType, index: number) => {
              return (
                <div key={index} className="sound-list__item">
                  <div className="text-wrap">
                    {((type === 'sounds' || type === 'favorites') && sound.prefix ? sound.prefix : '') + sound.name}
                  </div>

                  <ClipPlayerControl
                    showFavorite={type !== 'clips'}
                    onFavorite={() => !onFavorite || onFavorite(sound)}
                    hasModPermissions={hasModPermissions}
                    sound={sound}
                    type={type}
                    onPlayBrowser={(sound) => SoundService.playAudioInBrowser(sound, type)}
                    onPlayDiscord={SoundService.playSound}
                  />
                </div>
              );
            })
          : null}
      </div>
    );
  }
}
