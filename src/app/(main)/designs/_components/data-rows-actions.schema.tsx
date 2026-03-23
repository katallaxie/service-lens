import { TDesignSelectSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type DeleteDesignSchema = ZodFormState<TDesignSelectSchema> | null
