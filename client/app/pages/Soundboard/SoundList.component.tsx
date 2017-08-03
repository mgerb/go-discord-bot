import React from 'react';
import axios from 'axios';

import './SoundList.scss';

interface Props {

}

interface State {
    showAudioControls: boolean[];
    soundList: {
        extension: string;
        name: string;
        prefix: string;
    }[];
}

export class SoundList extends React.Component<Props, State> {
    
    private soundListCache: any;

    constructor() {
        super();
        this.state = {
            soundList: [],
            showAudioControls: [],
        };
    }
    
    componentDidMount() {
        this.getSoundList();
    }
    
    getSoundList() {
        if (!this.soundListCache) {
            axios.get("/soundlist").then((response) => {
                this.soundListCache = response.data;
                this.setState({
                    soundList: response.data,
                });
            }).catch(() => {
                //console.warn(error.response.data);
            });
        } else {
            this.setState({
                soundList: this.soundListCache,
            });
        }
    }
    
    checkExtension(extension: string) {
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

    handleShowAudio(index: any) {
        let temp = this.state.showAudioControls;
        temp[index] = true;

        this.setState({
            showAudioControls: temp,
        });
    }
    
    render() {
        return (
            <div className="Card">
                <div className="Card__header" style={{display:'flex'}}>
                    <div>
                        Sounds
                        <i className="fa fa fa-volume-up" aria-hidden="true"/>
                    </div>
                    <div style={{flex:1}}/>
                    <div>({this.state.soundList.length})</div>
                </div>
                
                {this.state.soundList.length > 0 ? this.state.soundList.map((sound, index) => {
                    return (
                        <div key={index} className="SoundList__item">
                            <div>
                                {sound.prefix + sound.name}
                            </div>
                            
                            {this.checkExtension(sound.extension) && this.state.showAudioControls[index] ?
                            <audio controls src={"/public/sounds/" + sound.name + "." + sound.extension}
                                    itemType={"audio/" + sound.extension}
                                    style={{width: "100px"}}/>
                            : <i className="fa fa-play link" aria-hidden="true" onClick={() => this.handleShowAudio(index)}/> }
                        </div>
                    );
                }) : null}
            </div>
        );
    }
}
