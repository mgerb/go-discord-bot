import React from 'react';
import { Link } from 'react-router';

import './Navbar.scss';

export default class Navbar extends React.Component {

    render() {
        return (
            <div className="Navbar">
                <div className="Navbar__header">Sound<br/>Bot</div>
                <Link to="/" className="Navbar__item" onlyActiveOnIndex activeClassName="Navbar__item--active">Soundboard</Link>
            </div>
        );
    }
}

Navbar.propTypes = {
    children: React.PropTypes.node,
};
