import { TDesignInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddDesignFormState = ZodFormState<TDesignInsertSchema> | null
