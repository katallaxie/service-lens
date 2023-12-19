import { protectedProcedure } from '../../trpc'
import { ProfileListSchema, ProfileGetSchema } from '../schemas/profile'

import { findAndCountProfiles, findOneProfile } from '@/db/services/profiles'

export const listProfiles = protectedProcedure
  .input(ProfileListSchema)
  .query(async opts => await findAndCountProfiles({ ...opts.input }))
export const getProfile = protectedProcedure
  .input(ProfileGetSchema)
  .query(async opts => await findOneProfile(opts.input))
