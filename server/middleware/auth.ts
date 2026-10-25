import { getCurrentUser } from '~~/server/utils/auth'

export default defineEventHandler(async (event) => {
  const url = getRequestURL(event)
  // Only protect API routes that need auth
  if (url.pathname.startsWith('/api/blood-pressure') || 
      url.pathname.startsWith('/api/weight')) {
    const user = await getCurrentUser(event)
    if (!user) {
      throw createError({
        statusCode: 401,
        statusMessage: 'Unauthorized'
      })
    }
    event.context.user = user
  }
})
