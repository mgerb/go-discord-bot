import React from 'react';
import './NotFound.scss';

interface Props {

}

interface State {

}

export class NotFound extends React.Component<Props, State> {
    render() {
        return (
            <div className="NotFound">
                404 Not Found
            </div>
        );
    }
}
