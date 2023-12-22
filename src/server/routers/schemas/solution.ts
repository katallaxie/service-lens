import { PaginationSchema } from './pagination'
import { z } from 'zod'

export const SolutionListSchema = PaginationSchema
export const SolutionAddSchema = z.object({
  title: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  body: z.string()
})
export const SolutionGetSchema = z.string().uuid()
export const SolutionCommentDeleteSchema = z.string()
export const SolutionTemplateListSchema = PaginationSchema
export const SolutionTemplateGetSchema = z.string()
export const SolutionTemplateDeleteSchema = z.string()
