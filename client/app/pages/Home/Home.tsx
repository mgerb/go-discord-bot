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
                    <h3>Audio Clipping</h3>
                    <p><em>NEW:</em> Audio clipping now supported! Try it out with the <code>clip</code> command!</p>

                    <br/>

                    <h3>PUBG Stats</h3>
                    <p>PUBG stats are pulled from the score API.</p>

                    <br/>

                    <h3>Youtube Downloader</h3>
                    <p>Convert Youtube URL's to MP3 files.</p>

                    <br/>

                    <h3>Soundboard Upload</h3>
                    <p>Drag and drop files to upload. Sounds can be played in discord by typing the commands on the next page.</p>

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
