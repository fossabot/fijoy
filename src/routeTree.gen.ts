// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as ProtectedImport } from './routes/_protected'
import { Route as ProtectedSetupImport } from './routes/_protected/setup'
import { Route as ProtectedWorkspaceIndexImport } from './routes/_protected/workspace/index'

// Create Virtual Routes

const SignupLazyImport = createFileRoute('/signup')()
const LoginLazyImport = createFileRoute('/login')()
const IndexLazyImport = createFileRoute('/')()
const ProtectedWorkspaceNamespaceIndexLazyImport = createFileRoute(
  '/_protected/workspace/$namespace/',
)()
const ProtectedWorkspaceNamespaceTransactionsLazyImport = createFileRoute(
  '/_protected/workspace/$namespace/transactions',
)()
const ProtectedWorkspaceNamespaceAccountsLazyImport = createFileRoute(
  '/_protected/workspace/$namespace/accounts',
)()

// Create/Update Routes

const SignupLazyRoute = SignupLazyImport.update({
  path: '/signup',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/signup.lazy').then((d) => d.Route))

const LoginLazyRoute = LoginLazyImport.update({
  path: '/login',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/login.lazy').then((d) => d.Route))

const ProtectedRoute = ProtectedImport.update({
  id: '/_protected',
  getParentRoute: () => rootRoute,
} as any)

const IndexLazyRoute = IndexLazyImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/index.lazy').then((d) => d.Route))

const ProtectedSetupRoute = ProtectedSetupImport.update({
  path: '/setup',
  getParentRoute: () => ProtectedRoute,
} as any)

const ProtectedWorkspaceIndexRoute = ProtectedWorkspaceIndexImport.update({
  path: '/workspace/',
  getParentRoute: () => ProtectedRoute,
} as any)

const ProtectedWorkspaceNamespaceIndexLazyRoute =
  ProtectedWorkspaceNamespaceIndexLazyImport.update({
    path: '/workspace/$namespace/',
    getParentRoute: () => ProtectedRoute,
  } as any).lazy(() =>
    import('./routes/_protected/workspace/$namespace/index.lazy').then(
      (d) => d.Route,
    ),
  )

const ProtectedWorkspaceNamespaceTransactionsLazyRoute =
  ProtectedWorkspaceNamespaceTransactionsLazyImport.update({
    path: '/workspace/$namespace/transactions',
    getParentRoute: () => ProtectedRoute,
  } as any).lazy(() =>
    import('./routes/_protected/workspace/$namespace/transactions.lazy').then(
      (d) => d.Route,
    ),
  )

const ProtectedWorkspaceNamespaceAccountsLazyRoute =
  ProtectedWorkspaceNamespaceAccountsLazyImport.update({
    path: '/workspace/$namespace/accounts',
    getParentRoute: () => ProtectedRoute,
  } as any).lazy(() =>
    import('./routes/_protected/workspace/$namespace/accounts.lazy').then(
      (d) => d.Route,
    ),
  )

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      preLoaderRoute: typeof IndexLazyImport
      parentRoute: typeof rootRoute
    }
    '/_protected': {
      preLoaderRoute: typeof ProtectedImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      preLoaderRoute: typeof LoginLazyImport
      parentRoute: typeof rootRoute
    }
    '/signup': {
      preLoaderRoute: typeof SignupLazyImport
      parentRoute: typeof rootRoute
    }
    '/_protected/setup': {
      preLoaderRoute: typeof ProtectedSetupImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/': {
      preLoaderRoute: typeof ProtectedWorkspaceIndexImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/$namespace/accounts': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceAccountsLazyImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/$namespace/transactions': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceTransactionsLazyImport
      parentRoute: typeof ProtectedImport
    }
    '/_protected/workspace/$namespace/': {
      preLoaderRoute: typeof ProtectedWorkspaceNamespaceIndexLazyImport
      parentRoute: typeof ProtectedImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  IndexLazyRoute,
  ProtectedRoute.addChildren([
    ProtectedSetupRoute,
    ProtectedWorkspaceIndexRoute,
    ProtectedWorkspaceNamespaceAccountsLazyRoute,
    ProtectedWorkspaceNamespaceTransactionsLazyRoute,
    ProtectedWorkspaceNamespaceIndexLazyRoute,
  ]),
  LoginLazyRoute,
  SignupLazyRoute,
])
