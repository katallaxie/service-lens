import { TLensInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddLensModalFormState = ZodFormState<TLensInsertSchema> | null
