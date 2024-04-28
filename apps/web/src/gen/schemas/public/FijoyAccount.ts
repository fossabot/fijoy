// @generated
// This file is automatically generated by Kanel. Do not modify manually.

import { fijoyWorkspaceId, type FijoyWorkspaceId } from './FijoyWorkspace';
import { fijoyAccountType, type default as FijoyAccountType } from './FijoyAccountType';
import { z } from 'zod';

export type FijoyAccountId = string;

/** Represents the table public.fijoy_account */
export default interface FijoyAccount {
  /** Database type: pg_catalog.text */
  id: FijoyAccountId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name: string;

  /** Database type: public.fijoy_account_type */
  accountType: FijoyAccountType;

  /** Database type: pg_catalog.numeric */
  balance: string;

  /** Database type: pg_catalog.text */
  currency: string;

  /** Database type: pg_catalog.text */
  institution: string;

  /** Database type: pg_catalog.bool */
  active: boolean;

  /** Database type: pg_catalog.timestamptz */
  createdAt: Date;

  /** Database type: pg_catalog.timestamptz */
  updatedAt: Date;
}

/** Represents the initializer for the table public.fijoy_account */
export interface FijoyAccountInitializer {
  /** Database type: pg_catalog.text */
  id: FijoyAccountId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name: string;

  /** Database type: public.fijoy_account_type */
  accountType: FijoyAccountType;

  /** Database type: pg_catalog.numeric */
  balance: string;

  /** Database type: pg_catalog.text */
  currency: string;

  /** Database type: pg_catalog.text */
  institution: string;

  /**
   * Database type: pg_catalog.bool
   * Default value: true
   */
  active?: boolean;

  /**
   * Database type: pg_catalog.timestamptz
   * Default value: now()
   */
  createdAt?: Date;

  /**
   * Database type: pg_catalog.timestamptz
   * Default value: now()
   */
  updatedAt?: Date;
}

/** Represents the mutator for the table public.fijoy_account */
export interface FijoyAccountMutator {
  /** Database type: pg_catalog.text */
  id?: FijoyAccountId;

  /** Database type: pg_catalog.text */
  workspaceId?: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name?: string;

  /** Database type: public.fijoy_account_type */
  accountType?: FijoyAccountType;

  /** Database type: pg_catalog.numeric */
  balance?: string;

  /** Database type: pg_catalog.text */
  currency?: string;

  /** Database type: pg_catalog.text */
  institution?: string;

  /** Database type: pg_catalog.bool */
  active?: boolean;

  /** Database type: pg_catalog.timestamptz */
  createdAt?: Date;

  /** Database type: pg_catalog.timestamptz */
  updatedAt?: Date;
}

export const fijoyAccountId = z.string();

export const fijoyAccount = z.object({
  id: fijoyAccountId,
  workspaceId: fijoyWorkspaceId,
  name: z.string(),
  accountType: fijoyAccountType,
  balance: z.string(),
  currency: z.string(),
  institution: z.string(),
  active: z.boolean(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export const fijoyAccountInitializer = z.object({
  id: fijoyAccountId,
  workspaceId: fijoyWorkspaceId,
  name: z.string(),
  accountType: fijoyAccountType,
  balance: z.string(),
  currency: z.string(),
  institution: z.string(),
  active: z.boolean().optional(),
  createdAt: z.coerce.date().optional(),
  updatedAt: z.coerce.date().optional(),
});

export const fijoyAccountMutator = z.object({
  id: fijoyAccountId.optional(),
  workspaceId: fijoyWorkspaceId.optional(),
  name: z.string().optional(),
  accountType: fijoyAccountType.optional(),
  balance: z.string().optional(),
  currency: z.string().optional(),
  institution: z.string().optional(),
  active: z.boolean().optional(),
  createdAt: z.coerce.date().optional(),
  updatedAt: z.coerce.date().optional(),
});
