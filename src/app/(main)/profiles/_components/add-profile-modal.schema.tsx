import { TProfileInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddProfileModalFormState = ZodFormState<TProfileInsertSchema> | null
