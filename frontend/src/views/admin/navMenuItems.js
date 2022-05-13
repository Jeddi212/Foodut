/*=========================================================================================
  File Name: sidebarItems.js
  Description: Sidebar Items list. Add / Remove menu items from here.
  Strucutre:
          url     => router path
          name    => name to display in sidebar
          slug    => router path name
          icon    => Feather Icon component/icon name
          tag     => text to display on badge
          tagColor  => class to apply on badge element
          i18n    => Internationalization
          submenu   => submenu of current item (current item will become dropdown )
                NOTE: Submenu don't have any icon(you can add icon if u want to display)
          isDisabled  => disable sidebar item/group
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
==========================================================================================*/


export default [
    {
      url: null,
      name: 'Authentication',
      icon: 'HomeIcon',
      i18n: 'Authentication' ,
      submenu: [
        {
          url: '/',
          name: 'Login',
          slug: 'pages-login',
          i18n: 'Login',
          target: '_blank'
        },
        {
          url: '/register',
          name: 'Register',
          slug: 'pages-register',
          i18n: 'Register',
          target: '_blank'
        }
      ]
    },
    {
      header: 'Apps',
      icon: 'PackageIcon',
      i18n: 'AppsPages',
      items: [
        {
          url: null,
          name: 'Store',
          icon: 'ShoppingCartIcon',
          i18n: 'Store',
          submenu: [
            {
              url: '/admin/browse',
              name: 'Browse',
              slug: 'admin-browse',
              i18n: 'Browse'
            },
            {
              url: '/admin/user',
              name: 'List User',
              slug: 'admin-user',
              i18n: 'ListUser'
            },
            {
              url: '/admin/store',
              name: 'List Store',
              slug: 'admin-store',
              i18n: 'ListStore'
            }
          ]
        }
      ]
    },
    {
      header: 'Test Api',
      icon: 'PackageIcon',
      i18n: 'Test Api',
      items: [
        {
          url: '/admin/test',
          name: 'Test Api',
          icon: 'ShoppingCartIcon',
          slug: 'admin-test',
          i18n: 'Test Api',
        }
      ]
    }
  ]
  
  