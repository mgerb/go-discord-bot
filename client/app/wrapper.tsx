import { inject, observer } from 'mobx-react';
import React from 'react';
import { withRouter } from 'react-router';
import { Header, Navbar } from './components';
import { Util } from './util';
import './wrapper.scss';

export const Wrapper = inject('appStore')(
  withRouter(
    // eslint-disable-next-line
    observer(({ appStore, children }: any) => {
      const openClass = appStore.navbarOpen ? 'wrapper--open' : '';
      const onNavClick = () => {
        if (Util.isMobileScreen()) {
          appStore.toggleNavbar();
        }
      };

      return (
        <div>
          <Header onButtonClick={appStore.toggleNavbar} />
          <Navbar appStore={appStore} onNavClick={onNavClick} />
          <div className={'wrapper ' + openClass}>{children}</div>
        </div>
      );
    }),
  ),
);
