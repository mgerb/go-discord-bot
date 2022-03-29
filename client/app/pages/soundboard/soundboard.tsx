import { inject, observer } from 'mobx-react';
import React from 'react';
import { SoundList, Uploader } from '../../components';
import { SoundType } from '../../model';
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
  private soundListCache: SoundType[];

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
        .then((response) => {
          this.soundListCache = response.data;
          this.setState({
            soundList: response.data,
          });
        })
        // eslint-disable-next-line
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

  onFavorite = (sound: SoundType) => {
    this.props.appStore?.addFavorite(sound);
  };

  onDeleteFavorite = (sound: SoundType) => {
    this.props.appStore?.removeFavorite(sound);
  };

  render() {
    const { soundList } = this.state;

    if (!this.props.appStore) {
      return null;
    }
    const { hasModPermissions, getFavorites } = this.props.appStore;

    return (
      <div className="content">
        <Uploader onComplete={this.onUploadComplete} />
        {((hasModPermissions && getFavorites().length) || 0 > 0) && (
          <SoundList
            soundList={getFavorites()}
            title="Favorites"
            type="favorites"
            onPlayDiscord={this.onPlayDiscord}
            hasModPermissions={hasModPermissions()}
            onFavorite={this.onDeleteFavorite}
          />
        )}
        <SoundList
          soundList={soundList}
          title="Sounds"
          type="sounds"
          onPlayDiscord={this.onPlayDiscord}
          hasModPermissions={hasModPermissions()}
          onFavorite={this.onFavorite}
        />
      </div>
    );
  }
}
