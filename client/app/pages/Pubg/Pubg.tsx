import React from 'react';
import axios from 'axios';
import * as _ from 'lodash';
import './Pubg.scss';

interface Props {

}

interface State {
    players: Player[];
}

interface Player {
    PlayerName: string;
    agg?: any;
    as?: any;
    na?: any;
    sa?: any;
}

export class Pubg extends React.Component<Props, State> {

    constructor() {
        super();
        this.state = {
            players: [],
        };
    }

    componentDidMount() {
        axios.get("/stats/pubg").then((res) => {

        console.log(res.data);
            this.setState({
                players: this.filterData(res.data),
            });

            console.log(this.state.players);
        });
    }

    filterData(data: any): Player[] {
        return _.map(_.values(data), (data: any) => {

            let regions: any = _.chain(data.Stats).groupBy('Region')
            /*
            .mapValues((val: any) => {
                return  _.groupBy(val, 'Match');
            })
            */
            .value();


            _.forIn(regions, (val: any, key: string) => {
                regions[key] = _.groupBy(val, 'Match');

                _.forIn(regions[key], (val2: any, key2: string) => {
                    //regions[key][key2] = _.groupBy(regions[key][key2][0].Stats, 'field');
                    regions[key][key2] = _.groupBy(_.flatten(regions[key][key2][0].Stats), 'field');
                    _.each(regions[key][key2], s => s = _.flatten(s));

                    //console.log(regions[key][key2][0]);

                });
            });

            //console.log(regions);

            return {
                ...{ PlayerName: data.PlayerName },
                ...regions,
            };
        });
    }

    insertTableData() {

        return this.state.players.map((stat: any, index: number) => {
            return (
                <tr key={index}>
                    <td>{stat.PlayerName}</td>
                </tr>
            );
        });
    }

    render() {
        return (
            <div className="pubg__container">
                <div className="card">
                    <div className="card__header">PUBG Stats</div>
                    <table className="pubg__table">
                        <thead>
                            <tr>
                                <th>Name</th>
                            </tr>
                        </thead>
                        <tbody>
                            {this.insertTableData()}
                        </tbody>
                    </table>

                </div>
            </div>
        );
    }
}
