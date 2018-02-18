import React from 'react';
import { Link } from 'react-router';

import './Navbar.scss';

// TODO: change url for build
const redirectUri = 'https://localhost/oauth';
const oauthUrl = `https://discordapp.com/api/oauth2/authorize?client_id=410818759746650140&redirect_uri=${redirectUri}&response_type=code&scope=guilds%20identify`;

interface Props {

}

interface State {

}

export class Navbar extends React.Component<Props, State> {

    render() {
        return (
            <div className="Navbar">
                <div className="Navbar__header">Go Discord Bot</div>
                <Link to="/" className="Navbar__item" onlyActiveOnIndex activeClassName="Navbar__item--active">Home</Link>
                <Link to="/soundboard" className="Navbar__item" activeClassName="Navbar__item--active">Soundboard</Link>
                <Link to="/downloader" className="Navbar__item" activeClassName="Navbar__item--active">Youtube Downloader</Link>
                <Link to="/pubg" className="Navbar__item" activeClassName="Navbar__item--active">Pubg</Link>
                <Link to="/clips" className="Navbar__item" activeClassName="Navbar__item--active">Clips</Link>
                <a href={oauthUrl} className="Navbar__item">Login</a>
            </div>
        );
    }
}
