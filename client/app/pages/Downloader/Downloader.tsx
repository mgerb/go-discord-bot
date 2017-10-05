import React from 'react';
import axios from 'axios';

import './Downloader.scss';

interface Props {

}

interface State {
    fileType: string;
    url: string;
    message: string;
    downloadLink: string;
    downLoadFileName: string;
    dataLoading: boolean;
    dataLoaded: boolean;
}

export class Downloader extends React.Component<Props, State> {

    constructor(props: Props) {
        super(props);
        this.state = {
            fileType: "mp3",
            url: "",
            message: "",
            dataLoaded: false,
            downloadLink: "",
            downLoadFileName: "",
            dataLoading: false,
        };
    }

    sendRequest() {
        if (this.state.url === "") {
            this.setState({
                message: "Invalid URL",
            });

            return;
        }

        this.setState({
            message: "",
            url: "",
            dataLoaded: false,
            dataLoading: true,
        });

        axios.get(`/api/ytdownloader`, {
            params: {
                fileType: this.state.fileType,
                url: this.state.url,
            }
        }).then((res) => {
            this.setState({
                dataLoaded: true,
                dataLoading: false,
                downloadLink: `/public/youtube/${res.data.fileName}`,
                downLoadFileName: res.data.fileName,
            });
        }).catch(() => {
            this.setState({
                message: "Internal error.",
                dataLoading: false,
            });
        });
    }

    render() {
        return (
            <div className="Downloader">
                <div className="card">
                    <div className="card__header">
                        Youtube to MP3
                    </div>

                    <input placeholder="Enter Youtube URL"
                            className="input Downloader__input"
                            value={this.state.url}
                            onChange={(event) => this.setState({url: event.target.value})}/>

                    <div style={{marginBottom:'10px'}}>
                        <button className="button button--primary"
                                style={{width: '100px', height: '40px', fontSize: 'large'}}
                                onClick={this.sendRequest.bind(this)}>
                            {this.state.dataLoading ? <i className="fa fa-spinner fa-spin fa-fw"/> : 'Convert'}
                        </button>
                    </div>

                    {this.state.message !== "" && <div>{this.state.message}</div>}
                    {this.state.dataLoaded && <a href={this.state.downloadLink} download>{this.state.downLoadFileName}</a>}
                </div>

            </div>
        );
    }
}
