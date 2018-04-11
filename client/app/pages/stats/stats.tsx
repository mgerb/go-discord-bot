import React, { Component } from 'react';
import { HorizontalBar } from 'react-chartjs-2';
import { chain, map } from 'lodash';
import { axios } from '../../services';
import './stats.scss';

interface IState {
  data: {
    username: string;
    count: number;
  }[];
}

/**
 * a page to show discord chat statistics
 * currently keeps track of number messages that contain external links
 */
export class Stats extends Component<any, IState> {
  constructor(props: any) {
    super(props);
    this.state = {
      data: [],
    };
  }

  componentDidMount() {
    this.getdata();
  }

  async getdata() {
    const messages = await axios.get('/api/logger/linkedmessages');
    const data: any = chain(messages.data)
      .map((v, k) => {
        return { username: k, count: v };
      })
      .orderBy(v => v.count, 'desc')
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

    return (
      <div className="content">
        <div className="card" style={{ maxWidth: '1000px' }}>
          <div className="card__header">Shitposts</div>
          <HorizontalBar data={data} height={500} />
        </div>
      </div>
    );
  }
}
