import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const query = getQuery(event)
  
  if (query.all === 'true') {
    const weights = await db.prepare(
      'SELECT * FROM weights WHERE user_id = ? AND deleted_at IS NULL ORDER BY entered_date DESC'
    ).bind(user.id).all()
    return { weights: weights.results }
  }
  
  const limit = query.limit ? parseInt(query.limit as string) : 10
  const weights = await db.prepare(
    'SELECT * FROM weights WHERE user_id = ? AND deleted_at IS NULL ORDER BY entered_date DESC LIMIT ?'
  ).bind(user.id, limit).all()
  
  return { weights: weights.results }
})
