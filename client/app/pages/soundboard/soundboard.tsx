import { inject, observer } from 'mobx-react';
import React from 'react';
import { SoundList, SoundType, Uploader } from '../../components';
import { axios, SoundService } from '../../services';
import { AppStore } from '../../stores';
import './soundboard.scss';

interface Props {
  appStore?: AppStore;
}

interface State {
  percentCompleted: number;
  uploaded: boolean;
  uploadError: string;
  soundList: SoundType[];
}

@inject('appStore')
@observer
export class Soundboard extends React.Component<Props, State> {
  private soundListCache: any;

  constructor(props: Props) {
    super(props);
    this.state = {
      percentCompleted: 0,
      uploaded: false,
      uploadError: ' ',
      soundList: [],
    };
  }

  componentDidMount() {
    this.getSoundList();
  }

  private getSoundList() {
    if (!this.soundListCache) {
      axios
        .get('/api/soundlist')
        .then(response => {
          this.soundListCache = response.data;
          this.setState({
            soundList: response.data,
          });
        })
        .catch((error: any) => {
          console.error(error.response.data);
        });
    } else {
      this.setState({
        soundList: this.soundListCache,
      });
    }
  }

  onUploadComplete = () => {
    delete this.soundListCache;
    this.getSoundList();
  };

  onPlayDiscord = (sound: SoundType) => {
    SoundService.playSound(sound);
  };

  render() {
    const { soundList } = this.state;
    const { appStore } = this.props;
    return (
      <div className="content">
        <Uploader onComplete={this.onUploadComplete} />
        <SoundList
          soundList={soundList}
          type="sounds"
          onPlayDiscord={this.onPlayDiscord}
          showDiscordPlay={appStore!.hasModPermissions()}
        />
      </div>
    );
  }
}
