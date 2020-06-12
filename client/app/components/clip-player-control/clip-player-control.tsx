import React from 'react';
import { SoundListType, SoundType } from '../../model';

interface IProps {
  sound: SoundType;
  type: SoundListType;
  hasModPermissions: boolean;
  showFavorite?: boolean;
  onFavorite?: () => void;
  onPlayBrowser: (sound: SoundType) => void;
  onPlayDiscord: (sound: SoundType) => void;
}

interface IState {}

export class ClipPlayerControl extends React.Component<IProps, IState> {
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
    const { onPlayBrowser, onPlayDiscord, sound, hasModPermissions, showFavorite, type } = this.props;

    return (
      this.checkExtension(sound.extension) && (
        <div className="flex flex--center">
          {showFavorite && hasModPermissions && (
            <i
              title="Favorite"
              className={'fa link fa-lg ' + (type === 'favorites' ? 'fa-trash' : 'fa-heart color__red')}
              aria-hidden="true"
              style={{ paddingRight: '5px' }}
              onClick={() => !this.props.onFavorite || this.props.onFavorite()}
            />
          )}
          <a
            href={`/public/${type.toLowerCase()}/` + sound.name + '.' + sound.extension}
            download
            title="Download"
            className="fa fa-download link"
            aria-hidden="true"
          />
          <i
            title="Play in browser"
            className="fa fa-play link"
            aria-hidden="true"
            style={{ paddingLeft: '15px' }}
            onClick={() => onPlayBrowser(sound)}
          />
          {hasModPermissions && (
            <i
              title="Play in discord"
              className="fa fa-play-circle link fa-lg"
              aria-hidden="true"
              style={{ paddingLeft: '10px' }}
              onClick={() => onPlayDiscord(sound)}
            />
          )}
        </div>
      )
    );
  }
}
