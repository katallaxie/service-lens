import { TTagInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddTagModalFormState = ZodFormState<TTagInsertSchema> | null
