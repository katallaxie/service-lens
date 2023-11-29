import { loggerLink } from '@trpc/client'
import { experimental_nextCacheLink } from '@trpc/next/app-dir/links/nextCache'
import { experimental_createTRPCNextAppDirServer } from '@trpc/next/app-dir/server'
import { auth } from '@/lib/auth'
import { appRouter } from '@/server/routers/_app'
import { cookies } from 'next/headers'
import SuperJSON from 'superjson'

export const api = experimental_createTRPCNextAppDirServer<typeof appRouter>({
  config() {
    return {
      transformer: SuperJSON,
      links: [
        loggerLink({
          enabled: op => true
        }),
        experimental_nextCacheLink({
          // requests are cached for 5 seconds
          revalidate: 5,
          router: appRouter,
          createContext: async () => {
            return {
              session: await auth(),
              headers: {
                cookie: cookies().toString(),
                'x-trpc-source': 'rsc-invoke'
              }
            }
          }
        })
      ]
    }
  }
})
