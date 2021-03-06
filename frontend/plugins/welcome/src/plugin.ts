import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/', login);
  },
});
