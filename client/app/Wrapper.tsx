import React from 'react';
import { Navbar } from './components/Navbar';

//styling
import './scss/index.scss';

export class Wrapper extends React.Component<any, any> {
  constructor(props: any) {
    super(props);
  }

  render() {
    return (
      <div>
        <Navbar />
        <div>{this.props.children}</div>
      </div>
    );
  }
}
