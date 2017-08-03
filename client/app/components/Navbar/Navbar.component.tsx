import React from 'react';
import { Link } from 'react-router';

import './Navbar.scss';

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
            </div>
        );
    }
}
