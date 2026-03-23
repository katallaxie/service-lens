import { TWorkloadInsertSchema } from '@/db/schema'
import { ZodFormState } from '@/types'

export type AddWorkloadModalFormState = ZodFormState<TWorkloadInsertSchema> | null
