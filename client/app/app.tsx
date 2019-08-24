import { Provider } from 'mobx-react';
import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Admin, Clips, Downloader, NotFound, Oauth, Soundboard, Stats, UploadHistory, VideoArchive } from './pages';
import './scss/index.scss';
import { rootStoreInstance } from './stores';
import { Wrapper } from './wrapper';

const App: any = (): any => {
  return (
    <BrowserRouter>
      <Provider {...rootStoreInstance}>
        <Wrapper>
          <Switch>
            <Route exact path="/" component={Soundboard} />
            <Route path="/upload-history" component={UploadHistory} />
            <Route path="/downloader" component={Downloader} />
            <Route path="/clips" component={Clips} />
            <Route path="/oauth" component={Oauth} />
            <Route path="/stats" component={Stats} />
            <Route path="/video-archive" component={VideoArchive} />
            <Route path="/admin" component={Admin} />
            <Route component={NotFound} />
          </Switch>
        </Wrapper>
      </Provider>
    </BrowserRouter>
  );
};

ReactDOM.render(<App /> as any, document.getElementById('app'));
