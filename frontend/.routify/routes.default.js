

export default {
  "meta": {},
  "id": "_default",
  "_regex": {},
  "_paramKeys": {},
  "file": {
    "path": "src/routes/_module.svelte",
    "dir": "src/routes",
    "base": "_module.svelte",
    "ext": ".svelte",
    "name": "_module"
  },
  "asyncModule": () => import('../src/routes/_module.svelte'),
  "rootName": "default",
  "routifyDir": import.meta.url,
  "children": [
    {
      "meta": {},
      "id": "_default_dashboard",
      "_regex": {},
      "_paramKeys": {},
      "name": "dashboard",
      "module": false,
      "file": {
        "path": "src/routes/dashboard",
        "dir": "src/routes",
        "base": "dashboard",
        "ext": "",
        "name": "dashboard"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_dashboard_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/dashboard/index.svelte",
            "dir": "src/routes/dashboard",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/dashboard/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_index_svelte",
      "_regex": {},
      "_paramKeys": {},
      "name": "index",
      "file": {
        "path": "src/routes/index.svelte",
        "dir": "src/routes",
        "base": "index.svelte",
        "ext": ".svelte",
        "name": "index"
      },
      "asyncModule": () => import('../src/routes/index.svelte'),
      "children": []
    },
    {
      "meta": {},
      "id": "_default_login",
      "_regex": {},
      "_paramKeys": {},
      "name": "login",
      "module": false,
      "file": {
        "path": "src/routes/login",
        "dir": "src/routes",
        "base": "login",
        "ext": "",
        "name": "login"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_login_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/login/index.svelte",
            "dir": "src/routes/login",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/login/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_profile",
      "_regex": {},
      "_paramKeys": {},
      "name": "profile",
      "module": false,
      "file": {
        "path": "src/routes/profile",
        "dir": "src/routes",
        "base": "profile",
        "ext": "",
        "name": "profile"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_profile_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/profile/index.svelte",
            "dir": "src/routes/profile",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/profile/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_register",
      "_regex": {},
      "_paramKeys": {},
      "name": "register",
      "module": false,
      "file": {
        "path": "src/routes/register",
        "dir": "src/routes",
        "base": "register",
        "ext": "",
        "name": "register"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_register_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/register/index.svelte",
            "dir": "src/routes/register",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/register/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {
        "dynamic": true,
        "dynamicSpread": true
      },
      "_regex": {},
      "_paramKeys": {},
      "name": "[...404]",
      "file": {
        "path": ".routify/components/[...404].svelte",
        "dir": ".routify/components",
        "base": "[...404].svelte",
        "ext": ".svelte",
        "name": "[...404]"
      },
      "asyncModule": () => import('./components/[...404].svelte'),
      "children": []
    }
  ]
}