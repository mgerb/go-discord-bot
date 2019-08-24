import React from 'react';
import { IUserEventLog } from '../../model';
import { UserEventLogService } from '../../services';

interface IProps {}

interface IState {
  userEventLogs: IUserEventLog[];
}

export class Admin extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      userEventLogs: [],
    };
  }
  componentDidMount() {
    UserEventLogService.getUserEventLogs().then(userEventLogs => {
      this.setState({
        userEventLogs,
      });
    });
  }

  renderUserEventLogs() {
    return this.state.userEventLogs.map(({ id, user, content, created_at }, index) => {
      return (
        <tr key={index}>
          <td>{id}</td>
          <td>{created_at}</td>
          <td>{user.username}</td>
          <td>{content}</td>
        </tr>
      );
    });
  }

  render() {
    return (
      <div className="content">
        <div className="card">
          <div className="card__header">User Event Log</div>
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Timestamp</th>
                <th>User</th>
                <th>Content</th>
              </tr>
            </thead>
            <tbody>{this.renderUserEventLogs()}</tbody>
          </table>
        </div>
      </div>
    );
  }
}
