import React from 'react';
import Dropzone from 'react-dropzone';
import { axios } from '../../services';
import { SoundList, SoundType } from '../../components/SoundList';
import './Soundboard.scss';
import { AxiosRequestConfig } from 'axios';

let self: any;

interface Props {}

interface State {
  percentCompleted: number;
  uploaded: boolean;
  uploadError: string;
  soundList: SoundType[];
}

export class Soundboard extends React.Component<Props, State> {
  private config: AxiosRequestConfig;
  private soundListCache: any;

  constructor(props: Props) {
    super(props);
    (this.state = {
      percentCompleted: 0,
      uploaded: false,
      uploadError: ' ',
      soundList: [],
    }),
      (self = this);
  }

  componentDidMount() {
    this.config = {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      onUploadProgress: progressEvent => {
        this.setState({
          percentCompleted: Math.round(progressEvent.loaded * 100 / progressEvent.total),
        });
      },
    };

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

  onDrop(acceptedFiles: any) {
    if (acceptedFiles.length > 0) {
      self.uploadFile(acceptedFiles[0]);
    }
  }

  uploadFile(file: any) {
    let formData = new FormData();
    formData.append('name', file.name);
    formData.append('file', file);

    axios
      .post('/api/upload', formData, this.config)
      .then(() => {
        this.setState({
          percentCompleted: 0,
          uploaded: true,
          uploadError: ' ',
        });

        this.soundListCache = undefined;
        this.getSoundList();
      })
      .catch(err => {
        this.setState({
          percentCompleted: 0,
          uploaded: false,
          uploadError: err.response.data,
        });
      });
  }

  render() {
    const { soundList } = this.state;
    return (
      <div className="Soundboard">
        <div className="column">
          <SoundList soundList={soundList} type="Sounds" />
        </div>

        <div className="column">
          <div>
            <Dropzone
              className="Dropzone"
              activeClassName="Dropzone--active"
              onDrop={this.onDrop}
              multiple={false}
              disableClick={true}
              maxSize={10000000000}
              accept={'audio/*'}
            >
              <div style={{ fontSize: '20px' }}>Drop file here to upload.</div>
              {this.state.percentCompleted > 0 ? <div>Uploading: {this.state.percentCompleted}</div> : ''}
              {this.state.uploaded ? <div style={{ color: 'green' }}>File uploded!</div> : ''}
              <div style={{ color: '#f95f59' }}>{this.state.uploadError}</div>
            </Dropzone>
          </div>
        </div>
      </div>
    );
  }
}
