import React from 'react';
import Dropzone from 'react-dropzone';
import { axios } from '../../services';
import './uploader.scss';

interface IProps {
  onComplete: () => void;
}

interface IState {
  percentCompleted: number;
  uploaded: boolean;
  uploadError: string;
}

export class Uploader extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      percentCompleted: 0,
      uploaded: false,
      uploadError: ' ',
    };
  }

  private config = {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
    onUploadProgress: (progressEvent: any) => {
      this.setState({
        percentCompleted: Math.round((progressEvent.loaded * 100) / progressEvent.total),
      });
    },
  };

  onDrop = (acceptedFiles: any) => {
    if (acceptedFiles.length > 0) {
      this.uploadFile(acceptedFiles[0]);
    }
  };

  uploadFile(file: any) {
    let formData = new FormData();
    formData.append('name', file.name);
    formData.append('file', file);

    axios
      .post('/api/sound', formData, this.config)
      .then(() => {
        this.setState({
          percentCompleted: 0,
          uploaded: true,
          uploadError: ' ',
        });

        this.props.onComplete();
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
    return (
      <Dropzone
        className="dropzone"
        activeClassName="dropzone--active"
        onDrop={this.onDrop}
        multiple={false}
        disableClick={false}
        maxSize={10000000000}
        accept={'audio/*'}
      >
        <div style={{ fontSize: '20px' }}>Click or drop file to upload</div>
        {this.state.percentCompleted > 0 ? <div>Uploading: {this.state.percentCompleted}</div> : ''}
        {this.state.uploaded ? <div style={{ color: 'green' }}>File uploded!</div> : ''}
        <div style={{ color: '#f95f59' }}>{this.state.uploadError}</div>
      </Dropzone>
    );
  }
}
