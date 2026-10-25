import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const id = getRouterParam(event, 'id')
  const body = await readBody(event)
  
  const reading = await db.prepare(
    'SELECT * FROM blood_pressures WHERE id = ? AND user_id = ? AND deleted_at IS NULL'
  ).bind(id, user.id).first()
  
  if (!reading) {
    throw createError({
      statusCode: 404,
      statusMessage: 'Reading not found'
    })
  }
  
  const { systolic, diastolic, pulse, reading_date } = body
  const now = new Date().toISOString()
  
  await db.prepare(
    'UPDATE blood_pressures SET systolic = ?, diastolic = ?, pulse = ?, reading_date = ?, updated_at = ? WHERE id = ?'
  ).bind(systolic, diastolic, pulse, reading_date, now, id).run()
  
  const updated = await db.prepare('SELECT * FROM blood_pressures WHERE id = ?').bind(id).first()
  
  return { reading: updated }
})
