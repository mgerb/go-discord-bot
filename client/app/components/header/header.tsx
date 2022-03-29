import React from 'react';
import './header.scss';

interface IProps {
  onButtonClick: () => void;
}

export class Header extends React.Component<IProps, unknown> {
  constructor(props: IProps) {
    super(props);
  }

  render() {
    return (
      <div className="header">
        <div className="header__title-container">
          <button className="header__nav-button" onClick={this.props.onButtonClick}>
            <i className="fa fa-lg fa-bars" />
          </button>
          <h2 style={{ margin: 0 }}>Sound Bot</h2>
        </div>
        <a
          rel="noreferrer"
          href="https://github.com/mgerb/go-discord-bot"
          className="fa fa-lg fa-github"
          target="_blank"
        />
      </div>
    );
  }
}
