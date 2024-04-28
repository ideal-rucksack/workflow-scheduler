export default [
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/dashboard',
    component: './Home',
    title: 'menu.dashboard',
    wrappers: ['@/wrappers/auth'],
  },
  {
    layout: false,
    path: '/signin',
    component: './signin',
    title: 'nav.left.signin',
  },
  {
    name: '权限演示',
    path: '/access',
    component: './Access',
    wrappers: ['@/wrappers/auth'],
  },
  {
    name: ' CRUD 示例',
    path: '/table',
    component: './Table',
    wrappers: ['@/wrappers/auth'],
  },
]
