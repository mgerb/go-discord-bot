import React from 'react';
import './sound-list.scss';

interface Props {
  soundList: SoundType[];
  type: string;
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

  handleShowAudio(index: any) {
    let temp = this.state.showAudioControls;
    temp[index] = true;

    this.setState({
      showAudioControls: temp,
    });
  }

  render() {
    const { soundList, type } = this.props;

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
                  <div className="text-wrap">{(sound.prefix || '') + sound.name}</div>

                  {this.checkExtension(sound.extension) && this.state.showAudioControls[index] ? (
                    <audio
                      controls
                      src={`/public/${type.toLowerCase()}/` + sound.name + '.' + sound.extension}
                      itemType={'audio/' + sound.extension}
                      style={{ width: '100px' }}
                    />
                  ) : (
                    <i className="fa fa-play link" aria-hidden="true" onClick={() => this.handleShowAudio(index)} />
                  )}
                </div>
              );
            })
          : null}
      </div>
    );
  }
}
