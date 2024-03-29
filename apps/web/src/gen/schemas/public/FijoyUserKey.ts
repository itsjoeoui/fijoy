// @generated
// This file is automatically generated by Kanel. Do not modify manually.

import { fijoyUserId, type FijoyUserId } from './FijoyUser';
import { z } from 'zod';

export type FijoyUserKeyId = string;

/** Represents the table public.fijoy_user_key */
export default interface FijoyUserKey {
  /** Database type: pg_catalog.text */
  id: FijoyUserKeyId;

  /** Database type: pg_catalog.text */
  userId: FijoyUserId;

  /** Database type: pg_catalog.text */
  hashedPassword: string | null;
}

/** Represents the initializer for the table public.fijoy_user_key */
export interface FijoyUserKeyInitializer {
  /** Database type: pg_catalog.text */
  id: FijoyUserKeyId;

  /** Database type: pg_catalog.text */
  userId: FijoyUserId;

  /** Database type: pg_catalog.text */
  hashedPassword?: string | null;
}

/** Represents the mutator for the table public.fijoy_user_key */
export interface FijoyUserKeyMutator {
  /** Database type: pg_catalog.text */
  id?: FijoyUserKeyId;

  /** Database type: pg_catalog.text */
  userId?: FijoyUserId;

  /** Database type: pg_catalog.text */
  hashedPassword?: string | null;
}

export const fijoyUserKeyId = z.string();

export const fijoyUserKey = z.object({
  id: fijoyUserKeyId,
  userId: fijoyUserId,
  hashedPassword: z.string().nullable(),
});

export const fijoyUserKeyInitializer = z.object({
  id: fijoyUserKeyId,
  userId: fijoyUserId,
  hashedPassword: z.string().optional().nullable(),
});

export const fijoyUserKeyMutator = z.object({
  id: fijoyUserKeyId.optional(),
  userId: fijoyUserId.optional(),
  hashedPassword: z.string().optional().nullable(),
});
