import React from 'react';
import './Home.scss';

export default class Home extends React.Component {

    render() {
        return (
            <div className="Home">
                <div className="Card">
                    <div className="Card__header">
                        GoBot Early Access Pre Pre Alpha
                    </div>
                    <p>This application is a work in progress.</p>
                    <p>Check out the source code on 
                        <a href="https://github.com/mgerb/GoBot" target="_blank"> GitHub 
                            <i className="fa fa-github" aria-hidden="true"></i>
                        </a>
                    </p>
                </div>
            </div>
        );
    }
}