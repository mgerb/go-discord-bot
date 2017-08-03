import React from 'react';

import './Home.scss';

interface Props {

}

interface State {

}

export class Home extends React.Component<Props, State> {

    render() {
        return (
            <div className="Home">
                <div className="Card">
                    <div className="Card__header">
                        Go Discord Bot
                    </div>
                    <p>Drag and drop files to upload. Sounds can be played in discord by typing the commands on the next page.</p>
                    <br/>
                    <p>Gif command now supported! Example: !gif awesome cat gifs</p>
                    <br/>
                    <p>Check out the source code on 
                        <a href="https://github.com/mgerb/GoBot" target="_blank"> GitHub 
                            <i className="fa fa-github" aria-hidden="true"/>
                        </a>
                    </p>
                </div>
            </div>
        );
    }
}
