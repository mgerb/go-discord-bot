import React from 'react';
import { Navbar } from './components/Navbar/Navbar.component';

//styling
import './scss/index.scss';

interface Props {

}

interface State {

}

export class Wrapper extends React.Component<Props, State> {

    constructor() {
        super();
    }

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
