import React from 'react';
import { Link } from 'react-router';

import './Navbar.scss';

export default class Navbar extends React.Component {

    render() {
        return (
            <div className="Navbar">
                <div className="Navbar__header">GoBot</div>
                <Link to="/" className="Navbar__item" onlyActiveOnIndex activeClassName="Navbar__item--active">Home</Link>
                <Link to="/soundboard" className="Navbar__item" activeClassName="Navbar__item--active">Soundboard</Link>
                <div className="link">
                    <a href="https://discordapp.com/invite/0Z2tzxKECEj2BHwj" target="_blank">Join the discord</a>
                </div>
            </div>
        );
    }
}

Navbar.propTypes = {
    children: React.PropTypes.node,
};