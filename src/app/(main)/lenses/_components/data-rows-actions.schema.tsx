import { TLensDeleteSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type DeleteLensSchema = ZodFormState<TLensDeleteSchema> | null
