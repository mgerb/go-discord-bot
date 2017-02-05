import React from 'react';
import Navbar from './components/Navbar/Navbar.component';

//styling
import './scss/index.scss';

export default class Wrapper extends React.Component {

    render() {
        return (
            <div>
                <Navbar/>
                <div>
                    {this.props.children}
                </div>
            </div>
        );
    }
}

Wrapper.propTypes = {
    children: React.PropTypes.node,
};
