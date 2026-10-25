import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const id = getRouterParam(event, 'id')
  const body = await readBody(event)
  
  const entry = await db.prepare(
    'SELECT * FROM weights WHERE id = ? AND user_id = ? AND deleted_at IS NULL'
  ).bind(id, user.id).first()
  
  if (!entry) {
    throw createError({
      statusCode: 404,
      statusMessage: 'Entry not found'
    })
  }
  
  const { weight, entered_date } = body
  const now = new Date().toISOString()
  
  await db.prepare(
    'UPDATE weights SET weight = ?, entered_date = ?, updated_at = ? WHERE id = ?'
  ).bind(weight, entered_date, now, id).run()
  
  const updated = await db.prepare('SELECT * FROM weights WHERE id = ?').bind(id).first()
  
  return { weight: updated }
})
