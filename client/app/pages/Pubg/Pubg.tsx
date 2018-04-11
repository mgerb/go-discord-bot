/**
 * DEPRECATED
 */

import React from 'react';
import { axios } from '../../services';
import * as _ from 'lodash';
import './Pubg.scss';

interface Props {}

interface State {
  players: Player[];
  selectedRegion: string;
  selectedMatch: string;
  statList: string[];
}

interface Player {
  PlayerName: string;
  agg?: any;
  as?: any;
  na?: any;
  sa?: any;
}

export class Pubg extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      players: [],
      selectedRegion: 'agg',
      selectedMatch: 'squad',
      statList: [],
    };
  }

  componentDidMount() {
    axios.get('/api/stats/pubg').then(res => {
      this.setState({
        players: _.map(res.data) as any,
      });

      this.setStatList();
    });
  }

  // get stat list
  setStatList() {
    // hacky way to find existing content -- to tired to make it pretty
    let i = 0;
    let stats;
    while (!stats) {
      if (i > this.state.players.length) {
        return;
      }

      stats = _.find(
        _.get(this.state, `players[${i}].Stats`),
        (s: any) => s.Match === this.state.selectedMatch.toLowerCase(),
      );
      i++;
    }

    if (stats) {
      this.setState({
        statList: _.sortBy(_.map(stats.Stats, 'field')) as any,
      });
    }
  }

  insertRows(): any {
    return this.state.statList.map((val: any, index: any) => {
      return (
        <tr key={index}>
          <td>{val}</td>
          {this.state.players.map((player: any, i: number) => {
            // find player stats for field
            let playerStat = _.find(player.Stats, (p: any) => {
              return (
                p.Match === this.state.selectedMatch.toLowerCase() &&
                p.Region === this.state.selectedRegion.toLowerCase()
              );
            });

            return (
              <td key={i}>{_.get(_.find(_.get(playerStat, 'Stats'), (p: any) => p.field === val), 'displayValue')}</td>
            );
          })}
        </tr>
      );
    });
  }

  buttonRegion(title: string) {
    let lowerTitle = title === 'All' ? 'agg' : title.toLowerCase();
    return (
      <button
        className={`button ${lowerTitle === this.state.selectedRegion ? 'button--primary' : ''}`}
        onClick={() => {
          this.setState({ selectedRegion: lowerTitle });
          this.setStatList();
        }}
      >
        {title}
      </button>
    );
  }

  buttonMatch(title: string) {
    let lowerTitle = title.toLowerCase();
    return (
      <button
        className={`button ${lowerTitle === this.state.selectedMatch ? 'button--primary' : ''}`}
        onClick={() => {
          this.setState({ selectedMatch: lowerTitle });
          this.setStatList();
        }}
      >
        {title}
      </button>
    );
  }

  render() {
    return (
      <div className="pubg__container">
        <div className="card" style={{ maxWidth: 'initial' }}>
          <div className="card__header">PUBG Stats</div>

          <div className="pubg__button-row">
            {this.buttonMatch('Solo')}
            {this.buttonMatch('Duo')}
            {this.buttonMatch('Squad')}
          </div>

          <div className="pubg__button-row">
            {this.buttonRegion('All')}
            {this.buttonRegion('Na')}
            {this.buttonRegion('As')}
            {this.buttonRegion('Au')}
          </div>

          <table className="pubg__table">
            <tbody>
              <tr>
                <th />
                {this.state.players.map((val: any, index: number) => {
                  return <th key={index}>{val.PlayerName}</th>;
                })}
              </tr>
              {this.insertRows()}
            </tbody>
          </table>
        </div>
      </div>
    );
  }
}
