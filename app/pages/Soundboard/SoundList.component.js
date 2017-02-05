import React from 'react';
import axios from 'axios';

import './SoundList.scss';

export default class SoundList extends React.Component {
    
    constructor() {
        super();
        this.state = {
            soundList: [],
        };
    }
    
    componentDidMount() {
        this.getSoundList();
    }
    
    getSoundList() {
        if (!this.soundListCache) {
            axios.get("/soundlist")
                .then((response) => {
                    this.soundListCache = response.data;
                    this.setState({
                        soundList: response.data,
                    });
                }).catch((error) => {
                    console.warn(error.response.data);
                });
        } else {
            this.setState({
                soundList: this.soundListCache,
            });
        }
    }
    
    checkExtension(extension) {
        switch(extension) {
            case "wav":
                return true;
            case "mp3":
                return true;
            case "mpeg":
                return true;
            default:
                return false;
        }
    }
    
    render() {
        return (
            <div className="Card">
                <div className="Card__header">
                    Sounds
                <i className="fa fa fa-volume-up" aria-hidden="true"></i>
                </div>
                
                {this.state.soundList.length > 0 ? this.state.soundList.map((sound, index) => {
                    return <div key={index} className="SoundList__item">
                        <div>
                            {sound.prefix + sound.name}
                        </div>
                        
                        {this.checkExtension(sound.extension) ?
                        <audio controls src={"/sounds/" + sound.name + "." + sound.extension}
                                type={"audio/" + sound.extension}
                                style={{width: "100px"}}/>
                        : null}
                    </div>
                }) : null}
            </div>
        );
    }
}
