// @generated
// This file is automatically generated by Kanel. Do not modify manually.

import { fijoyWorkspaceId, type FijoyWorkspaceId } from './FijoyWorkspace';
import { z } from 'zod';

export type FijoyTagId = string;

/** Represents the table public.fijoy_tag */
export default interface FijoyTag {
  /** Database type: pg_catalog.text */
  id: FijoyTagId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name: string;
}

/** Represents the initializer for the table public.fijoy_tag */
export interface FijoyTagInitializer {
  /** Database type: pg_catalog.text */
  id: FijoyTagId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name: string;
}

/** Represents the mutator for the table public.fijoy_tag */
export interface FijoyTagMutator {
  /** Database type: pg_catalog.text */
  id?: FijoyTagId;

  /** Database type: pg_catalog.text */
  workspaceId?: FijoyWorkspaceId;

  /** Database type: pg_catalog.text */
  name?: string;
}

export const fijoyTagId = z.string();

export const fijoyTag = z.object({
  id: fijoyTagId,
  workspaceId: fijoyWorkspaceId,
  name: z.string(),
});

export const fijoyTagInitializer = z.object({
  id: fijoyTagId,
  workspaceId: fijoyWorkspaceId,
  name: z.string(),
});

export const fijoyTagMutator = z.object({
  id: fijoyTagId.optional(),
  workspaceId: fijoyWorkspaceId.optional(),
  name: z.string().optional(),
});
