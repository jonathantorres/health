import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const id = getRouterParam(event, 'id')
  
  const reading = await db.prepare(
    'SELECT * FROM blood_pressures WHERE id = ? AND user_id = ? AND deleted_at IS NULL'
  ).bind(id, user.id).first()
  
  if (!reading) {
    throw createError({
      statusCode: 404,
      statusMessage: 'Reading not found'
    })
  }
  
  const now = new Date().toISOString()
  await db.prepare('UPDATE blood_pressures SET deleted_at = ? WHERE id = ?').bind(now, id).run()
  
  return { success: true }
})
