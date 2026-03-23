import { TEnvironmentInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddEnvironmentModalFormState = ZodFormState<TEnvironmentInsertSchema> | null
