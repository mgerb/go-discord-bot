import React from 'react';
import { IUser } from '../../model';
import { UserService } from '../../services/user.service';

interface IState {
  users: IUser[];
  showSavedMessage: boolean;
}

export class Users extends React.Component<any, IState> {
  constructor(props: any) {
    super(props);
    this.state = {
      users: [],
      showSavedMessage: false,
    };
  }

  componentDidMount() {
    UserService.getUsers().then((users) => {
      this.setState({ users });
    });
  }

  onUserChange = (type: 'permissions' | 'voice_join_sound', event: any, index: number) => {
    this.setState({ showSavedMessage: false });

    let users = [...this.state.users];
    const val = type === 'permissions' ? parseInt(event.target.value) : event.target.value;
    users[index] = {
      ...users[index],
      [type]: val,
    };

    this.setState({ users });
  };

  renderUsers = () => {
    return this.state.users.map((u: IUser, i: number) => {
      return (
        <tr key={i}>
          <td>{u.username}</td>
          <td>{u.email}</td>
          <td>
            <input
              className="input"
              type="number"
              min={1}
              max={3}
              value={u.permissions}
              onChange={(e) => this.onUserChange('permissions', e, i)}
            />
          </td>
          <td>
            <input
              className="input"
              type="text"
              value={u.voice_join_sound || ''}
              onChange={(e) => this.onUserChange('voice_join_sound', e, i)}
            />
          </td>
        </tr>
      );
    });
  };

  save = (event: any) => {
    this.setState({ showSavedMessage: false });
    event.preventDefault();
    UserService.putUsers(this.state.users).then((users: IUser[]) => {
      this.setState({ users, showSavedMessage: true });
    });
  };

  render() {
    return (
      <form className="content" onSubmit={this.save}>
        <div className="card card--wide">
          <div className="card__header">Users</div>
          <div className="overflow-x-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Username</th>
                  <th>Email</th>
                  <th>Permissions</th>
                  <th>Join Sound</th>
                </tr>
              </thead>
              <tbody>{this.renderUsers()}</tbody>
            </table>
          </div>

          <br />
          <button className="button button--primary" type="submit">
            Save
          </button>
          {this.state.showSavedMessage && <span style={{ marginLeft: 5 }}>Users updated</span>}
        </div>
      </form>
    );
  }
}
