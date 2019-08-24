import { chain, map } from 'lodash';
import { inject, observer } from 'mobx-react';
import React, { Component } from 'react';
import { HorizontalBar } from 'react-chartjs-2';
import { UploadHistory } from '../../components';
import { ISound } from '../../model';
import { axios, SoundService } from '../../services';
import { AppStore } from '../../stores';
import './stats.scss';

interface IProps {
  appStore: AppStore;
}

interface IState {
  data: {
    username: string;
    count: number;
  }[];
  sounds: ISound[];
}

/**
 * a page to show discord chat statistics
 * currently keeps track of number messages that contain external links
 */
@inject('appStore')
@observer
export class Stats extends Component<IProps, IState> {
  constructor(props: any) {
    super(props);
    this.state = {
      data: [],
      sounds: [],
    };
  }

  componentDidMount() {
    this.getdata();
    SoundService.getSounds().then(sounds => {
      this.setState({ sounds });
    });
  }

  async getdata() {
    const messages = await axios.get('/api/logger/linkedmessages');
    const data: any = chain(messages.data)
      .map((v, k) => {
        return { username: k, count: v };
      })
      .orderBy(v => v.count, 'desc')
      .slice(0, 10)
      .value();

    this.setState({ data });
  }

  render() {
    const data: any = {
      labels: map(this.state.data, v => v.username),
      datasets: [
        {
          label: 'Count',
          backgroundColor: 'rgba(114,137,218, 0.5)',
          borderColor: 'rgba(114,137,218, 0.9)',
          borderWidth: 1,
          hoverBackgroundColor: 'rgba(114,137,218, 0.7)',
          hoverBorderColor: 'rgba(114,137,218, 1)',
          data: map(this.state.data, v => v.count),
        },
      ],
      options: {
        responsive: true,
      },
    };

    const { claims, hasModPermissions } = this.props.appStore;

    return (
      <div className="content">
        <div className="card">
          <div className="card__header">Posts containing links</div>
          <HorizontalBar data={data} />
        </div>
        {claims && <UploadHistory sounds={this.state.sounds} showDiscordPlay={hasModPermissions()} />}
      </div>
    );
  }
}
