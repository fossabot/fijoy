// @generated by protoc-gen-connect-query v1.3.1 with parameter "target=ts"
// @generated from file fijoy/v1/workspace.proto (package fijoy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodIdempotency, MethodKind } from "@bufbuild/protobuf";
import { CreateWorkspaceRequest, GetWorkspaceByIdRequest, GetWorkspaceByNamespaceRequest, UpdateLocaleRequest, UpdatePrimaryCurrencyRequest, UpdateWorkspaceNameRequest, UpdateWorkspaceNamespaceRequest, Workspace, Workspaces } from "./workspace_pb.js";

/**
 * @generated from rpc fijoy.v1.WorkspaceService.CreateWorkspace
 */
export const createWorkspace = {
  localName: "createWorkspace",
  name: "CreateWorkspace",
  kind: MethodKind.Unary,
  I: CreateWorkspaceRequest,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.GetWorkspaces
 */
export const getWorkspaces = {
  localName: "getWorkspaces",
  name: "GetWorkspaces",
  kind: MethodKind.Unary,
  I: Empty,
  O: Workspaces,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.GetWorkspaceById
 */
export const getWorkspaceById = {
  localName: "getWorkspaceById",
  name: "GetWorkspaceById",
  kind: MethodKind.Unary,
  I: GetWorkspaceByIdRequest,
  O: Workspace,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.GetWorkspaceByNamespace
 */
export const getWorkspaceByNamespace = {
  localName: "getWorkspaceByNamespace",
  name: "GetWorkspaceByNamespace",
  kind: MethodKind.Unary,
  I: GetWorkspaceByNamespaceRequest,
  O: Workspace,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.UpdateWorkspaceName
 */
export const updateWorkspaceName = {
  localName: "updateWorkspaceName",
  name: "UpdateWorkspaceName",
  kind: MethodKind.Unary,
  I: UpdateWorkspaceNameRequest,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.UpdateWorkspaceNamespace
 */
export const updateWorkspaceNamespace = {
  localName: "updateWorkspaceNamespace",
  name: "UpdateWorkspaceNamespace",
  kind: MethodKind.Unary,
  I: UpdateWorkspaceNamespaceRequest,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.UpdatePrimaryCurrency
 */
export const updatePrimaryCurrency = {
  localName: "updatePrimaryCurrency",
  name: "UpdatePrimaryCurrency",
  kind: MethodKind.Unary,
  I: UpdatePrimaryCurrencyRequest,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.UpdateLocale
 */
export const updateLocale = {
  localName: "updateLocale",
  name: "UpdateLocale",
  kind: MethodKind.Unary,
  I: UpdateLocaleRequest,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;

/**
 * @generated from rpc fijoy.v1.WorkspaceService.DeleteWorkspace
 */
export const deleteWorkspace = {
  localName: "deleteWorkspace",
  name: "DeleteWorkspace",
  kind: MethodKind.Unary,
  I: Empty,
  O: Workspace,
  service: {
    typeName: "fijoy.v1.WorkspaceService"
  }
} as const;
