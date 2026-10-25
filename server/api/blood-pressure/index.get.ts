import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const query = getQuery(event)
  
  if (query.all === 'true') {
    const readings = await db.prepare(
      'SELECT * FROM blood_pressures WHERE user_id = ? AND deleted_at IS NULL ORDER BY reading_date DESC'
    ).bind(user.id).all()
    return { readings: readings.results }
  }
  
  const limit = query.limit ? parseInt(query.limit as string) : 10
  const readings = await db.prepare(
    'SELECT * FROM blood_pressures WHERE user_id = ? AND deleted_at IS NULL ORDER BY reading_date DESC LIMIT ?'
  ).bind(user.id, limit).all()
  
  return { readings: readings.results }
})
