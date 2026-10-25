import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const id = getRouterParam(event, 'id')
  
  const entry = await db.prepare(
    'SELECT * FROM weights WHERE id = ? AND user_id = ? AND deleted_at IS NULL'
  ).bind(id, user.id).first()
  
  if (!entry) {
    throw createError({
      statusCode: 404,
      statusMessage: 'Entry not found'
    })
  }
  
  const now = new Date().toISOString()
  await db.prepare('UPDATE weights SET deleted_at = ? WHERE id = ?').bind(now, id).run()
  
  return { success: true }
})
