/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as PublicImport } from './routes/_public'
import { Route as ProtectedImport } from './routes/_protected'
import { Route as ProtectedWorkspaceIndexImport } from './routes/_protected/workspace/index'
import { Route as ProtectedWorkspaceNamespaceRouteImport } from './routes/_protected/workspace/$namespace/route'
import { Route as ProtectedWorkspaceNamespaceNamespaceImport } from './routes/_protected/workspace/$namespace/_namespace'
import { Route as ProtectedWorkspaceNamespaceNamespaceIndexImport } from './routes/_protected/workspace/$namespace/_namespace/index'
import { Route as ProtectedWorkspaceNamespaceNamespaceTransactionsImport } from './routes/_protected/workspace/$namespace/_namespace/transactions'
import { Route as ProtectedWorkspaceNamespaceNamespaceSettingsImport } from './routes/_protected/workspace/$namespace/_namespace/settings'
import { Route as ProtectedWorkspaceNamespaceNamespaceCategoriesImport } from './routes/_protected/workspace/$namespace/_namespace/categories'
import { Route as ProtectedWorkspaceNamespaceNamespaceAccountsIndexImport } from './routes/_protected/workspace/$namespace/_namespace/accounts/index'

// Create Virtual Routes

const PublicIndexLazyImport = createFileRoute('/_public/')()
const PublicSignupLazyImport = createFileRoute('/_public/signup')()
const PublicLoginLazyImport = createFileRoute('/_public/login')()
const ProtectedSetupLazyImport = createFileRoute('/_protected/setup')()
const ProtectedWorkspaceNamespaceNamespaceAccountsAccountIdLazyImport =
  createFileRoute(
    '/_protected/workspace/$namespace/_namespace/accounts/$accountId',
  )()

// Create/Update Routes

const PublicRoute = PublicImport.update({
  id: '/_public',
  getParentRoute: () => rootRoute,
} as any)

const ProtectedRoute = ProtectedImport.update({
  id: '/_protected',
  getParentRoute: () => rootRoute,
} as any)

const PublicIndexLazyRoute = PublicIndexLazyImport.update({
  path: '/',
  getParentRoute: () => PublicRoute,
} as any).lazy(() => import('./routes/_public/index.lazy').then((d) => d.Route))

const PublicSignupLazyRoute = PublicSignupLazyImport.update({
  path: '/signup',
  getParentRoute: () => PublicRoute,
} as any).lazy(() =>
  import('./routes/_public/signup.lazy').then((d) => d.Route),
)

const PublicLoginLazyRoute = PublicLoginLazyImport.update({
  path: '/login',
  getParentRoute: () => PublicRoute,
} as any).lazy(() => import('./routes/_public/login.lazy').then((d) => d.Route))

const ProtectedSetupLazyRoute = ProtectedSetupLazyImport.update({
  path: '/setup',
  getParentRoute: () => ProtectedRoute,
} as any).lazy(() =>
  import('./routes/_protected/setup.lazy').then((d) => d.Route),
)

const ProtectedWorkspaceIndexRoute = ProtectedWorkspaceIndexImport.update({
  path: '/workspace/',
  getParentRoute: () => ProtectedRoute,
} as any)

const ProtectedWorkspaceNamespaceRouteRoute =
  ProtectedWorkspaceNamespaceRouteImport.update({
    path: '/workspace/$namespace',
    getParentRoute: () => ProtectedRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceRoute =
  ProtectedWorkspaceNamespaceNamespaceImport.update({
    id: '/_namespace',
    getParentRoute: () => ProtectedWorkspaceNamespaceRouteRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceIndexRoute =
  ProtectedWorkspaceNamespaceNamespaceIndexImport.update({
    path: '/',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceTransactionsRoute =
  ProtectedWorkspaceNamespaceNamespaceTransactionsImport.update({
    path: '/transactions',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceSettingsRoute =
  ProtectedWorkspaceNamespaceNamespaceSettingsImport.update({
    path: '/settings',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceCategoriesRoute =
  ProtectedWorkspaceNamespaceNamespaceCategoriesImport.update({
    path: '/categories',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceAccountsIndexRoute =
  ProtectedWorkspaceNamespaceNamespaceAccountsIndexImport.update({
    path: '/accounts/',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any)

const ProtectedWorkspaceNamespaceNamespaceAccountsAccountIdLazyRoute =
  ProtectedWorkspaceNamespaceNamespaceAccountsAccountIdLazyImport.update({
    path: '/accounts/$accountId',
    getParentRoute: () => ProtectedWorkspaceNamespaceNamespaceRoute,
  } as any).lazy(() =>
    import(
      './routes/_protected/workspace/$namespace/_namespace/accounts/$accountId.lazy'
    ).then((d) => d.Route),
  )

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/_protected': {
      preLoaderRoute: typeof ProtectedImport
      parentRoute: typeof rootRoute
    }
    '/_public': {
      preLoaderRoute: typeof PublicImport
      parentRoute: typeof rootRoute
    }
    '/_protected/setup': {
      preLoaderRoute: typeof ProtectedSetupLazyImport
      parentRoute: typeof ProtectedImport
    }
    '/_public/login': {
      preLoaderRoute: typeof PublicLoginLazyImport
      parentRoute: typeof PublicImport
    }
    '/_public/signup': {
      preLoaderRoute: typeof PublicSignupLazyImport
      parentRoute: typeof PublicImport
    }
    '/_public/': {
      preLoaderRoute: typeof PublicIndexLazyImport
      parentRoute: typeof PublicImport
    }
    '/_protected/workspace/$namespace': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceRouteImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/': {
      preLoaderRoute: typeof ProtectedWorkspaceIndexImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/$namespace/_namespace': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
      parentRoute: typeof ProtectedWorkspaceNamespaceRouteImport
    }
    '/_protected/workspace/$namespace/_namespace/categories': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceCategoriesImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
    '/_protected/workspace/$namespace/_namespace/settings': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceSettingsImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
    '/_protected/workspace/$namespace/_namespace/transactions': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceTransactionsImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
    '/_protected/workspace/$namespace/_namespace/': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceIndexImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
    '/_protected/workspace/$namespace/_namespace/accounts/$accountId': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceAccountsAccountIdLazyImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
    '/_protected/workspace/$namespace/_namespace/accounts/': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceNamespaceAccountsIndexImport
      parentRoute: typeof ProtectedWorkspaceNamespaceNamespaceImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  ProtectedRoute.addChildren([
    ProtectedSetupLazyRoute,
    ProtectedWorkspaceNamespaceRouteRoute.addChildren([
      ProtectedWorkspaceNamespaceNamespaceRoute.addChildren([
        ProtectedWorkspaceNamespaceNamespaceCategoriesRoute,
        ProtectedWorkspaceNamespaceNamespaceSettingsRoute,
        ProtectedWorkspaceNamespaceNamespaceTransactionsRoute,
        ProtectedWorkspaceNamespaceNamespaceIndexRoute,
        ProtectedWorkspaceNamespaceNamespaceAccountsAccountIdLazyRoute,
        ProtectedWorkspaceNamespaceNamespaceAccountsIndexRoute,
      ]),
    ]),
    ProtectedWorkspaceIndexRoute,
  ]),
  PublicRoute.addChildren([
    PublicLoginLazyRoute,
    PublicSignupLazyRoute,
    PublicIndexLazyRoute,
  ]),
])

/* prettier-ignore-end */
