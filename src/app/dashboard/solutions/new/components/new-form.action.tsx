'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { addSolution } from '@/db/services/solutions'
import { v4 as uuidv4 } from 'uuid'

export const rhfAction = createAction(
  protectedProcedure.input(rhfActionSchema).mutation(
    async opts =>
      await addSolution({
        id: uuidv4(),
        ...opts.input
      })
  )
)