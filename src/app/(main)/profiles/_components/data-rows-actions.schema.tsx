import { TEnvironmentSelectSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type DeleteEnvironmentSchema = ZodFormState<TEnvironmentSelectSchema> | null
