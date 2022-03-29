import { DateTime } from 'luxon';
import { inject, observer } from 'mobx-react';
import React from 'react';
import { EmbeddedYoutube } from '../../components';
import { IVideoArchive, Permissions } from '../../model';
import { ArchiveService } from '../../services/archive.service';
import { AppStore } from '../../stores';
import './video-archive.scss';

interface IProps {
  appStore: AppStore;
}

interface IState {
  archives: IVideoArchive[];
  url: string;
  error?: string;
}

@inject('appStore')
@observer
export class VideoArchive extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      archives: [],
      url: '',
    };
  }

  componentDidMount() {
    this.loadArchives();
  }

  async loadArchives() {
    const archives = await ArchiveService.getVideoArchive();
    if (archives) {
      this.setState({
        archives,
      });
    }
  }

  onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const { url } = this.state;
    this.setState({ error: undefined });

    try {
      await ArchiveService.postVideoArchive({ url });
    } catch (e) {
      this.setState({ error: 'Invalid URL' });
      return;
    }

    this.setState({ url: '' });
    this.loadArchives();
  };

  renderForm() {
    const { error, url } = this.state;
    return (
      <form className="flex flex--v-center" onSubmit={this.onSubmit}>
        <input
          className="input video-archive__text-input"
          placeholder="Enter Youtube URL or ID..."
          value={url}
          onChange={(e) => this.setState({ url: e.target.value })}
        />
        <input type="submit" className="button button--primary" style={{ marginLeft: '10px' }} />
        {error && (
          <span className="color__red" style={{ marginLeft: '5px' }}>
            {error}
          </span>
        )}
      </form>
    );
  }

  renderArchives() {
    return this.state.archives.map((v, k) => (
      <div key={k} className="card video-archive__card test">
        <EmbeddedYoutube id={v.youtube_id} />
        <div style={{ padding: '10px 10px 0' }}>
          <h3 style={{ margin: '0 0 5px' }} className="ellipsis" title={v.title}>
            {v.title}
          </h3>
          <div>
            <span className="color__red">{v.uploaded_by}</span> -{' '}
            <small>{DateTime.fromISO(v.created_at).toLocaleString()}</small>
          </div>
        </div>
      </div>
    ));
  }

  render() {
    const { claims } = this.props.appStore;
    return (
      <div className="content">
        {claims && claims.permissions > Permissions.User && this.renderForm()}
        <div className="archive-grid">{this.renderArchives()}</div>
      </div>
    );
  }
}
