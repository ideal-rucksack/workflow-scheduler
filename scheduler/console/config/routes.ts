export default [
  {
    path: '/*',
    component: './404.tsx',
    layout: false,
  },
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/dashboard',
    component: './dashboard',
    title: 'menu.dashboard',
    icon: 'DashboardOutlined',
    wrappers: ['@/wrappers/auth'],
  },
  {
    layout: false,
    path: '/signin',
    component: './signin',
    title: 'nav.left.signin',
  },
]
