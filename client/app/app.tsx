import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { Wrapper } from './Wrapper';
import { Home } from './pages/Home/Home';
import { Soundboard } from './pages/Soundboard/Soundboard';
import { NotFound } from './pages/NotFound/NotFound';
import { Downloader } from './pages/Downloader/Downloader';
import { Clips } from './pages/Clips';
import { Oauth } from './pages/oauth/oauth';
import { Stats } from './pages/stats/stats';
import 'babel-polyfill';

const App: any = (): any => {
  return (
    <BrowserRouter>
      <Wrapper>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/soundboard" component={Soundboard} />
          <Route path="/downloader" component={Downloader} />
          <Route path="/clips" component={Clips} />
          <Route path="/oauth" component={Oauth} />
          <Route path="/stats" component={Stats} />
          <Route component={NotFound} />
        </Switch>
      </Wrapper>
    </BrowserRouter>
  );
};

ReactDOM.render(<App /> as any, document.getElementById('app'));
