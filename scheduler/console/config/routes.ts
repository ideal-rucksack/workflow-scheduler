export default [
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/dashboard',
    component: './Home',
    title: 'menu.dashboard',
  },
  {
    layout: false,
    path: '/login',
    component: './login',
    title: 'nav.left.login',
  },
  {
    name: '权限演示',
    path: '/access',
    component: './Access',
  },
  {
    name: ' CRUD 示例',
    path: '/table',
    component: './Table',
  },
]
