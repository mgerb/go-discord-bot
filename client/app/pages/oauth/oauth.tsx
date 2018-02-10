import React from 'react';
import { get } from 'lodash';
import axios from 'axios';

interface Props {
    
}

interface State {
    
}

export class Oauth extends React.Component<Props, State> {
    
    constructor(props: Props) {
        super(props);
    }
    
    componentDidMount() {
        const code = get(this, 'props.location.query.code');
        
        if (code) {
            // do stuff here
            this.fetchOauth(code as string);
        }
    }
    
    private async fetchOauth(code: string) {
        const res = await axios.post('/api/oauth', { code });
        console.log(res);
    }
    
    render() {
        return <div></div>
    }
}
