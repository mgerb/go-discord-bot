import React from 'react';
import { SoundList } from '../../components';
import { SoundType } from '../../model';
import { axios } from '../../services';

interface Props {}

interface State {
  clipList: SoundType[];
}

export class Clips extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      clipList: [],
    };
  }

  componentDidMount() {
    this.getClipList();
  }

  private getClipList() {
    axios
      .get('/api/cliplist')
      .then((response) => {
        this.setState({
          clipList: response.data,
        });
      })
      .catch((error: any) => {
        console.error(error.response.data);
      });
  }

  render() {
    return (
      <div className="content">
        <div className="column">
          {/* no need for permissions on this component - set false */}
          <SoundList soundList={this.state.clipList} type="clips" title="Clips" hasModPermissions={false} />
        </div>
      </div>
    );
  }
}
