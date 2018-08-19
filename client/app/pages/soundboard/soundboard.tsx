import React from 'react';
import { SoundList, SoundType, Uploader } from '../../components';
import { axios } from '../../services';
import './soundboard.scss';

interface Props {}

interface State {
  percentCompleted: number;
  uploaded: boolean;
  uploadError: string;
  soundList: SoundType[];
}

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

  render() {
    const { soundList } = this.state;
    return (
      <div className="content">
        <Uploader onComplete={this.onUploadComplete} />
        <SoundList soundList={soundList} type="Sounds" />
      </div>
    );
  }
}
