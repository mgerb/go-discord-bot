import React from 'react';
import { SoundService } from '../../services';

interface IProps {
  sound: SoundType;
  type: 'sounds' | 'clips';
  showDiscordPlay?: boolean;
}

interface IState {}

export interface SoundType {
  extension: string;
  name: string;
}

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

  handlePlayAudioInBrowser(sound: SoundType, type: string) {
    const url = `/public/${type.toLowerCase()}/` + sound.name + '.' + sound.extension;
    const audio = new Audio(url);
    audio.play();
  }

  onPlayDiscord = (sound: SoundType) => {
    SoundService.playSound(sound);
  };

  render() {
    const { sound, showDiscordPlay, type } = this.props;

    return (
      this.checkExtension(sound.extension) && (
        <div className="flex flex--v-center">
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
            onClick={() => this.handlePlayAudioInBrowser(sound, type)}
          />
          {showDiscordPlay && (
            <i
              title="Play in discord"
              className="fa fa-play-circle link fa-lg"
              aria-hidden="true"
              style={{ paddingLeft: '10px' }}
              onClick={() => this.onPlayDiscord(sound)}
            />
          )}
        </div>
      )
    );
  }
}
