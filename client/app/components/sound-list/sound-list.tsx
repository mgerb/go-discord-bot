import React from 'react';
import { ClipPlayerControl } from '../clip-player-control/clip-player-control';
import './sound-list.scss';

interface Props {
  soundList: SoundType[];
  type: 'sounds' | 'clips';
  onPlayDiscord?: (sound: SoundType) => void;
  showDiscordPlay?: boolean;
}

interface State {
  showAudioControls: boolean[];
}

export interface SoundType {
  extension: string;
  name: string;
  prefix?: string;
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

  handlePlayAudioInBrowser(sound: SoundType, type: string) {
    const url = `/public/${type.toLowerCase()}/` + sound.name + '.' + sound.extension;
    const audio = new Audio(url);
    audio.play();
  }

  render() {
    const { showDiscordPlay, soundList, type } = this.props;

    return (
      <div className="card">
        <div className="card__header" style={{ display: 'flex' }}>
          <div>
            <span>{type}</span>
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
                    {(type === 'sounds' && sound.prefix ? sound.prefix : '') + sound.name}
                  </div>

                  <ClipPlayerControl showDiscordPlay={showDiscordPlay} sound={sound} type={type} />
                </div>
              );
            })
          : null}
      </div>
    );
  }
}
